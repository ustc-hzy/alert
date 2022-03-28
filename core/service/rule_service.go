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
	"strconv"
	"time"
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
		log.Println(err)
	}
	return result
}
func (i RuleServiceImpl) Add(rule rule_dao.Rule, ruleJson rule_dao.RuleJson) bool {

	rule.Expression = i.Serialization(ruleJson)
	if i.IsRuleExist(rule.RuleCode) {
		log.Println("the rule code already exist")
		return false
	}

	res1 := dao.DB.Debug().Create(&rule)
	if res1.Error != nil {
		log.Println(res1.Error)
		return false
	}

	return true
}
func (i RuleServiceImpl) Delete(ruleCode string) bool {
	res := dao.DB.Debug().Table(RULETABLENAME).Where("rule_code = ? ", ruleCode).Update("is_delete", 1)
	if res.Error != nil {
		log.Println(res.Error)
		return false
	}

	return true
}
func (i RuleServiceImpl) Query(ruleCode string) vo.RuleVo {

	rule := rule_dao.Rule{}
	res := dao.DB.Debug().Table(RULETABLENAME).Where("rule_code = ?", ruleCode).Where("is_delete = ?", 0).Find(&rule)
	if res.Error != nil {
		log.Println(res.Error)
	}
	if res.RowsAffected > 1 {
		log.Println("code is not only")
	} else if res.RowsAffected == 0 {
		log.Print("not found")
		return vo.RuleVo{}
	}
	fmt.Println(res.RowsAffected)
	//fmt.Println(rule.Expression)
	if rule.Expression == "" {
		log.Println("empty expression")
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
	if !i.IsRuleExist(rule.RuleCode) {
		log.Println("the indicator code not exist")
		return false
	}
	rule.UpdateTime = time.Now()
	res := dao.DB.Debug().Omit("create_time").Where("rule_code", rule.RuleCode).Where("is_delete = ?", 0).Save(&rule)
	if res.Error != nil {
		log.Println(res.Error)
	}
	return true
}

func (i RuleServiceImpl) IsRuleExist(ruleCode string) bool {
	var count int64
	res := dao.DB.Debug().Table(RULETABLENAME).Where("rule_code = ?", ruleCode).Where("is_delete = ?", 0).Count(&count)
	if res.Error != nil {
		log.Println(res.Error)
		return false
	}
	if count != 0 {
		return true
	}
	return false
}

func (i RuleServiceImpl) RuleExpression(rule vo.RuleVo) (string, error) {
	if rule.Rules == nil && rule.Logic == -1 && rule.Op != -1 && rule.Value != 0 && len(rule.IndicatorCode) != 0 {
		//if leaf
		var op string
		switch rule.Op {
		case core.LARGER:
			op = ">"
			break
		case core.SMALLER:
			op = "<"
			break
		case core.EQUAL:
			op = "="
			break
		case core.NOTEQUAL:
			op = "!="
			break
		}
		ret := rule.IndicatorCode + op + strconv.FormatFloat(rule.Value, 'f', 0, 64)
		return ret, nil
	} else if rule.Rules != nil && rule.Logic != -1 && rule.Op == -1 && rule.Value == 0 && len(rule.IndicatorCode) == 0 {
		//if not leaf
		r1, e1 := i.RuleExpression(rule.Rules[0])
		r2, e2 := i.RuleExpression(rule.Rules[1])
		switch rule.Logic {
		case core.AND:
			ret := "(" + r1 + "&&" + r2 + ")"
			return ret, e1
		case core.OR:
			ret := "(" + r1 + "||" + r2 + ")"
			return ret, e2
		}

	}
	return "", errors.New("data error")
}

func (i RuleCheckImpl) Check(ruleCode string) (vo.CheckResult, error) {
	rule := RuleServiceImpl{}.Query(ruleCode)
	value := make(map[string]int)
	res, err := i.CheckLeaf(rule, value)

	result := vo.CheckResult{
		Res:   res,
		Value: value,
	}
	return result, err
}

func (i RuleCheckImpl) CheckLeaf(rule vo.RuleVo, value map[string]int) (bool, error) {
	if rule.Rules == nil && rule.Logic == -1 && rule.Op != -1 && rule.Value != 0 && len(rule.IndicatorCode) != 0 {
		//if leaf
		ind := IndComputeImpl{}.Compute(rule.IndicatorCode, vo.Condition{
			RoomID:    rule.RoomId,
			StartTime: "",
			EndTime:   "",
		})
		value[rule.IndicatorCode] = int(ind)
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
		r1, e1 := i.CheckLeaf(rule.Rules[0], value)
		r2, e2 := i.CheckLeaf(rule.Rules[1], value)
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
