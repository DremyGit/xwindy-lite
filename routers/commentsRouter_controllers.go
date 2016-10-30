package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["github.com/dremygit/xwindy-lite/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/dremygit/xwindy-lite/controllers:CommentController"],
		beego.ControllerComments{
			Method: "GetCommentListByNewsID",
			Router: `/:newsid/comments`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/dremygit/xwindy-lite/controllers:CommentController"] = append(beego.GlobalControllerRouter["github.com/dremygit/xwindy-lite/controllers:CommentController"],
		beego.ControllerComments{
			Method: "CreateComment",
			Router: `/:newsid/comments`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/dremygit/xwindy-lite/controllers:NewsController"] = append(beego.GlobalControllerRouter["github.com/dremygit/xwindy-lite/controllers:NewsController"],
		beego.ControllerComments{
			Method: "GetNewsList",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/dremygit/xwindy-lite/controllers:NewsController"] = append(beego.GlobalControllerRouter["github.com/dremygit/xwindy-lite/controllers:NewsController"],
		beego.ControllerComments{
			Method: "GetNewsByID",
			Router: `/:newsid`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["github.com/dremygit/xwindy-lite/controllers:NewsController"] = append(beego.GlobalControllerRouter["github.com/dremygit/xwindy-lite/controllers:NewsController"],
		beego.ControllerComments{
			Method: "IncreaseClickCount",
			Router: `/:newsid/click_count`,
			AllowHTTPMethods: []string{"put"},
			Params: nil})

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

	beego.GlobalControllerRouter["github.com/dremygit/xwindy-lite/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/dremygit/xwindy-lite/controllers:UserController"],
		beego.ControllerComments{
			Method: "Authorize",
			Router: `/authorization`,
			AllowHTTPMethods: []string{"post"},
			Params: nil})

}
