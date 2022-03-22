package service

import (
	"alert/core/vo"
	"fmt"
	"log"
	"time"
)

type WorkServiceImpl struct{}

func (i WorkServiceImpl) Work(ruleCode string, ch chan string) {
	taskName := <-ch
	ret, err := RuleCheckImpl{}.Check(ruleCode)
	if err != nil {
		log.Println(err)
	}
	if ret == true {
		res := RuleServiceImpl{}.Query(ruleCode)
		expression, err := RuleServiceImpl{}.RuleExpression(res)
		if err != nil {
			log.Println(err)
		}
		AlertServiceImpl{}.Add(taskName, res.RuleName, expression, res.RoomId, time.Now())
		fmt.Println(ruleCode + " has been triggered")
	}

}

func (i WorkServiceImpl) ComputeWork(IndicatorCode string, condition vo.Condition) float64 {
	return IndComputeImpl{}.Compute(IndicatorCode, condition)
}
