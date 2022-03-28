package alert_dao

import "time"

type Alert struct {
	RuleName   string    `json:"rule_name"`
	RuleCode   string    `json:"rule_code"`
	TaskName   string    `json:"task_name"`
	Expression string    `json:"expression"`
	Value      string    `json:"value"`
	RoomID     uint      `json:"room_id"`
	Time       time.Time `json:"time"`
}
