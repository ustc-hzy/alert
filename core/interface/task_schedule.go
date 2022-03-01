package _interface

import "time"

type TaskScheduleInterface interface {
	Schedule(Frequency time.Duration)
}
