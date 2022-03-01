package indicator_dao

import "time"

type Op_type int32

const (
	ADD      Op_type = 0
	SUBTRACT Op_type = 1
	MULTIPLY Op_type = 2
	DIVIDE   Op_type = 3
)

type Indicator struct {
	IndicatorCode string    `gorm:"column:code"`
	Name          string    `gorm:"column:indicatorName"`
	Expression    string    `gorm:"column:expression"`
	Description   string    `gorm:"column:description"`
	CreateTime    time.Time `gorm:"column:create_time"`
	UpdateTime    time.Time `gorm:"column:update_time"`
}

type IndicatorJson struct {
	Indicators []Indicator
	Op         Op_type
	Value      string
}
