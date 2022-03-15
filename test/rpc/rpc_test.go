package rpc

import (
	"alert/kitex_gen/api"
	"alert/kitex_gen/api/crud"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"log"
	"testing"
	"time"
)

const ADDR = "127.0.0.1:8888"

var indicator = &api.Indicator{
	IndicatorCode: "test10",
	Name:          "test",
	Expression:    "test",
	Description:   "test",
	TimeCreate:    "2006-01-02 15:04:05",
	TimeUpdate:    "2006-01-02 15:04:05",
}
var indicatorJson = &api.IndicatorJson{
	Indicators: nil,
	Calculate:  -1,
	Value:      "select deal_amount from `deal_infos`",
}

func Connect() crud.Client {
	c, err := crud.NewClient("crud", client.WithHostPorts(ADDR))
	if err != nil {
		log.Fatalln(err)
	}
	return c
}

func TestRpcAddIndicator(t *testing.T) {
	c := Connect()
	addIndicatorReq := &api.AddIndicatorRequest{
		Indicator:     indicator,
		IndicatorJson: indicatorJson,
	}
	addIndicatorResp, err := c.AddIndicator(context.Background(), addIndicatorReq, callopt.WithConnectTimeout(3*time.Second))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(addIndicatorResp)
}
