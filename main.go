package main

import (
	"github.com/astaxie/beego"

	"github.com/dremygit/xwindy-lite/controllers"
	_ "github.com/dremygit/xwindy-lite/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}
