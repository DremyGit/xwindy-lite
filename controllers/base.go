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

// BaseController is the base controller of all controllers
type BaseController struct {
	beego.Controller
}

var requestBegin time.Time

// Prepare to set the time of the request begin
func (c *BaseController) Prepare() {
	requestBegin = time.Now()
}

// ErrorResponse is the error payload
type ErrorResponse struct {
	StatusCode int         `json:"status_code"`
	Message    interface{} `json:"message"`
}

// Success to response data and set status code
func (c *BaseController) Success(code int, data interface{}) {
	c.Data["json"] = data
	toolbox.StatisticsMap.AddStatistics(c.Ctx.Input.Method(), c.Ctx.Input.URL(), "&BaseController", time.Since(requestBegin))
	c.Ctx.Output.SetStatus(code)
	c.ServeJSON()
}

// Failure to response error message and set status code
func (c *BaseController) Failure(code int, message interface{}) {
	c.Data["json"] = &ErrorResponse{code, message}
	toolbox.StatisticsMap.AddStatistics(c.Ctx.Input.Method(), c.Ctx.Input.URL(), "&BaseController", time.Since(requestBegin))
	c.Ctx.Output.SetStatus(code)
	c.ServeJSON()
}

// ParsePayload to parse the request payload to map
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

// ParseToken to parse the authorization token in request header to map
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
