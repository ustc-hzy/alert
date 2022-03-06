package indicator_dao

import (
	"alert/core"
	"alert/core/vo"
	"time"
)

type Indicator struct {
	IndicatorCode string    `gorm:"column:code"`
	Name          string    `gorm:"column:indicatorName"`
	Expression    string    `gorm:"column:Expression"`
	Description   string    `gorm:"column:Description"`
	CreateTime    time.Time `gorm:"column:create_time"`
	UpdateTime    time.Time `gorm:"column:update_time"`
}

type IndicatorJson struct {
	Indicators []vo.IndicatorVO
	Caculate   core.CaculateType
	Value      string
}
