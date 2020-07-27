package controllers

import (
	"github.com/astaxie/beego"
	"math/big"
)

type PrivatePlacementTxsController struct {
	beego.Controller
}

type TxsResult struct {
	*JsonRpcResult
	TotalTxCount uint
	Offset int `json:"offset"`
	Limit int `json:"limit"`
	Txs []Tx `json:"txs"`
}

type Tx struct {
	Timestamp uint `json:"timestamp"`
	SendEthValue *big.Int `json:"send_eth_value"`
	ReceiveMiniAmount *big.Int `json:"receive_mini_amount"`
	Fee *big.Int `json:"fee"`
	TxHash string `json:"tx_hash"`
	Status string `json:"status"`
}

func (c PrivatePlacementTxsController) Get() {
	tr := TxsResult{}
	tr.JsonRpcResult = DefaultJsonRpcResult()
	tr.TotalTxCount = 100
	tr.Limit,_= c.GetInt("limit")
	tr.Offset,_ = c.GetInt("offset")
	tr.Txs = make([]Tx,tr.Limit,tr.Limit)
	for i:=0;i< tr.Limit;i++{
		tr.Txs[i] = Tx{1588598533,big.NewInt(2200000000000000),big.NewInt(2200000000000000),big.NewInt(2200000000000000),"0x0d7b1a9d64a284a98df5373191c0ab30c9368aec03a127f9e7a72a79187bec2d","success"}
	}
	c.Data["json"] = tr
	c.ServeJSON()
}