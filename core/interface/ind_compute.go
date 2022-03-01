package _interface

import "time"

type IndComputeInterface interface {
	Compute(IndicatorCode string, RoomID uint, StartTime time.Time, EndTime time.Time) uint
}
