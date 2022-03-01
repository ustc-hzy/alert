package service

import (
	"alert/core/dao"
	"alert/core/dao/indicator_dao"
	"alert/core/dto"
	"time"
)

var DB = dao.InitDB()

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
func (i IndicatorServiceImpl) Query(IndicatorCode string) dto.IndicatorVO {
	return dto.IndicatorVO{}
}
func (i IndicatorServiceImpl) Modify(indicator indicator_dao.Indicator) bool {
	return true
}

func (i IndComputeImpl) Compute(IndicatorCode string, RoomID uint, StartTime time.Time, EndTime time.Time) uint {
	ind := IndicatorServiceImpl{}.Query(IndicatorCode)
	return i.ComputeLeaf(ind, RoomID, StartTime, EndTime)
}

func (i IndComputeImpl) ComputeLeaf(ind dto.IndicatorVO, id uint, start time.Time, end time.Time) uint {
	if ind.Indicators == nil && len(ind.Value) != 0 && ind.Op == -1 {
		var ret uint
		DB.Raw(ind.Value).Scan(&ret)
	} else if ind.Indicators != nil && len(ind.Value) == 0 && ind.Op != -1 {
		i1 := i.ComputeLeaf(ind.Indicators[0], id, start, end)
		i2 := i.ComputeLeaf(ind.Indicators[1], id, start, end)
		switch ind.Op {
		case 0:
			return i1 + i2
		case 1:
			return i1 - i2
		case 2:
			return i1 * i2
		case 3:
			return i1 / i2
		}
	}
	return 0
}
