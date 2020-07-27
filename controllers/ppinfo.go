package controllers

import (
	"github.com/astaxie/beego"
)

type PrivatePlacementInfoController struct {
	beego.Controller
}

type InfoResult struct {
	*JsonRpcResult
	CurIndex        uint   `json:"round"`
	CurEthValue     uint   `json:"cur_eth_value"`
	RemainingRation string `json:"remaining_ration"`
	RemainingMini   uint   `json:"remaining_mini"`
}

// @Title getCurrentPrivatepalcementInfo
// @Description get the current private placement info
// @Success 200 {object} controllers.ppinfo.InfoResult
// @Failure 404 User not found
// @router / [get]
func (c *PrivatePlacementInfoController) Get() {
	ir := InfoResult{}
	ir.JsonRpcResult = DefaultJsonRpcResult()
	ir.CurIndex = 1
	ir.CurEthValue = 0
	ir.RemainingRation = "100%"
	ir.RemainingMini = 2018093
	c.Data["json"] = ir
	c.ServeJSON()
}
