package service

import (
	"alert/core"
	"alert/core/dao/rule_dao"
	"fmt"
	"testing"
	"time"
)

var rule_service = RuleServiceImpl{}
var rule_check_service = RuleCheckImpl{}

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
