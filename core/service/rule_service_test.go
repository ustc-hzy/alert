package service

import (
	"alert/core/dao/rule_dao"
	"testing"
	"time"
)

var rule_service = RuleServiceImpl{}

func TestRuleAdd(t *testing.T) {

	rule := rule_dao.Rule{RuleCode: "test", RuleName: "test", RoomId: 0, Expression: "test", Description: "test",
		StartTime: time.Now(), EndTime: time.Now(), CreateTime: time.Now(), UpdateTime: time.Now()}
	ruleJson := rule_dao.RuleJson{}
	result := rule_service.Add(rule, ruleJson)
	if !result {
		t.Fatal("error")
	}
}
