package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["Miniswap-Api/controllers:PrivatePlacementConfigController"] = append(beego.GlobalControllerRouter["Miniswap-Api/controllers:PrivatePlacementConfigController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Miniswap-Api/controllers:PrivatePlacementInfoController"] = append(beego.GlobalControllerRouter["Miniswap-Api/controllers:PrivatePlacementInfoController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["Miniswap-Api/controllers:PrivatePlacementTxsController"] = append(beego.GlobalControllerRouter["Miniswap-Api/controllers:PrivatePlacementTxsController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
