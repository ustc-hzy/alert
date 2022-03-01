package service

import (
	"alert/core/dao/rule_dao"
)

type RuleServiceImpl struct{}
type RuleCheckImpl struct{}

func (i RuleServiceImpl) Serialization(json rule_dao.RuleJson) string {

	return ""
}
func (i RuleServiceImpl) AntiSerialization(Expression string) rule_dao.RuleJson {
	return rule_dao.RuleJson{}
}
func (i RuleServiceImpl) Add(rule rule_dao.Rule) bool {

	return false
}
func (i RuleServiceImpl) Delete(RuleCode string) bool {
	return false
}
func (i RuleServiceImpl) Query(RuleCode string) rule_dao.Rule {
	return rule_dao.Rule{}
}
func (i RuleServiceImpl) Modify(rule rule_dao.Rule) bool {
	return false
}

func (i RuleCheckImpl) Check(RuleCode string) bool {
	RuleServiceImpl{}.Query(RuleCode)

	return false
}
