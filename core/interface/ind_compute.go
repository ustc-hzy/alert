package _interface

import (
	"alert/core/vo"
	"time"
)

type IndComputeInterface interface {
	Compute(IndicatorCode string, RoomID uint, StartTime time.Time, EndTime time.Time) uint
	ComputeLeaf(vo vo.IndicatorVO) uint
}
