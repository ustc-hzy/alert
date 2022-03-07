package service

import (
	"alert/core/dao/task_dao"
	"alert/core/vo"
	"fmt"
	"log"
	"time"
)

type TaskServiceImpl struct{}
type TaskScheduleImpl struct{}

const TASKTABLENAME = "tasks"

func (i TaskServiceImpl) Add(task task_dao.Task) bool {
	var count int64
	res := DB.Debug().Table(TASKTABLENAME).Where("task_code = ?", task.TaskCode).Count(&count)
	if res.Error != nil {
		log.Fatalln(res.Error)
		return false
	}
	if count != 0 {
		log.Fatalln("the task already exists")
		return false
	}
	resp := DB.Debug().Table(TASKTABLENAME).Create(&task)
	if resp.Error != nil {
		log.Fatalln(resp.Error)
		return false
	}
	return true
}

func (i TaskServiceImpl) Delete(TaskCode string) bool {
	res := DB.Debug().Table(TASKTABLENAME).Where("task_code = ? ", TaskCode).Delete(&task_dao.Task{})
	if res.Error != nil {
		log.Fatalln(res.Error)
		return false
	}
	return true
}

func (i TaskServiceImpl) Query(TaskCode string) vo.TaskVO {
	task := task_dao.Task{}
	res := DB.Debug().Table(TASKTABLENAME).Where("task_code = ?", TaskCode).Find(&task)
	if res.Error != nil {
		log.Fatalln(res.Error)
	}
	if res.RowsAffected > 1 {
		log.Fatalln("task duplicated")
	} else if res.RowsAffected == 0 {
		log.Print("task not found")
		return vo.TaskVO{}
	}
	taskVo := vo.TaskVO{
		TaskCode:  task.TaskCode,
		RuleCode:  task.RuleCode,
		NextTime:  task.NextTime,
		Frequency: task.Frequency,
		Status:    task.Status,
	}

	fmt.Println(taskVo)
	return taskVo
}

func (i TaskServiceImpl) Modify(task task_dao.Task) bool {
	res := DB.Debug().Omit("next_time", "status").Where("task_code", task.TaskCode).Save(&task)
	if res.Error != nil {
		log.Fatalln(res.Error)
	}
	return true
}

func (i TaskServiceImpl) UpdateTime(vo vo.TaskVO) bool {
	vo.NextTime = vo.NextTime.Add(vo.Frequency)
	return true
}

func (i TaskServiceImpl) UpdateStatus(vo vo.TaskVO, status bool) bool {
	vo.Status = status
	return true
}

func (i TaskScheduleImpl) Schedule(Frequency time.Duration, TaskList []vo.TaskVO) {
	for {
		for j := 0; j < len(TaskList); j++ {
			go WorkServiceImpl{}.Work(TaskList[j].RuleCode)
			go TaskServiceImpl{}.UpdateTime(TaskList[j])
			time.Sleep(Frequency)
		}
	}
}
