package dto

import "time"

type OpType int32

const (
	LARGER   OpType = 0
	SMALLER  OpType = 1
	EQUAL    OpType = 2
	NOTEQUAL OpType = 3
)

type LogicType int32

const (
	AND LogicType = 0
	OR  LogicType = 1
)

type RuleVo struct {
	RuleCode    string `gorm:"primary_key"`
	RuleName    string
	RoomId      uint
	Rules       []RuleVo
	Logic       LogicType
	Op          OpType
	Value       uint
	Description string
	StartTime   time.Time
	EndTime     time.Time
	CreateTime  time.Time
	UpdateTime  time.Time
}
