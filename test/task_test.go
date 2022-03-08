package test

import (
	"alert/core/dao/task_dao"
	"alert/core/service"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestTaskAdd(*testing.T) {
	task := task_dao.Task{
		TaskCode:  "taskTest",
		TaskName:  "sum",
		RuleCode:  "test12",
		Frequency: 30,
		NextTime:  time.Now(),
		Status:    true,
	}
	result := service.TaskServiceImpl{}.Add(task)
	if !result {
		log.Fatal("error")
	}
}

func TestTaskQuery(*testing.T) {
	result := service.TaskServiceImpl{}.Query("task")
	fmt.Println(result)
}

func TestTaskModify(*testing.T) {
	task := task_dao.Task{
		TaskCode:  "task",
		TaskName:  "sum",
		RuleCode:  "rule",
		Frequency: 90,
		NextTime:  time.Now(),
		Status:    false,
	}
	result := service.TaskServiceImpl{}.Modify(task)
	if !result {
		log.Fatal("error")
	}
}

func TestTaskDelete(*testing.T) {
	result := service.TaskServiceImpl{}.Delete("task")
	if !result {
		log.Fatal("error")
	}
}
