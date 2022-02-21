package _interface

import (
	indicator_dao "alert/core/dao/indicator_dao"
	"time"
)

type IndicatorInterface interface {
	Serialization(Indicators []indicator_dao.Indicator, Op indicator_dao.Op_type, Value string) string
	AntiSerialization(Expression string) indicator_dao.IndicatorJson
	Add(IndicatorCode string, Name string, Expression string, Description string, CreateTime time.Time,
		UpdateTime time.Time) bool
	Delete(IndicatorCode string) bool
	Query(IndicatorCode string) indicator_dao.Indicator
	Modify(IndicatorCode string, Name string, Expression string, Description string, CreateTime time.Time,
		UpdateTime time.Time) bool
}
