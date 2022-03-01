package dto

import "time"

type TaskVO struct {
	TaskCode  string
	RuleCode  string
	NextTime  time.Time
	Frequency time.Duration
	Status    bool
}
