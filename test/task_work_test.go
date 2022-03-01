package test

import (
	"alert/core/dao/task_dao"
	"alert/core/service"
	"fmt"
	"log"
	"testing"
	"time"
)

func Test(t *testing.T) {
	task := task_dao.Task{TaskCode: "task1", TaskName: "task", RuleCode: "rule", Frequency: 30, NextTime: time.Now(), Status: true}
	result := service.TaskServiceImpl{}.Add(task)
	if !result {
		log.Fatal("error")
	}
}

func TestQuery(*testing.T) {
	result := service.TaskServiceImpl{}.Query("task")
	fmt.Println(result)
}
