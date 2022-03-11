package service

import (
	"alert/core/vo"
	"fmt"
	"log"
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

func (i WorkServiceImpl) ComputeWork(IndicatorCode string, condition vo.Condition) uint {
	return IndComputeImpl{}.Compute(IndicatorCode, condition)
}
