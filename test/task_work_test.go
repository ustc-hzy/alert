package test

import (
	"alert/core/dao/task_dao"
	"alert/core/service"
	"log"
	"testing"
	"time"
)

func Test(t *testing.T) {
	task := task_dao.Task{TaskCode: "task", TaskName: "task", RuleCode: "rule", Frequency: 30, NextTime: time.Now(), Status: true}
	result := service.TaskServiceImpl{}.Add(task)
	if !result {
		log.Fatal("error")
	}
}
