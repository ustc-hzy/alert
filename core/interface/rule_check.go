package _interface

import "alert/core/vo"

type RuleCheckInterface interface {
	Check(ruleCode string) (bool, error)
	CheckLeaf(rule vo.RuleVo) (bool, error)
}
