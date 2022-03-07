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
	// TODO: Your code here...
	var TaskList []task_dao.Task
	res := DB.Debug().Table(TASKTABLENAME).Find(&TaskList)
	if res.Error != nil {
		log.Fatalln(res.Error)
	}

	for i := 1; ; i++ {
		for j := range TaskList {
			go WorkServiceImpl{}.Work(TaskList[j].RuleCode)
			TaskServiceImpl{}.UpdateTime(TaskList[j])
		}
		if i&256 == 0 {
			for k := range TaskList {
				TaskServiceImpl{}.Modify(TaskList[k])
			}
		}
		time.Sleep(time.Duration(req.Frequency))
	}
	return &api.ScheduleResponse{Success: true}, nil
}
