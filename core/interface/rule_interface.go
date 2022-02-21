package _interface

import (
	rule_dao "alert/core/dao/rule_dao"
	"time"
)

type RuleInterface interface {
	Serialization(Rules []rule_dao.Rule, Logic rule_dao.LogicType, Op rule_dao.OpType, Value uint) string
	AntiSerialization(Expression string) rule_dao.RuleJson
	Add(RuleCode string, RuleName string, RoomId uint, Expression string, Description string, StartTime time.Time,
		EndTime time.Time, CreateTime time.Time, UpdateTime time.Time) bool
	Delete(RuleCode string) bool
	Query(RuleCode string) rule_dao.Rule
	Modify(RuleCode string, RuleName string, RoomId uint, Expression string, Description string, StartTime time.Time,
		EndTime time.Time, CreateTime time.Time, UpdateTime time.Time) bool
}
