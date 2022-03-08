package service

import (
	"alert/core"
	"alert/core/dao"
	"alert/core/dao/rule_dao"
	"alert/core/vo"
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

type RuleServiceImpl struct{}
type RuleCheckImpl struct{}

const RULETABLENAME = "rules"

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
	if IsRuleExist(rule.RuleCode) {
		log.Fatalln("the indicator code already exist")
		return false
	}

	res1 := dao.DB.Debug().Create(&rule)
	if res1.Error != nil {
		log.Fatalln(res1.Error)
		return false
	}

	return true
}
func (i RuleServiceImpl) Delete(ruleCode string) bool {
	res := dao.DB.Debug().Table(RULETABLENAME).Where("rule_code = ? ", ruleCode).Update("is_delete", 1)
	if res.Error != nil {
		log.Fatalln(res.Error)
		return false
	}

	return true
}
func (i RuleServiceImpl) Query(ruleCode string) vo.RuleVo {

	rule := rule_dao.Rule{}
	res := dao.DB.Debug().Table(RULETABLENAME).Where("rule_code = ?", ruleCode).Where("is_delete = ?", 0).Find(&rule)
	if res.Error != nil {
		log.Fatalln(res.Error)
	}
	if res.RowsAffected > 1 {
		log.Fatalln("code is not only")
	} else if res.RowsAffected == 0 {
		log.Print("not found")
		return vo.RuleVo{}
	}
	fmt.Println(res.RowsAffected)
	//fmt.Println(rule.Expression)
	if rule.Expression == "" {
		log.Fatalln("empty expression")
	}
	//construct VO
	rulejson := i.AntiSerialization(rule.Expression)
	ruleVo := vo.RuleVo{
		RuleCode:      rule.RuleCode,
		RuleName:      rule.RuleName,
		RoomId:        rule.RoomId,
		Rules:         rulejson.Rules,
		Logic:         rulejson.Logic,
		Op:            rulejson.Op,
		Value:         rulejson.Value,
		IndicatorCode: rulejson.IndicatorCode,
		Description:   rule.Description,
		StartTime:     rule.StartTime,
		EndTime:       rule.EndTime,
		CreateTime:    rule.CreateTime,
		UpdateTime:    rule.UpdateTime,
	}

	fmt.Println(ruleVo)
	return ruleVo
}
func (i RuleServiceImpl) Modify(rule rule_dao.Rule) bool {
	if !IsRuleExist(rule.RuleCode) {
		log.Fatalln("the indicator code not exist")
		return false
	}
	res := dao.DB.Debug().Omit("create_time").Where("rule_code", rule.RuleCode).Where("is_delete = ?", 0).Save(&rule)
	if res.Error != nil {
		log.Fatalln(res.Error)
	}
	return true
}

func IsRuleExist(ruleCode string) bool {
	var count int64
	res := dao.DB.Debug().Table(RULETABLENAME).Where("rule_code = ?", ruleCode).Where("is_delete = ?", 0).Count(&count)
	if res.Error != nil {
		log.Fatalln(res.Error)
		return false
	}
	if count != 0 {
		return true
	}
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
