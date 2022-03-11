package service

import (
	"alert/core"
	"alert/core/dao"
	"alert/core/dao/indicator_dao"
	"alert/core/vo"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type IndicatorServiceImpl struct{}
type IndComputeImpl struct{}

const INDICATORABLENAME = "indicators"

func (i IndicatorServiceImpl) Serialization(indicatorJson indicator_dao.IndicatorJson) string {
	result, _ := json.Marshal(indicatorJson)
	return string(result)
}
func (i IndicatorServiceImpl) AntiSerialization(expression string) indicator_dao.IndicatorJson {
	byteStream := []byte(expression)
	var result indicator_dao.IndicatorJson
	err := json.Unmarshal(byteStream, &result)
	if err != nil {
		log.Fatalln(err)
	}
	return result
}

func (i IndicatorServiceImpl) Add(indicator indicator_dao.Indicator, indicatorJson indicator_dao.IndicatorJson) bool {

	indicator.Expression = i.Serialization(indicatorJson)
	if IsIndicatorExist(indicator.IndicatorCode) {
		log.Fatalln("the indicator code already exist")
		return false
	}
	res1 := dao.DB.Debug().Table(INDICATORABLENAME).Create(&indicator)
	if res1.Error != nil {
		log.Fatalln(res1.Error)
		return false
	}

	return true
}
func (i IndicatorServiceImpl) Delete(indicatorCode string) bool {
	res := dao.DB.Debug().Table(INDICATORABLENAME).Where("code = ? ", indicatorCode).Where("is_delete = ?", 0).Update("is_delete", 1)
	if res.Error != nil {
		log.Fatalln(res.Error)
		return false
	}

	return true
}
func (i IndicatorServiceImpl) Query(indicatorCode string) vo.IndicatorVO {
	indicators := indicator_dao.Indicator{}
	res := dao.DB.Debug().Table(INDICATORABLENAME).Where("code = ?", indicatorCode).Where("is_delete = ?", 0).Find(&indicators)
	if res.Error != nil {
		log.Fatalln(res.Error)
	}
	if res.RowsAffected > 1 {
		log.Fatalln("code is not only")
	} else if res.RowsAffected == 0 {
		log.Print("not found")
		//TODO: return what?
		return vo.IndicatorVO{}
	}
	fmt.Println(res.RowsAffected)
	//fmt.Println(indicators.Expression)
	if indicators.Expression == "" {
		log.Fatalln("empty expression")
	}
	//construct VO
	indicatorJson := i.AntiSerialization(indicators.Expression)
	indicatorVo := vo.IndicatorVO{IndicatorCode: indicators.IndicatorCode,
		Name:        indicators.Name,
		Indicators:  indicatorJson.Indicators,
		Caculate:    indicatorJson.Caculate,
		Value:       indicatorJson.Value,
		Description: indicators.Description,
		CreateTime:  indicators.CreateTime,
		UpdateTime:  indicators.UpdateTime}

	fmt.Println(indicatorVo)
	return indicatorVo
}
func (i IndicatorServiceImpl) Modify(indicator indicator_dao.Indicator) bool {
	if !IsIndicatorExist(indicator.IndicatorCode) {
		log.Fatalln("the indicator code not exist")
		return false
	}
	res := dao.DB.Debug().Omit("create_time").Where("code", indicator.IndicatorCode).Where("is_delete = ?", 0).Save(&indicator)
	if res.Error != nil {
		log.Fatalln(res.Error)
	}
	return true
}

func IsIndicatorExist(indicatorCode string) bool {
	var count int64
	res := dao.DB.Debug().Table(INDICATORABLENAME).Where("code = ?", indicatorCode).Where("is_delete = ?", 0).Count(&count)
	if res.Error != nil {
		log.Fatalln(res.Error)
		return false
	}
	if count != 0 {

		return true
	}
	return false
}

//TODO: param condition
func (i IndComputeImpl) Compute(indicatorCode string, condition vo.Condition) uint {
	ind := IndicatorServiceImpl{}.Query(indicatorCode)
	return i.ComputeLeaf(ind, condition)
}

func (i IndComputeImpl) ComputeLeaf(ind vo.IndicatorVO, condition vo.Condition) uint {
	sql := ind.Value
	if ind.Indicators == nil && len(sql) != 0 && ind.Caculate == -1 {
		//leaf
		sql += ConstructSql(condition)
		var amount uint
		res := dao.DB.Debug().Raw(sql).Scan(&amount)
		if res.Error != nil {
			log.Fatalln(res.Error)
		}
		return amount
	} else {
		switch ind.Caculate {
		case core.ADD:
			value := i.ComputeLeaf(ind.Indicators[0], condition) + i.ComputeLeaf(ind.Indicators[1], condition)
			return value
			break
		case core.SUBTRACT:
			value := i.ComputeLeaf(ind.Indicators[0], condition) - i.ComputeLeaf(ind.Indicators[1], condition)
			return value
			break
		case core.MULTIPLY:
			value := i.ComputeLeaf(ind.Indicators[0], condition) * i.ComputeLeaf(ind.Indicators[1], condition)
			return value
			break
		case core.DIVIDE:
			value := i.ComputeLeaf(ind.Indicators[0], condition) / i.ComputeLeaf(ind.Indicators[1], condition)
			return value
			break
		}
	}
	return 0
}

func ConstructSql(condition vo.Condition) string {
	roomIdCondition := " where room_id = " + strconv.Itoa(int(condition.RoomID))
	conditionAll := roomIdCondition
	var startTimeCondition string
	var endTimeCondition string
	if condition.StartTime != "" {
		starTime := condition.StartTime
		startTimeCondition = " and deal_time >= " + starTime
		conditionAll = conditionAll + startTimeCondition
	}
	if condition.EndTime != "" {
		endTime := condition.EndTime
		endTimeCondition = " and deal_time <= " + endTime
		conditionAll = conditionAll + endTimeCondition
	}
	return conditionAll
}
