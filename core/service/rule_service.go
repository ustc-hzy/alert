package service

import (
	"alert/core/dao/rule_dao"
	"alert/core/dto"
	"errors"
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
func (i RuleServiceImpl) Query(RuleCode string) dto.RuleVo {
	return dto.RuleVo{}
}
func (i RuleServiceImpl) Modify(rule rule_dao.Rule) bool {
	return false
}

func (i RuleCheckImpl) Check(RuleCode string) (bool, error) {
	rule := RuleServiceImpl{}.Query(RuleCode)
	return i.CheckLeaf(rule)
}

func (i RuleCheckImpl) CheckLeaf(rule dto.RuleVo) (bool, error) {
	if rule.Rules == nil && rule.Logic == -1 && rule.Op != -1 && rule.Value != 0 && len(rule.IndicatorCode) != 0 {
		//if leaf
		ind := IndComputeImpl{}.Compute(rule.IndicatorCode, rule.RoomId, rule.StartTime, rule.EndTime)
		switch rule.Op {
		case dto.OpType(rule_dao.LARGER):
			if ind > rule.Value {
				return true, nil
			} else {
				return false, nil
			}
			break
		case dto.OpType(rule_dao.SMALLER):
			if ind < rule.Value {
				return true, nil
			} else {
				return false, nil
			}
			break
		case dto.OpType(rule_dao.EQUAL):
			if ind == rule.Value {
				return true, nil
			} else {
				return false, nil
			}
			break
		case dto.OpType(rule_dao.NOTEQUAL):
			if ind != rule.Value {
				return true, nil
			} else {
				return false, nil
			}
			break
		}
	} else if rule.Rules != nil && rule.Logic != -1 && rule.Op == -1 && rule.Value == 0 && len(rule.IndicatorCode) == 0 {
		//if not leaf
		r1, e1 := i.CheckLeaf(rule.Rules[0])
		r2, e2 := i.CheckLeaf(rule.Rules[1])
		switch rule.Logic {
		case dto.LogicType(rule_dao.AND):
			if e1 == nil {
				return r1 && r2, e2
			} else {
				return r1 && r2, e1
			}
			break
		case dto.LogicType(rule_dao.OR):
			if e1 == nil {
				return r1 || r2, e2
			} else {
				return r1 || r2, e1
			}
			break
		}
	}
	return false, errors.New("data error")

}
