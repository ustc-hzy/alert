package service

import (
	"alert/core/dao"
	"alert/core/dao/task_dao"
	"alert/core/vo"
	"fmt"
	"log"
	"time"
)

type TaskServiceImpl struct{}
type ScheduleImpl struct{}

const (
	TASKTABLENAME = "tasks"
	POOLSIZE      = 10
)

func (i TaskServiceImpl) Add(task task_dao.Task) bool {
	var count int64
	res := dao.DB.Debug().Table(TASKTABLENAME).Where("task_code = ?", task.TaskCode).Count(&count)
	if res.Error != nil {
		log.Println(res.Error)
		return false
	}
	if count != 0 {
		log.Println("the task already exists")
		return false
	}

	//if rules don't exist
	m := RuleServiceImpl{}.IsRuleExist(task.RuleCode)
	if m == false {
		log.Println("the rule doesn't exist")
		return false
	}

	resp := dao.DB.Debug().Table(TASKTABLENAME).Create(&task)
	if resp.Error != nil {
		log.Println(resp.Error)
		return false
	}
	return true
}

func (i TaskServiceImpl) Delete(taskCode string) bool {
	res := dao.DB.Debug().Table(TASKTABLENAME).Where("task_code = ? ", taskCode).Delete(&task_dao.Task{})
	if res.Error != nil {
		log.Println(res.Error)
		return false
	}
	return true
}

func (i TaskServiceImpl) Query(taskCode string) vo.TaskVO {
	task := task_dao.Task{}
	res := dao.DB.Debug().Table(TASKTABLENAME).Where("task_code = ?", taskCode).Find(&task)
	if res.Error != nil {
		log.Println(res.Error)
	}
	if res.RowsAffected > 1 {
		log.Println("task duplicated")
	} else if res.RowsAffected == 0 {
		log.Print("task not found")
		return vo.TaskVO{}
	}
	taskVo := TaskServiceImpl{}.TransferTaskVo(task)

	fmt.Println(taskVo)
	return taskVo
}

func (i TaskServiceImpl) Modify(task task_dao.Task) bool {
	//if rules don't exist
	m := RuleServiceImpl{}.IsRuleExist(task.RuleCode)
	if m == false {
		log.Println("the rule doesn't exist")
		return false
	}

	res := dao.DB.Debug().Omit("next_time", "status").Where("task_code", task.TaskCode).Save(&task)
	if res.Error != nil {
		log.Println(res.Error)
	}
	return true
}

func (i TaskServiceImpl) UpdateTime(task task_dao.Task) bool {
	res := dao.DB.Debug().Model(&task).Update("next_time", time.Now().Add(task.Frequency))
	if res.Error != nil {
		log.Println(res.Error)
	}
	return true
}

func (i TaskServiceImpl) UpdateStatus(task task_dao.Task, status bool) bool {
	res := dao.DB.Debug().Model(&task).Update("status", status)
	if res.Error != nil {
		log.Println(res.Error)
	}
	return true
}

func (i TaskServiceImpl) TransferTaskVo(task task_dao.Task) vo.TaskVO {
	taskVo := vo.TaskVO{
		TaskCode:  task.TaskCode,
		RuleCode:  task.RuleCode,
		NextTime:  task.NextTime,
		Frequency: task.Frequency,
		Status:    task.Status,
	}
	return taskVo
}

func (s ScheduleImpl) Schedule(frequency time.Duration) {

	ch := make(chan string, POOLSIZE)
	//schedule
	for i := 1; ; i++ {
		//get taskList
		var taskList []task_dao.Task
		res := dao.DB.Debug().Table(TASKTABLENAME).Find(&taskList)
		if res.Error != nil {
			log.Println(res.Error)
		}

		//check
		for _, task := range taskList {
			if task.NextTime.Before(time.Now()) {
				ch <- task.TaskName
				go WorkServiceImpl{}.Work(task.RuleCode, ch)
				TaskServiceImpl{}.UpdateTime(task)
			}
		}

		//sleep
		time.Sleep(frequency)
	}

}
