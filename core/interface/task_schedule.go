package _interface

import "time"

type ScheduleInterface interface {
	Schedule(Frequency time.Duration)
}
