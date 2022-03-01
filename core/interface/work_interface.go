package _interface

import "time"

type WorkInterface interface {
	Work(RuleCode string) bool
	ComputeWork(IndicatorCode string, RoomID uint, StartTime time.Time, EndTime time.Time) uint
}
