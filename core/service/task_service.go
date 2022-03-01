package service

import (
	"alert/core/dao/task_dao"
	"alert/core/dto"
	"gorm.io/gorm"
	"time"
)

type TaskServiceImpl struct{}
type TaskScheduleImpl struct{}

var DB *gorm.DB

func (i TaskServiceImpl) Add(TaskCode string, TaskName string, RuleCode string, Frequency time.Duration, NextTime time.Time, Status bool) bool {
	task := task_dao.Task{TaskCode: TaskCode, TaskName: TaskName, RuleCode: RuleCode, Frequency: Frequency, NextTime: NextTime, Status: Status}
	DB.Create(task)
	return true
}

func (i TaskServiceImpl) Delete(TaskCode string) bool {
	DB.Delete(&task_dao.Task{}, TaskCode)
	return true
}

func (i TaskServiceImpl) Query(TaskCode string) task_dao.Task {
	var task task_dao.Task
	DB.Where("TaskCode = ?", TaskCode).First(task)
	return task
}

func (i TaskServiceImpl) Modify(TaskCode string, TaskName string, RuleCode string, Frequency time.Duration, NextTime time.Time, Status bool) bool {
	task := task_dao.Task{TaskCode: TaskCode, TaskName: TaskName, RuleCode: RuleCode, Frequency: Frequency, NextTime: NextTime, Status: Status}
	DB.Save(task)
	return true
}

func (i TaskServiceImpl) UpdateTime(vo dto.TaskVO) bool {
	vo.NextTime = vo.NextTime.Add(vo.Frequency)
	return true
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
