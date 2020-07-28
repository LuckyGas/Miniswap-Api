package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"math"
	"math/big"
)

type PrivatePlacementConfigController struct {
	beego.Controller
}

type JsonRpcResult struct {
	Version string `json:"jsonrpc"`
	Id      uint   `json:"id"`
}

func DefaultJsonRpcResult() *JsonRpcResult {
	return &JsonRpcResult{"2.0", 1}
}

type ConfigResult struct {
	*JsonRpcResult
	Configs []Config
}

type Config struct {
	Index      int      `json:"round"`
	EthAmount  *big.Int `json:"ethValue"`
	MiniAmount *big.Int `json:"miniAmount"`
	Ratio      *big.Int `json:"ratio"`
}

func initConfig() {
	if len(Configs) == 0 {
		l := 30
		Configs = make([]Config, l, l)
		for i := 1; i <= l; i++ {
			ethAmount := 20 + 10*(i-1)
			price := math.Pow(0.93, float64(l-i)) * 0.02
			miniAmount := math.Floor(float64(ethAmount) * 246 / price)
			ratio := new(big.Int).SetInt64(int64(math.Floor(miniAmount / float64(ethAmount))))
			Configs[i-1] = Config{i, new(big.Int).SetInt64(int64(ethAmount)), new(big.Int).SetInt64(int64(miniAmount)), ratio}
		}
	}
}

func GetConfig(index int64) (config Config) {
	initConfig()
	config = Configs[index]
	fmt.Println(config)
	return
}

var Configs []Config

// @Title getConfig
// @Description get the private placement config
// @Success 200 {object} controllers.ppconfig.ConfigResult
// @Failure 404 User not found
// @router / [get]
func (c *PrivatePlacementConfigController) Get() {
	initConfig()
	jr := ConfigResult{DefaultJsonRpcResult(), Configs}
	c.Data["json"] = jr
	c.ServeJSON()
}
