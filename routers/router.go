package routers

import (
	"Miniswap-Api/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/ppconfig", &controllers.PrivatePlacementConfigController{})
	beego.Router("/ppinfo", &controllers.PrivatePlacementInfoController{})
	beego.Router("/pptxs", &controllers.PrivatePlacementTxsController{})
}
