package service

import (
	"alert/core/dao"
	"alert/core/dao/indicator_dao"
	"gorm.io/gorm"
	"alert/core/dto"
	"time"
)

var DB *gorm.DB = dao.InitDB()

type IndicatorServiceImpl struct{}
type IndComputeImpl struct{}

func (i IndicatorServiceImpl) Serialization(indicatorJson indicator_dao.IndicatorJson) string {
	return ""
}
func (i IndicatorServiceImpl) AntiSerialization(expression string) indicator_dao.IndicatorJson {
	return indicator_dao.IndicatorJson{}
}
func (i IndicatorServiceImpl) IndicatorAdd(indicator indicator_dao.Indicator) bool {
	res := DB.Debug().Create(&indicator)
	if res.Error != nil {
		return false
	}

	return true
}
func (i IndicatorServiceImpl) Delete(IndicatorCode string) bool {
	return true
}
func (i IndicatorServiceImpl) Query(IndicatorCode string) indicator_dao.Indicator {
	return indicator_dao.Indicator{}
}
func (i IndicatorServiceImpl) Modify(indicator indicator_dao.Indicator) bool {
	return true
}

func (i IndComputeImpl) Compute(IndicatorCode string, RoomID uint, StartTime time.Time, EndTime time.Time) uint {
	ind := IndicatorServiceImpl{}.Query(IndicatorCode)
	return i.ComputeLeaf(ind)
}

func (i IndComputeImpl) ComputeLeaf(ind dto.IndicatorVO) uint {
	if ind.Indicators == nil && len(ind.Value) != 0 && ind.Op == -1 {

	}
}
