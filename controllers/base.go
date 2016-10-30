package controllers

import (
	"errors"
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/toolbox"
	. "github.com/bitly/go-simplejson"
	jwt "github.com/dgrijalva/jwt-go"
)

var (
	API_BASE string
)

type BaseController struct {
	beego.Controller
}

var requestBegin time.Time

func (c *BaseController) Prepare() {
	requestBegin = time.Now()
}

type ErrorResponse struct {
	StatusCode int         `json:"status_code"`
	Message    interface{} `json:"message"`
}

func (c *BaseController) ParsePayload() (map[string]interface{}, error) {
	js, err := NewJson(c.Ctx.Input.RequestBody)
	if err != nil {
		return nil, errors.New("JSON 格式错误")
	}

	payload, err := js.Map()
	if err != nil {
		return nil, errors.New("JSON 结构错误")
	}

	return payload, nil
}

func (c *BaseController) Success(code int, data interface{}) {
	c.Data["json"] = data
	toolbox.StatisticsMap.AddStatistics(c.Ctx.Input.Method(), c.Ctx.Input.URL(), "&BaseController", time.Since(requestBegin))
	c.Ctx.Output.SetStatus(code)
	c.ServeJSON()
}

func (c *BaseController) Failure(code int, message interface{}) {
	fmt.Printf("[Err] - %s Request [%s]\t %s\t IP:%s\t %d: %s\n", time.Now().Format("2006-01-02 15:04:05"), c.Ctx.Input.Method(), c.Ctx.Input.URL(), c.Ctx.Input.IP(), code, message)
	c.Data["json"] = &ErrorResponse{code, message}
	toolbox.StatisticsMap.AddStatistics(c.Ctx.Input.Method(), c.Ctx.Input.URL(), "&BaseController", time.Since(requestBegin))
	c.Ctx.Output.SetStatus(code)
	c.ServeJSON()
}

func (c *BaseController) ParseToken() (map[string]interface{}, error) {
	token, err := jwt.ParseFromRequest(c.Ctx.Request, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("test"), nil
	})
	if err != nil {
		return nil, err
	}

	return token.Claims, nil
}

func init() {
	iniconf, err := config.NewConfig("ini", "./conf/app.conf")
	if err != nil {
		panic("Config file not found")
	}

	API_BASE = iniconf.String("path::hostname") + iniconf.String("path::basepath")
}
