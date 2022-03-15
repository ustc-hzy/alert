package service

import (
	"alert/core/vo"
	"fmt"
	"log"
	"time"
)

type WorkServiceImpl struct{}

func (i WorkServiceImpl) Work(ruleCode string, ch chan int) {
	ret, err := RuleCheckImpl{}.Check(ruleCode)
	if err != nil {
		log.Println(err)
	}
	if ret == true {
		res := RuleServiceImpl{}.Query(ruleCode)
		AlertServiceImpl{}.Add(res.RuleName, res.RoomId, time.Now())
		fmt.Println(ruleCode + " has been triggered")
	}
	<-ch
}

func (i WorkServiceImpl) ComputeWork(IndicatorCode string, condition vo.Condition) uint {
	return IndComputeImpl{}.Compute(IndicatorCode, condition)
}
