package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/dremygit/xwindy-lite/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/dremygit/xwindy-lite/controllers:UserController"],
		beego.ControllerComments{
			Method: "GetBySno",
			Router: `/:userId`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/dremygit/xwindy-lite/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/dremygit/xwindy-lite/controllers:UserController"],
		beego.ControllerComments{
			Method: "CreateUser",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
