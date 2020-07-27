// @APIVersion 1.0.0
// @Title Miniswap API
// @Description Data related to MINISwap for developers
package routers

import (
	"Miniswap-Api/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns :=
		beego.NewNamespace("/v1",
			beego.NSNamespace("/pptxs",
				beego.NSInclude(
					&controllers.PrivatePlacementTxsController{},
				),
			),
			beego.NSNamespace("/ppconfig",
				beego.NSInclude(
					&controllers.PrivatePlacementConfigController{},
				),
			),
			beego.NSNamespace("/ppinfo",
				beego.NSInclude(
					&controllers.PrivatePlacementInfoController{},
				),
			),
		)
	beego.AddNamespace(ns)
	beego.Router("/ppconfig", &controllers.PrivatePlacementConfigController{})
	beego.Router("/ppinfo", &controllers.PrivatePlacementInfoController{})
	beego.Router("/pptxs", &controllers.PrivatePlacementTxsController{})
}
