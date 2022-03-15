package test

import (
	"alert/core"
	"alert/core/dao/indicator_dao"
	"alert/core/service"
	"alert/core/vo"
	"fmt"
	"testing"
	"time"
)

var indicator_service = service.IndicatorServiceImpl{}
var ind_compute_service = service.IndComputeImpl{}
var testCondition = vo.Condition{
	RoomID: 1,
	//StartTime: "'2022-03-01'",
	//EndTime:   "'2022-03-02'",
}

func TestAdd(t *testing.T) {

	indicator := indicator_dao.Indicator{
		IndicatorCode: "userAmount",
		Name:          "",
		Expression:    "",
		Description:   "user num",
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	}
	indicatorJson := indicator_dao.IndicatorJson{
		Indicators: nil,
		Caculate:   core.NIL,
		Value:      "select count(1) as userNum,room_id from ( select distinct user_id,room_id from `deal_infos ) as m`",
	}
	result := indicator_service.Add(indicator, indicatorJson)
	if !result {
		t.Fatal("error")
	}
}

func TestDelete(t *testing.T) {

	result := indicator_service.Delete("test98")
	if !result {
		t.Fatal("error")
	}
}

func TestQuery(t *testing.T) {

	query := indicator_service.Query("test99")
	fmt.Println(query)
}

func TestModify(t *testing.T) {
	indicator := indicator_dao.Indicator{
		IndicatorCode: "amount",
		Name:          "test",
		Description:   "trade sum",
		UpdateTime:    time.Now(),
	}
	indicator_service.Modify(indicator)
}

func TestCompute(t *testing.T) {
	// left close right open
	val := ind_compute_service.Compute("userAmount", testCondition)
	fmt.Println(val)
}

func TestAddPricePerPeople(t *testing.T) {
	moneyAmount := indicator_service.Query("moneyAmount")
	userAmount := indicator_service.Query("userAmount")
	indicator := indicator_dao.Indicator{
		IndicatorCode: "pricePerPeople",
		Name:          "",
		Expression:    "",
		Description:   "pricePerPeople",
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	}
	inds := []vo.IndicatorVO{moneyAmount, userAmount}
	indicatorJson := indicator_dao.IndicatorJson{
		Indicators: inds,
		Caculate:   core.DIVIDE,
		Value:      "",
	}
	result := indicator_service.Add(indicator, indicatorJson)
	if !result {
		t.Fatal("error")
	}

}

func TestAddCompound(t *testing.T) {

	indicatorA := vo.IndicatorVO{
		IndicatorCode: "nodeA",
		Name:          "test",
		Indicators:    nil,
		Caculate:      core.NIL,
		Value:         "select deal_amount from `deal_infos`",
		Description:   "test",
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	}
	indicatorB := vo.IndicatorVO{
		IndicatorCode: "nodeB",
		Name:          "test",
		Indicators:    nil,
		Caculate:      core.NIL,
		Value:         "select deal_amount from `deal_infos`",
		Description:   "test",
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	}

	indicator := indicator_dao.Indicator{
		IndicatorCode: "nodeALL2",
		Name:          "test",
		Expression:    "test",
		Description:   "test",
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	}
	inds := []vo.IndicatorVO{indicatorA, indicatorB}
	indicatorJson := indicator_dao.IndicatorJson{
		Indicators: inds,
		Caculate:   core.ADD,
		Value:      "",
	}

	result := indicator_service.Add(indicator, indicatorJson)
	if !result {
		t.Fatal("error")
	}
}

func TestComputeCompound(t *testing.T) {
	val := ind_compute_service.Compute("pricePerPeople", testCondition)
	fmt.Println(val)
}
