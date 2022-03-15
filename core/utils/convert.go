package utils

import (
	"alert/core"
	"alert/core/dao/indicator_dao"
	"alert/core/dao/rule_dao"
	"alert/kitex_gen/api"
	"github.com/jinzhu/copier"
	"time"
)

func ConvertIndicator(req *api.AddIndicatorRequest) (indicator_dao.Indicator, indicator_dao.IndicatorJson) {
	indicator := indicator_dao.Indicator{}
	indicatorJson := indicator_dao.IndicatorJson{}
	copier.Copy(&indicator, req.Indicator)
	createTimeStr := req.Indicator.GetTimeCreate()
	updateTimeStr := req.Indicator.GetTimeUpdate()
	createTime, _ := time.ParseInLocation(core.TIMETAMPLATE, createTimeStr, time.Local)
	updateTime, _ := time.ParseInLocation(core.TIMETAMPLATE, updateTimeStr, time.Local)
	indicator.CreateTime = createTime
	indicator.UpdateTime = updateTime
	copier.Copy(&indicatorJson, req.IndicatorJson)
	return indicator, indicatorJson
}

func ConvertRule(req *api.AddRuleRequest) rule_dao.Rule {
	rule := rule_dao.Rule{}

	copier.Copy(&rule, req.Rule)
	createTimeStr := req.Rule.GetTimeCreate()
	updateTimeStr := req.Rule.GetTimeUpdate()
	startTimeStr := req.Rule.GetTimeStart()
	endTimeStr := req.Rule.GetTimeEnd()
	createTime, _ := time.ParseInLocation(core.TIMETAMPLATE, createTimeStr, time.Local)
	updateTime, _ := time.ParseInLocation(core.TIMETAMPLATE, updateTimeStr, time.Local)
	startTime, _ := time.ParseInLocation(core.TIMETAMPLATE, startTimeStr, time.Local)
	endTime, _ := time.ParseInLocation(core.TIMETAMPLATE, endTimeStr, time.Local)
	rule.CreateTime = createTime
	rule.UpdateTime = updateTime
	rule.StartTime = startTime
	rule.EndTime = endTime

	return rule
}

func ConvertRuleJson(req *api.AddRuleRequest) rule_dao.RuleJson {
	ruleJson := rule_dao.RuleJson{}
	copier.Copy(&ruleJson, req.RuleJson)
	return ruleJson
}

func ConvertRule2(req *api.ModifyRuleRequest) rule_dao.Rule {
	rule := rule_dao.Rule{}

	copier.Copy(&rule, req.Rule)
	createTimeStr := req.Rule.GetTimeCreate()
	updateTimeStr := req.Rule.GetTimeUpdate()
	startTimeStr := req.Rule.GetTimeStart()
	endTimeStr := req.Rule.GetTimeEnd()
	createTime, _ := time.ParseInLocation(core.TIMETAMPLATE, createTimeStr, time.Local)
	updateTime, _ := time.ParseInLocation(core.TIMETAMPLATE, updateTimeStr, time.Local)
	startTime, _ := time.ParseInLocation(core.TIMETAMPLATE, startTimeStr, time.Local)
	endTime, _ := time.ParseInLocation(core.TIMETAMPLATE, endTimeStr, time.Local)
	rule.CreateTime = createTime
	rule.UpdateTime = updateTime
	rule.StartTime = startTime
	rule.EndTime = endTime

	return rule
}
