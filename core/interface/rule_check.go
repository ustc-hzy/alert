package _interface

import "alert/core/dto"

type RuleCheckInterface interface {
	Check(RuleCode string) (bool, error)
	CheckLeaf(rule dto.RuleVo) (bool, error)
}
