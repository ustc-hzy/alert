package service

import (
	"alert/core/dao/task_dao"
	"alert/core/dto"
	"time"
)

type TaskServiceImpl struct{}
type TaskScheduleImpl struct{}

func (i TaskServiceImpl) Add(task task_dao.Task) bool {
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

func (i TaskServiceImpl) Query(TaskCode string) dto.TaskVO {
	var task task_dao.Task
	DB.Debug().Where("task_code = ?", TaskCode).First(task)
	TaskVo := i.Transfer(task)
	return TaskVo
}

func (i TaskServiceImpl) Modify(TaskCode, TaskName, RuleCode string, Frequency time.Duration, NextTime time.Time, Status bool) bool {
	task := task_dao.Task{TaskCode: TaskCode, TaskName: TaskName, RuleCode: RuleCode, Frequency: Frequency, NextTime: NextTime, Status: Status}
	res := DB.Debug().Save(task)
	if res.Error != nil {
		return false
	}
	return true
}

func (i TaskServiceImpl) UpdateTime(vo dto.TaskVO) bool {
	vo.NextTime = vo.NextTime.Add(vo.Frequency)
	return true
}

func (i TaskServiceImpl) Transfer(task task_dao.Task) dto.TaskVO {
	TaskVo := dto.TaskVO{TaskCode: task.TaskCode, RuleCode: task.RuleCode, NextTime: task.NextTime, Frequency: task.Frequency, Status: task.Status}
	return TaskVo
}

func (i TaskScheduleImpl) Schedule(Frequency time.Duration, TaskList []dto.TaskVO) {
	for {
		for j := 0; j < len(TaskList); j++ {
			WorkServiceImpl{}.Work(TaskList[j].RuleCode)
			TaskServiceImpl{}.UpdateTime(TaskList[j])
			time.Sleep(Frequency)
		}
	}
}
