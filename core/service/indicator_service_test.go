package service

import (
	"alert/core/dao/indicator_dao"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {

	indicator_service := IndicatorServiceImpl{}
	indicator := indicator_dao.Indicator{IndicatorCode: "test", Name: "test", Expression: "test", Description: "test", CreateTime: time.Now(), UpdateTime: time.Now()}
	result := indicator_service.IndicatorAdd(indicator)
	if !result {
		t.Fatal("error")
	}
}

//func TestAdd1(t *testing.T) {
//	var TrueResult = 2
//	result := add(1, 1)
//	if result != TrueResult {
//		t.Fatal("error")
//	}
//}
