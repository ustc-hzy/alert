package service

import (
	"alert/core/dao/rule_dao"
	"time"
)

type RuleServiceImpl struct{}

func (i RuleServiceImpl) Serialization(Rules []rule_dao.Rule, Logic rule_dao.LogicType, Op rule_dao.OpType, Value uint) string {

}
func (i RuleServiceImpl) AntiSerialization(Expression string) rule_dao.RuleJson {

}
func (i RuleServiceImpl) Add(RuleCode string, RuleName string, RoomId uint, Expression string, Description string,
	StartTime time.Time, EndTime time.Time, CreateTime time.Time, UpdateTime time.Time) bool {

}
func (i RuleServiceImpl) Delete(RuleCode string) bool {

}
func (i RuleServiceImpl) Query(RuleCode string) rule_dao.Rule {

}
func (i RuleServiceImpl) Modify(RuleCode string, RuleName string, RoomId uint, Expression string, Description string,
	StartTime time.Time, EndTime time.Time, CreateTime time.Time, UpdateTime time.Time) bool {

}
