package alert_dao

import "time"

type Alert struct {
	RuleName string    `json:"rule_name"`
	RoomID   uint      `json:"room_id"`
	Time     time.Time `json:"time"`
}
