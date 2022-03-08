package service

import (
	"fmt"
	"log"
	"time"
)

type WorkServiceImpl struct{}

func (i WorkServiceImpl) Work(ruleCode string, ch chan int) {
	ret, err := RuleCheckImpl{}.Check(ruleCode)
	if err != nil {
		log.Fatal(err)
	}
	if ret == true {
		//TODO Alert
		fmt.Println(ruleCode + " has been triggered")
	}
	<-ch
}

func (i WorkServiceImpl) ComputeWork(IndicatorCode string, RoomID uint, StartTime, EndTime time.Time) uint {
	return IndComputeImpl{}.Compute(IndicatorCode, RoomID, StartTime, EndTime)
}
