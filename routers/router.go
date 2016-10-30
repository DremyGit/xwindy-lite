// @APIVersion 1.0.0
// @Title xwindy-lite API
// @Description xwindy-lite is a mini community of news of students
// @Contact fhgsm123@gmail.com
// @TermsOfServiceUrl http://api.xwindy.com/
// @License GPL v3.0
// @LicenseUrl https://github.com/DremyGit/xwindy-lite/blob/master/LICENSE
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
		beego.NSNamespace("/news",
			beego.NSInclude(
				&controllers.CommentController{},
				&controllers.NewsController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
