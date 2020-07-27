package controllers

import (
	"github.com/astaxie/beego"
	"math"
)

type PrivatePlacementConfigController struct {
	beego.Controller
}

type JsonRpcResult struct {
	Version string `json:"jsonrpc"`
	Id uint `json:"id"`
}

func DefaultJsonRpcResult() *JsonRpcResult {
	return &JsonRpcResult{"2.0",1}
}

type ConfigResult struct {
	*JsonRpcResult
	Configs []Config
}

type Config struct {
	Index int `json:"round"`
	EthAmount float64 `json:"ethValue"`
	MiniAmount float64 `json:"miniAmount"`
	Ratio float64 `json:"ratio"`
}

func initConfigResult() (jr *ConfigResult){
	l:=30
	jr = &ConfigResult{DefaultJsonRpcResult(),make([]Config,l,l)}
	for i:=1;i<=l;i++{
		ethAmount := float64(20+10*(i-1))
		price := math.Pow(0.93,float64(l-i)) * 0.02
		miniAmount := math.Floor(float64(ethAmount) * 246 / price)
		ratio := math.Floor(miniAmount / ethAmount)
		jr.Configs[i-1] = Config{i,ethAmount,miniAmount,ratio}
	}
	return
}

func (c *PrivatePlacementConfigController) Get() {
	jr:=initConfigResult()
	c.Data["json"] = jr
	c.ServeJSON()
}
