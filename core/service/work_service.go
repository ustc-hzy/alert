package service

import (
	"fmt"
	"log"
	"time"
)

type WorkServiceImpl struct{}

func (i WorkServiceImpl) Work(RuleCode string) {
	ret, err := RuleCheckImpl{}.Check(RuleCode)
	if err != nil {
		log.Fatal(err)
	}
	if ret == true {
		//TODO Alert
		fmt.Println(RuleCode + " has been triggered")
	}
}

func (i WorkServiceImpl) ComputeWork(IndicatorCode string, RoomID uint, StartTime time.Time, EndTime time.Time) uint {
	return IndComputeImpl{}.Compute(IndicatorCode, RoomID, StartTime, EndTime)
}
