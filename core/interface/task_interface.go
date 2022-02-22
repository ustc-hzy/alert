package _interface

import (
	"alert/core/dao/task_dao"
	"time"
)

type TaskInterface interface {
	Add(TaskCode string, TaskName string, RuleCode string, Frequency time.Duration, NextTime time.Time, Status bool) bool
	Delete(TaskCode string) bool
	Query(TaskCode string) task_dao.Task
	Modify(TaskCode string, TaskName string, RuleCode string, Frequency time.Duration, NextTime time.Time, Status bool) bool
	UpdateTime() bool
}
