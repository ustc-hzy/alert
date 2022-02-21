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
	IndicatorCode string
	Name          string
	Expression    string
	Description   string
	CreateTime    time.Time
	UpdateTime    time.Time
}

type IndicatorJson struct {
	Indicators []Indicator
	Op         Op_type
	Value      string
}
