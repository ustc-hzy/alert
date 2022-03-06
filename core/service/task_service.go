package service

import (
	"alert/core/dao/task_dao"
	"alert/core/vo"
	"time"
)

type TaskServiceImpl struct{}
type TaskScheduleImpl struct{}

func (i TaskServiceImpl) Add(TaskCode string, TaskName string, RuleCode string, Frequency time.Duration, NextTime time.Time, Status bool) bool {
	task := task_dao.Task{TaskCode: TaskCode, TaskName: TaskName, RuleCode: RuleCode, Frequency: Frequency, NextTime: NextTime, Status: Status}
	res := DB.Debug().Create(task)
	if res.Error != nil {
		return false
	}
	return true
}

func (i TaskServiceImpl) Delete(TaskCode string) bool {
	res := DB.Debug().Delete(&task_dao.Task{}, TaskCode)
	if res.Error != nil {
		return false
	}
	return true
}

func (i TaskServiceImpl) Query(TaskCode string) task_dao.Task {
	var task task_dao.Task
	DB.Where("TaskCode = ?", TaskCode).First(task)
	return task
}

func (i TaskServiceImpl) Modify(TaskCode string, TaskName string, RuleCode string, Frequency time.Duration, NextTime time.Time, Status bool) bool {
	task := task_dao.Task{TaskCode: TaskCode, TaskName: TaskName, RuleCode: RuleCode, Frequency: Frequency, NextTime: NextTime, Status: Status}
	res := DB.Debug().Save(task)
	if res.Error != nil {
		return false
	}
	return true
}

func (i TaskServiceImpl) UpdateTime(vo vo.TaskVO) bool {
	vo.NextTime = vo.NextTime.Add(vo.Frequency)
	return true
}

func (i TaskScheduleImpl) Schedule(Frequency time.Duration, TaskList []vo.TaskVO) {
	for {
		for j := 0; j < len(TaskList); j++ {
			WorkServiceImpl{}.Work(TaskList[j].RuleCode)
			TaskServiceImpl{}.UpdateTime(TaskList[j])
			time.Sleep(Frequency)
		}
	}
}
