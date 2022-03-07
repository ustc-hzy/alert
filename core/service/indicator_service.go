package service

import (
	"alert/core"
	"alert/core/dao"
	"alert/core/dao/indicator_dao"
	"alert/core/vo"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

var DB = dao.InitDB()

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
	var count int64
	res := DB.Debug().Table(INDICATORABLENAME).Where("code = ?", indicator.IndicatorCode).Where("is_delete = ?", 0).Count(&count)
	if res.Error != nil {
		log.Fatalln(res.Error)
		return false
	}
	if count != 0 {
		log.Fatalln("the indicator code already exist")
		return false
	}

	res1 := DB.Debug().Table(INDICATORABLENAME).Create(&indicator)
	if res1.Error != nil {
		log.Fatalln(res1.Error)
		return false
	}

	return true
}
func (i IndicatorServiceImpl) Delete(indicatorCode string) bool {
	res := DB.Debug().Table(INDICATORABLENAME).Where("code = ? ", indicatorCode).Where("is_delete = ?", 0).Update("is_delete", 1)
	if res.Error != nil {
		log.Fatalln(res.Error)
		return false
	}

	return true
}
func (i IndicatorServiceImpl) Query(indicatorCode string) vo.IndicatorVO {
	indicators := indicator_dao.Indicator{}
	res := DB.Debug().Table(INDICATORABLENAME).Where("code = ?", indicatorCode).Where("is_delete = ?", 0).Find(&indicators)
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
	res := DB.Debug().Omit("create_time").Where("code", indicator.IndicatorCode).Where("is_delete = ?", 0).Save(&indicator)
	if res.Error != nil {
		log.Fatalln(res.Error)
	}
	return true
}

//TODO: param condition
func (i IndComputeImpl) Compute(indicatorCode string, roomID uint, startTime time.Time, endTime time.Time) uint {
	ind := IndicatorServiceImpl{}.Query(indicatorCode)
	return i.ComputeLeaf(ind)
}

func (i IndComputeImpl) ComputeLeaf(ind vo.IndicatorVO) uint {
	if ind.Indicators == nil && len(ind.Value) != 0 && ind.Caculate == -1 {
		//leaf
		var amount uint
		res := DB.Debug().Raw(ind.Value).Scan(&amount)
		if res.Error != nil {
			log.Fatalln(res.Error)
		}
		return amount
	} else {
		switch ind.Caculate {
		case core.ADD:
			value := i.ComputeLeaf(ind.Indicators[0]) + i.ComputeLeaf(ind.Indicators[1])
			return value
			break
		case core.SUBTRACT:
			value := i.ComputeLeaf(ind.Indicators[0]) - i.ComputeLeaf(ind.Indicators[1])
			return value
			break
		case core.MULTIPLY:
			value := i.ComputeLeaf(ind.Indicators[0]) * i.ComputeLeaf(ind.Indicators[1])
			return value
			break
		case core.DIVIDE:
			value := i.ComputeLeaf(ind.Indicators[0]) / i.ComputeLeaf(ind.Indicators[1])
			return value
			break
		}
	}
	return 0
}
