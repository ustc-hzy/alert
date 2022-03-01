package service

import (
	"alert/core/dao"
	"alert/core/dao/indicator_dao"
	"alert/core/dto"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB = dao.InitDB()

type IndicatorServiceImpl struct{}
type IndComputeImpl struct{}

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

//TODO: keep only one
func (i IndicatorServiceImpl) Add(indicator indicator_dao.Indicator, indicatorJson indicator_dao.IndicatorJson) bool {

	indicator.Expression = i.Serialization(indicatorJson)

	res := DB.Debug().Create(&indicator)
	if res.Error != nil {
		log.Fatalln(res.Error)
		return false
	}

	return true
}
func (i IndicatorServiceImpl) Delete(IndicatorCode string) bool {
	res := DB.Debug().Table("indicators").Where("code = ? ", IndicatorCode).Update("is_delete", 1)
	if res.Error != nil {
		log.Fatalln(res.Error)
		return false
	}

	return true
}
func (i IndicatorServiceImpl) Query(IndicatorCode string) dto.IndicatorVO {
	indicators := indicator_dao.Indicator{}
	res := DB.Debug().Table("indicators").Where("code = ?", IndicatorCode).Where("is_delete = ?", 0).Find(&indicators)
	if res.Error != nil {
		log.Fatalln(res.Error)
	}
	if res.RowsAffected > 1 {
		log.Fatalln("code is not only")
	}
	fmt.Println(res.RowsAffected)
	fmt.Println(indicators.Expression)
	//construct VO
	indicatorJson := i.AntiSerialization(indicators.Expression)
	indicatorVo := dto.IndicatorVO{IndicatorCode: indicators.IndicatorCode,
		Name:        indicators.Name,
		Indicators:  indicatorJson.Indicators,
		Op:          dto.Op_type(indicatorJson.Op),
		Value:       indicatorJson.Value,
		Description: indicators.Description,
		CreateTime:  indicators.CreateTime,
		UpdateTime:  indicators.UpdateTime}

	fmt.Println(indicatorVo)
	return indicatorVo
}
func (i IndicatorServiceImpl) Modify(indicator indicator_dao.Indicator) bool {
	res := DB.Debug().Omit("create_time").Where("code", indicator.IndicatorCode).Save(&indicator)
	if res.Error != nil {
		log.Fatalln(res.Error)
	}
	return true
}

//TODO: param condition
func (i IndComputeImpl) Compute(IndicatorCode string, RoomID uint, StartTime time.Time, EndTime time.Time) uint {
	ind := IndicatorServiceImpl{}.Query(IndicatorCode)
	return i.ComputeLeaf(ind)
}

func (i IndComputeImpl) ComputeLeaf(ind dto.IndicatorVO) uint {
	if ind.Indicators == nil && len(ind.Value) != 0 && ind.Op == -1 {
		//leaf
		var amount uint
		res := DB.Debug().Table("deal_infos").Exec(ind.Value).First(&amount)
		if res.Error != nil {
			log.Fatalln(res.Error)
		}
		return amount
	} else {
		switch ind.Op {
		case dto.Op_type(indicator_dao.ADD):
			value := i.ComputeLeaf(ind.Indicators[0]) + i.ComputeLeaf(ind.Indicators[1])
			return value
			break
		case dto.Op_type(indicator_dao.SUBTRACT):
			value := i.ComputeLeaf(ind.Indicators[0]) - i.ComputeLeaf(ind.Indicators[1])
			return value
			break
		case dto.Op_type(indicator_dao.MULTIPLY):
			value := i.ComputeLeaf(ind.Indicators[0]) * i.ComputeLeaf(ind.Indicators[1])
			return value
			break
		case dto.Op_type(indicator_dao.DIVIDE):
			value := i.ComputeLeaf(ind.Indicators[0]) / i.ComputeLeaf(ind.Indicators[1])
			return value
			break
		}
	}
	return 0
}
