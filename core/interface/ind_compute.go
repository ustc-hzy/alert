package _interface

import (
	"alert/core/vo"
)

type IndComputeInterface interface {
	Compute(IndicatorCode string, condition vo.Condition) uint
	ComputeLeaf(vo vo.IndicatorVO, condition vo.Condition) uint
}
