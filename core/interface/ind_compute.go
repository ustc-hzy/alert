package _interface

import (
	"alert/core/dto"
	"time"
)

type IndComputeInterface interface {
	Compute(IndicatorCode string, RoomID uint, StartTime time.Time, EndTime time.Time) uint
	ComputeLeaf(vo dto.IndicatorVO) uint
}
