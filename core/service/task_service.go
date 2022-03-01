package service

import (
	"alert/core/dao/task_dao"
	"alert/core/dto"
	"time"
)

type TaskServiceImpl struct{}
type TaskScheduleImpl struct{}

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

func (i TaskServiceImpl) UpdateTime(vo dto.TaskVO) bool {
	return false
}

func (i TaskScheduleImpl) Schedule(Frequency time.Duration, TaskList []dto.TaskVO) {
	for {
		for j := 0; j < len(TaskList); j++ {
			WorkServiceImpl{}.Work(TaskList[j].RuleCode)
			TaskServiceImpl{}.UpdateTime(TaskList[j])
			time.Sleep(1)
		}
	}
}
