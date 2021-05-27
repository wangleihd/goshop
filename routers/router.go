// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"goshop/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/api/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/index",
			beego.NSNamespace("/index",
				beego.NSInclude(
					&controllers.IndexController{},
				),
			),
			beego.NSNamespace("/lists",
				beego.NSInclude(
					&controllers.ListsController{},
				),
			),
		),
		beego.NSNamespace("/goods",
			beego.NSNamespace("/detail",
				beego.NSInclude(
					&controllers.DetailController{},
				),
			),
		),
	)
	beego.AddNamespace(ns)
}
