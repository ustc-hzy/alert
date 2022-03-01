package service

import (
	"alert/core/dao/indicator_dao"
	"alert/core/dto"
	"gorm.io/gorm"
	"time"
)

type IndicatorServiceImpl struct{}
type IndComputeImpl struct{}

var DB *gorm.DB

func (i IndicatorServiceImpl) Serialization(Indicators []indicator_dao.Indicator, Op indicator_dao.Op_type, Value string) string {

}
func (i IndicatorServiceImpl) AntiSerialization(Expression string) indicator_dao.IndicatorJson {

}
func (i IndicatorServiceImpl) Add(IndicatorCode string, Name string, Expression string, Description string,
	CreateTime time.Time, UpdateTime time.Time) bool {

}
func (i IndicatorServiceImpl) Delete(IndicatorCode string) bool {

}
func (i IndicatorServiceImpl) Query(IndicatorCode string) dto.IndicatorVO {

}
func (i IndicatorServiceImpl) Modify(IndicatorCode string, Name string, Expression string, Description string,
	CreateTime time.Time, UpdateTime time.Time) bool {

}

func (i IndComputeImpl) Compute(IndicatorCode string, RoomID uint, StartTime time.Time, EndTime time.Time) uint {
	ind := IndicatorServiceImpl{}.Query(IndicatorCode)
	return i.ComputeLeaf(ind)
}

func (i IndComputeImpl) ComputeLeaf(ind dto.IndicatorVO) uint {
	if ind.Indicators == nil && len(ind.Value) != 0 && ind.Op == -1 {

	}
}
