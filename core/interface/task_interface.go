package _interface

import (
	"alert/core/dao/task_dao"
	"alert/core/vo"
)

type TaskInterface interface {
	Add(task task_dao.Task) bool
	Delete(taskCode string) bool
	Query(taskCode string) vo.TaskVO
	Modify(task task_dao.Task) bool
	UpdateTime(task task_dao.Task) bool
	UpdateStatus(task task_dao.Task, status bool) bool
	TransferTaskVo(task task_dao.Task) vo.TaskVO
}
