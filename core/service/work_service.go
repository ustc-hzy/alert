package service

import "time"

type WorkServiceImpl struct{}

func (i WorkServiceImpl) Work(RuleCode string) bool {
	return RuleCheckImpl{}.Check(RuleCode)
}

func (i WorkServiceImpl) ComputeWork(IndicatorCode string, RoomID uint, StartTime time.Time, EndTime time.Time) uint {
	return IndComputeImpl{}.Compute(IndicatorCode, RoomID, StartTime, EndTime)
}
