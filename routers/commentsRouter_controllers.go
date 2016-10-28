package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/dremygit/xwindy-lite/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/dremygit/xwindy-lite/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetBySno",
			Router: `/:sno`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/dremygit/xwindy-lite/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/dremygit/xwindy-lite/controllers:UserController"],
		beego.ControllerComments{
			Method: "CreateUser",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/dremygit/xwindy-lite/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/dremygit/xwindy-lite/controllers:UserController"],
		beego.ControllerComments{
			Method: "UpdateInfo",
			Router: `/:sno`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/dremygit/xwindy-lite/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/dremygit/xwindy-lite/controllers:UserController"],
		beego.ControllerComments{
			Method: "ResetPassword",
			Router: `/:sno/password`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

}
