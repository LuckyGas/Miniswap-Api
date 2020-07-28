package controllers

import (
	"github.com/astaxie/beego"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"strconv"
	"strings"
)

type PrivatePlacementInfoController struct {
	beego.Controller
}

type InfoResult struct {
	*JsonRpcResult
	CurIndex        uint     `json:"round"`
	CurEthValue     *big.Int `json:"cur_eth_value"`
	RemainingRation *big.Int `json:"remaining_ration"`
	RemainingMini   *big.Int `json:"remaining_mini"`
}

func getDataFromContract() (k int64, ethAmount *big.Int) {
	rpcDial, err := rpc.Dial("https://ropsten.infura.io/v3/638c755c81fe495e85debe581520b373")
	if err != nil {
		panic(err)
	}
	result := new(interface{})
	args := map[string]interface{}{
		"to":   "0x9220502202c420B2249176aAD4dAba5E1C3281E1",
		"data": "0xb4f40c61",
	}
	err = rpcDial.Call(result, "eth_call", args, "latest")
	if strings.HasPrefix((*result).(string), "0x") {
		*result = strings.Replace((*result).(string), "0x", "", -1)
		args["data"] = "0xba2cf3aa" + (*result).(string)
	}
	k, err = strconv.ParseInt((*result).(string), 16, 10)
	err = rpcDial.Call(result, "eth_call", args, "latest")
	if strings.HasPrefix((*result).(string), "0x") {
		*result = strings.Replace((*result).(string), "0x", "", -1)
		args["data"] = "0xba2cf3aa" + (*result).(string)
	}
	ethAmount, _ = new(big.Int).SetString((*result).(string), 16)
	rpcDial.Close()
	return
}

// @Title getCurrentPrivatepalcementInfo
// @Description get the current private placement info
// @Success 200 {object} controllers.ppinfo.InfoResult
// @Failure 404 User not found
// @router / [get]
func (c *PrivatePlacementInfoController) Get() {
	k, ethAmount := getDataFromContract()
	ir := InfoResult{}
	ir.JsonRpcResult = DefaultJsonRpcResult()
	ir.CurIndex = uint(k)
	ir.CurEthValue = ethAmount
	needEth := GetConfig(k - 1).EthAmount
	needEth = new(big.Int).Mul(needEth, new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil))
	ir.RemainingRation = new(big.Int).Div(new(big.Int).Mul(ethAmount, big.NewInt(100)), needEth)
	ir.RemainingMini = new(big.Int).Div(new(big.Int).Mul(ir.RemainingRation, GetConfig(k-1).MiniAmount), big.NewInt(100))
	c.Data["json"] = ir
	c.ServeJSON()
}
