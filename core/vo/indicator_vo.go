package vo

import (
	"alert/core"
	"time"
)

type IndicatorVO struct {
	IndicatorCode string
	Name          string
	Indicators    []IndicatorVO
	Caculate      core.CaculateType
	Value         string
	Description   string
	CreateTime    time.Time
	UpdateTime    time.Time
}

type Condition struct {
	RoomID    uint
	StartTime string
	EndTime   string
}
