package service

import (
	"alert/core/dao/task_dao"
	"time"
)

type TaskServiceImpl struct{}

func (i TaskServiceImpl) Add(TaskCode string, TaskName string, RuleCode string, Frequency time.Duration, NextTime time.Time, Status bool) bool {
	return false
}

func (i TaskServiceImpl) Delete(TaskCode string) bool {
	return false
}

func (i TaskServiceImpl) Query(TaskCode string) task_dao.Task {
	return task_dao.Task{}
}

func (i TaskServiceImpl) Modify(TaskCode string, TaskName string, RuleCode string, Frequency time.Duration, NextTime time.Time, Status bool) bool {
	return false
}

func (i TaskServiceImpl) UpdateTime() bool {
	return false
}
