package test

import (
	"alert/core"
	"alert/core/dao/rule_dao"
	"alert/core/service"
	"alert/core/vo"
	"fmt"
	"testing"
	"time"
)

var rule_service = service.RuleServiceImpl{}
var rule_check_service = service.RuleCheckImpl{}

func TestRuleAdd(t *testing.T) {

	rule := rule_dao.Rule{RuleCode: "test12", RuleName: "test", RoomId: 0, Expression: "test", Description: "test",
		StartTime: time.Now(), EndTime: time.Now(), CreateTime: time.Now(), UpdateTime: time.Now()}
	ruleJson := rule_dao.RuleJson{
		Rules:         nil,
		Logic:         core.NIL,
		Op:            core.EQUAL,
		Value:         2,
		IndicatorCode: "test1",
	}
	result := rule_service.Add(rule, ruleJson)
	if !result {
		t.Fatal("error")
	}
}

func TestRuleDelete(t *testing.T) {

	result := rule_service.Delete("test")
	if !result {
		t.Fatal("error")
	}
}

func TestRuleModify(t *testing.T) {
	rule := rule_dao.Rule{RuleCode: "test", RuleName: "test", Description: "modify", UpdateTime: time.Now(), StartTime: time.Now(), EndTime: time.Now()}
	rule_service.Modify(rule)
}

func TestRuleQuery(t *testing.T) {

	query := rule_service.Query("test1")
	fmt.Println(query)
}

func TestCheck(t *testing.T) {

	val, _ := rule_check_service.Check("test12")
	fmt.Println(val)
}

func TestRuleAddCompound(t *testing.T) {

	ruleA := vo.RuleVo{
		RuleCode:      "ruleA",
		RuleName:      "ruleA",
		RoomId:        0,
		Rules:         nil,
		Logic:         core.NIL,
		Op:            core.EQUAL,
		Value:         0,
		IndicatorCode: "test6",
		Description:   "test",
		StartTime:     time.Now(),
		EndTime:       time.Now(),
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	}
	ruleB := vo.RuleVo{
		RuleCode:      "ruleB",
		RuleName:      "ruleB",
		RoomId:        0,
		Rules:         nil,
		Logic:         core.NIL,
		Op:            core.EQUAL,
		Value:         4,
		IndicatorCode: "nodeALL",
		Description:   "test",
		StartTime:     time.Now(),
		EndTime:       time.Now(),
		CreateTime:    time.Now(),
		UpdateTime:    time.Now(),
	}

	rule := rule_dao.Rule{
		RuleCode:    "ruleAll2",
		RuleName:    "ruleAll2",
		RoomId:      0,
		Expression:  "test",
		Description: "test",
		StartTime:   time.Now(),
		EndTime:     time.Now(),
		CreateTime:  time.Now(),
		UpdateTime:  time.Now(),
	}
	rules := []vo.RuleVo{ruleA, ruleB}
	ruleJson := rule_dao.RuleJson{
		Rules:         rules,
		Logic:         core.OR,
		Op:            core.NIL,
		Value:         0,
		IndicatorCode: "",
	}

	result := rule_service.Add(rule, ruleJson)
	if !result {
		t.Fatal("error")
	}
}

func TestRuleComputeCompound(t *testing.T) {
	val, _ := rule_check_service.Check("ruleAll2")
	fmt.Println(val)
}
