package _interface

import (
	"alert/core/dao/rule_dao"
)

type RuleInterface interface {
	Serialization(RuleJson rule_dao.RuleJson) string
	AntiSerialization(Expression string) rule_dao.RuleJson
	Add(rule rule_dao.Rule) bool
	Delete(RuleCode string) bool
	Query(RuleCode string) rule_dao.Rule
	Modify(rule rule_dao.Rule) bool
}
