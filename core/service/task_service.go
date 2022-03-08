package service

import (
	"alert/core/dao/task_dao"
	"alert/core/vo"
	"alert/kitex_gen/api"
	"context"
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

func (i TaskServiceImpl) Delete(taskCode string) bool {
	res := DB.Debug().Table(TASKTABLENAME).Where("task_code = ? ", taskCode).Delete(&task_dao.Task{})
	if res.Error != nil {
		log.Fatalln(res.Error)
		return false
	}
	return true
}

func (i TaskServiceImpl) Query(taskCode string) vo.TaskVO {
	task := task_dao.Task{}
	res := DB.Debug().Table(TASKTABLENAME).Where("task_code = ?", taskCode).Find(&task)
	if res.Error != nil {
		log.Fatalln(res.Error)
	}
	if res.RowsAffected > 1 {
		log.Fatalln("task duplicated")
	} else if res.RowsAffected == 0 {
		log.Print("task not found")
		return vo.TaskVO{}
	}
	taskVo := TaskServiceImpl{}.TransferTaskVo(task)

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

func (i TaskServiceImpl) UpdateTime(task task_dao.Task) bool {
	task.NextTime = task.NextTime.Add(task.Frequency)
	return true
}

func (i TaskServiceImpl) UpdateStatus(task task_dao.Task, status bool) bool {
	task.Status = status
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

// Schedule implements the ScheduleImpl interface.
func (s *ScheduleImpl) Schedule(ctx context.Context, req *api.ScheduleRequest) (resp *api.ScheduleResponse, err error) {
	//get taskList
	var taskList []task_dao.Task
	res := DB.Debug().Table(TASKTABLENAME).Find(&taskList)
	if res.Error != nil {
		log.Fatalln(res.Error)
	}

	ch := make(chan int, POOLSIZE)
	//schedule
	for i := 1; ; i++ {
		for j, task := range taskList {
			ch <- j
			go WorkServiceImpl{}.Work(task.RuleCode, ch)
			TaskServiceImpl{}.UpdateTime(taskList[j])
		}
		//write back to DB
		if i&256 == 0 {
			for _, m := range taskList {
				TaskServiceImpl{}.Modify(m)
			}
		}
		//sleep
		time.Sleep(time.Duration(req.Frequency))
	}
	//return &api.ScheduleResponse{Success: true}, nil
}
