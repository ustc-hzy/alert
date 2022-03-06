package service

import (
	"alert/core"
	"alert/core/dao/rule_dao"
	"alert/core/vo"
	"encoding/json"
	"errors"
	"log"
)

type RuleServiceImpl struct{}
type RuleCheckImpl struct{}

func (i RuleServiceImpl) Serialization(ruleJson rule_dao.RuleJson) string {
	result, _ := json.Marshal(ruleJson)
	return string(result)
}
func (i RuleServiceImpl) AntiSerialization(expression string) rule_dao.RuleJson {
	byteStream := []byte(expression)
	var result rule_dao.RuleJson
	err := json.Unmarshal(byteStream, &result)
	if err != nil {
		log.Fatalln(err)
	}
	return result
}
func (i RuleServiceImpl) Add(rule rule_dao.Rule, ruleJson rule_dao.RuleJson) bool {
	rule.Expression = i.Serialization(ruleJson)
	if DB.Debug().Table("rules").Where("code = ?", rule.RuleCode) != nil {
		log.Fatalln("the rule code already exist")
		return false
	}
	res := DB.Debug().Create(&rule)
	if res.Error != nil {
		log.Fatalln(res.Error)
		return false
	}

	return true
}
func (i RuleServiceImpl) Delete(ruleCode string) bool {
	return false
}
func (i RuleServiceImpl) Query(ruleCode string) vo.RuleVo {
	return vo.RuleVo{}
}
func (i RuleServiceImpl) Modify(rule rule_dao.Rule) bool {
	return false
}

func (i RuleCheckImpl) Check(ruleCode string) (bool, error) {
	rule := RuleServiceImpl{}.Query(ruleCode)
	return i.CheckLeaf(rule)
}

func (i RuleCheckImpl) CheckLeaf(rule vo.RuleVo) (bool, error) {
	if rule.Rules == nil && rule.Logic == -1 && rule.Op != -1 && rule.Value != 0 && len(rule.IndicatorCode) != 0 {
		//if leaf
		ind := IndComputeImpl{}.Compute(rule.IndicatorCode, rule.RoomId, rule.StartTime, rule.EndTime)
		switch rule.Op {
		case core.LARGER:
			if ind > rule.Value {
				return true, nil
			} else {
				return false, nil
			}
			break
		case core.SMALLER:
			if ind < rule.Value {
				return true, nil
			} else {
				return false, nil
			}
			break
		case core.EQUAL:
			if ind == rule.Value {
				return true, nil
			} else {
				return false, nil
			}
			break
		case core.NOTEQUAL:
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
		case core.AND:
			if e1 == nil {
				return r1 && r2, e2
			} else {
				return r1 && r2, e1
			}
			break
		case core.OR:
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
