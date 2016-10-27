package controllers

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
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

func (c *BaseController) Success(code int, data interface{}) {
	c.Data["json"] = data
	toolbox.StatisticsMap.AddStatistics(c.Ctx.Input.Method(), c.Ctx.Input.URL(), "&MyController", time.Since(requestBegin))
	c.Ctx.ResponseWriter.WriteHeader(code)
	c.ServeJSON()
}

func (c *BaseController) Failure(code int, message interface{}) {
	fmt.Printf("[Err] - %s Request [%s]\t %s\t IP:%s\t %d: %s\n", time.Now().Format("2006-01-02 15:04:05"), c.Ctx.Input.Method(), c.Ctx.Input.URL(), c.Ctx.Input.IP(), code, message)
	c.Data["json"] = &ErrorResponse{code, message}
	toolbox.StatisticsMap.AddStatistics(c.Ctx.Input.Method(), c.Ctx.Input.URL(), "&MyController", time.Since(requestBegin))
	c.Ctx.ResponseWriter.WriteHeader(code)
	c.ServeJSON()
}
