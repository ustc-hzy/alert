package task_dao

import "time"

type Task struct {
	TaskCode  string
	TaskName  string
	RuleCode  string
	Frequency time.Duration
	NextTime  time.Time
	Status    bool
}
