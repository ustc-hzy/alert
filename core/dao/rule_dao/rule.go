package rule_dao

import (
	"alert/core"
	"alert/core/vo"
	"time"
)

type Rule struct {
	RuleCode    string    `gorm:"rule_code"`
	RuleName    string    `gorm:"rule_name"`
	RoomId      uint      `gorm:"room_id"`
	Expression  string    `gorm:"expression"`
	Description string    `gorm:"description"`
	StartTime   time.Time `gorm:"start_time"`
	EndTime     time.Time `gorm:"end_time"`
	CreateTime  time.Time `gorm:"create_time"`
	UpdateTime  time.Time `gorm:"update_time"`
}

type RuleJson struct {
	Rules         []vo.RuleVo
	Logic         core.LogicType
	Op            core.OpType
	Value         float64
	IndicatorCode string
}
