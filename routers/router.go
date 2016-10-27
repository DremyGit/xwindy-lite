package routers

import (
	"github.com/astaxie/beego"
	"github.com/dremygit/xwindy-lite/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/users",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
