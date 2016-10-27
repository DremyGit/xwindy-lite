package routers

import (
	"github.com/astaxie/beego"
	"github.com/dremygit/xwindy-lite/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
}
