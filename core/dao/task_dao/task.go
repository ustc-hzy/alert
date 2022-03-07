package task_dao

import (
	"gorm.io/gorm"
	"time"
)

type Task struct {
	TaskCode  string        `gorm:"primaryKey" json:"task_code"`
	TaskName  string        `json:"task_name"`
	RuleCode  string        `json:"rule_code"`
	Frequency time.Duration `json:"frequency"`
	NextTime  time.Time     `json:"next_time"`
	Status    bool          `json:"status"`
	DeletedAt gorm.DeletedAt
}
