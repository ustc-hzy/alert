package service

import (
	"log"
	"time"
)

type WorkServiceImpl struct{}

func (i WorkServiceImpl) Work(RuleCode string) bool {
	ret, err := RuleCheckImpl{}.Check(RuleCode)
	if err != nil {
		log.Fatal(err)
	}
	return ret
}

func (i WorkServiceImpl) ComputeWork(IndicatorCode string, RoomID uint, StartTime time.Time, EndTime time.Time) uint {
	return IndComputeImpl{}.Compute(IndicatorCode, RoomID, StartTime, EndTime)
}
