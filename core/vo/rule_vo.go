package vo

import (
	"alert/core"
	"time"
)

type RuleVo struct {
	RuleCode      string `gorm:"primary_key"`
	RuleName      string
	RoomId        uint
	Rules         []RuleVo
	Logic         core.LogicType
	Op            core.OpType
	Value         float64
	IndicatorCode string
	Description   string
	StartTime     time.Time
	EndTime       time.Time
	CreateTime    time.Time
	UpdateTime    time.Time
}
