package _interface

import (
	"alert/core/vo"
)

type WorkInterface interface {
	Work(RuleCode string, ch chan int)
	ComputeWork(IndicatorCode string, condition vo.Condition) uint
}
