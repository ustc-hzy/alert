package service

import (
	"alert/core"
	"alert/core/dao/indicator_dao"
	"fmt"
	"testing"
	"time"
)

var indicator_service = IndicatorServiceImpl{}
var ind_compute_service = IndComputeImpl{}

func TestAdd(t *testing.T) {

	indicator := indicator_dao.Indicator{IndicatorCode: "test7", Name: "test", Expression: "test", Description: "test", CreateTime: time.Now(), UpdateTime: time.Now()}
	indicatorJson := indicator_dao.IndicatorJson{Indicators: nil, Caculate: core.NIL, Value: "select deal_amount from `deal_infos`"}
	result := indicator_service.Add(indicator, indicatorJson)
	if !result {
		t.Fatal("error")
	}
}

func TestDelete(t *testing.T) {

	result := indicator_service.Delete("test")
	if !result {
		t.Fatal("error")
	}
}

func TestQuery(t *testing.T) {

	query := indicator_service.Query("test1")
	fmt.Println(query)
}

func TestModify(t *testing.T) {
	indicator := indicator_dao.Indicator{IndicatorCode: "test", Name: "test", Description: "modify", UpdateTime: time.Now()}
	indicator_service.Modify(indicator)
}

func TestCompute(t *testing.T) {
	val := ind_compute_service.Compute("test2", 0, time.Now(), time.Now())
	fmt.Println(val)
}
