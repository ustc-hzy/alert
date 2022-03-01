package _interface

import (
	indicator_dao "alert/core/dao/indicator_dao"
)

type IndicatorInterface interface {
	Serialization(indicatorJson indicator_dao.IndicatorJson) string
	AntiSerialization(expression string) indicator_dao.IndicatorJson
	IndicatorAdd(indicator indicator_dao.Indicator) bool
	Delete(IndicatorCode string) bool
	Query(IndicatorCode string) indicator_dao.Indicator
	Modify(indicator indicator_dao.Indicator) bool
}
