package rule_dao

import (
	"time"
)

type OpType int32

const (
	NULL     OpType = -1
	LARGER   OpType = 0
	SMALLER  OpType = 1
	EQUAL    OpType = 2
	NOTEQUAL OpType = 3
)

type LogicType int32

const (
	NIL LogicType = -1
	AND LogicType = 0
	OR  LogicType = 1
)

type Rule struct {
	RuleCode    string `gorm:"primary_key"`
	RuleName    string
	RoomId      uint
	Expression  string
	Description string
	StartTime   time.Time
	EndTime     time.Time
	CreateTime  time.Time
	UpdateTime  time.Time
}

type RuleJson struct {
	Rules         []Rule
	Logic         LogicType
	Op            OpType
	Value         uint
	IndicatorCode string
}
