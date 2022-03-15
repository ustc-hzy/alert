package main

import (
	"alert/core/dao/indicator_dao"
	"alert/core/dao/rule_dao"
	"alert/core/service"
	"alert/core/dao/task_dao"
	"alert/kitex_gen/api"
	"context"
	"github.com/jinzhu/copier"
	"time"
)

// CRUDImpl implements the last service interface defined in the IDL.
type CRUDImpl struct{}

// AddIndicator implements the CRUDImpl interface.
func (s *CRUDImpl) AddIndicator(ctx context.Context, req *api.AddIndicatorRequest) (resp *api.AddIndicatorResponse, err error) {
	// TODO: Your code here...
	indicator, indicatorJson := ConvertIndicator(req)
	service.IndicatorServiceImpl{}.Add(indicator, indicatorJson)
	return
}

func ConvertIndicator(req *api.AddIndicatorRequest) (indicator_dao.Indicator, indicator_dao.IndicatorJson) {
	indicator := indicator_dao.Indicator{}
	indicatorJson := indicator_dao.IndicatorJson{}
	copier.Copy(&indicator, req.Indicator)
	timeTemplate := "2006-01-02 15:04:05"
	createTimeStr := req.Indicator.GetTimeCreate()
	updateTimeStr := req.Indicator.GetTimeUpdate()
	createTime, _ := time.ParseInLocation(timeTemplate, createTimeStr, time.Local)
	updateTime, _ := time.ParseInLocation(timeTemplate, updateTimeStr, time.Local)
	indicator.CreateTime = createTime
	indicator.UpdateTime = updateTime
	copier.Copy(&indicatorJson, req.IndicatorJson)
	return indicator, indicatorJson
}

// DeleteIndicator implements the CRUDImpl interface.
func (s *CRUDImpl) DeleteIndicator(ctx context.Context, req *api.DeleteIndicatorRequest) (resp *api.DeleteIndicatorResponse, err error) {
	// TODO: Your code here...
	service.IndicatorServiceImpl{}.Delete(req.IndicatorCode)
	return
}

// QueryIndicator implements the CRUDImpl interface.
func (s *CRUDImpl) QueryIndicator(ctx context.Context, req *api.QueryIndicatorRequest) (resp *api.QueryIndicatorResponse, err error) {
	// TODO: Your code here...
	service.IndicatorServiceImpl{}.Query(req.IndicatorCode)
	return
}

// ModifyIndicator implements the CRUDImpl interface.
func (s *CRUDImpl) ModifyIndicator(ctx context.Context, req *api.ModifyIndicatorRequest) (resp *api.ModifyIndicatorResponse, err error) {
	// TODO: Your code here...
	indicator := indicator_dao.Indicator{}
	copier.Copy(&indicator, req.Indicator)
	service.IndicatorServiceImpl{}.Modify(indicator)
	return
}

// AddRule implements the CRUDImpl interface.
func (s *CRUDImpl) AddRule(ctx context.Context, req *api.AddRuleRequest) (resp *api.AddRuleResponse, err error) {
	// TODO: Your code here...
	rule, ruleJson := ConvertRule(req)
	service.RuleServiceImpl{}.Add(rule, ruleJson)
	return
}

func ConvertRule(req *api.AddRuleRequest) (rule_dao.Rule, rule_dao.RuleJson) {
	rule := rule_dao.Rule{}
	ruleJson := rule_dao.RuleJson{}
	copier.Copy(&rule, req.Rule)
	timeTemplate := "2006-01-02 15:04:05"
	createTimeStr := req.Rule.GetTimeCreate()
	updateTimeStr := req.Rule.GetTimeUpdate()
	createTime, _ := time.ParseInLocation(timeTemplate, createTimeStr, time.Local)
	updateTime, _ := time.ParseInLocation(timeTemplate, updateTimeStr, time.Local)
	rule.CreateTime = createTime
	rule.UpdateTime = updateTime
	copier.Copy(&ruleJson, req.RuleJson)
	return rule, ruleJson
}

// DeleteRule implements the CRUDImpl interface.
func (s *CRUDImpl) DeleteRule(ctx context.Context, req *api.DeleteRuleRequest) (resp *api.DeleteRuleResponse, err error) {
	// TODO: Your code here...
	service.RuleServiceImpl{}.Delete(req.RuleCode)
	return
}

// QueryRule implements the CRUDImpl interface.
func (s *CRUDImpl) QueryRule(ctx context.Context, req *api.QueryRuleRequest) (resp *api.QueryRuleResponse, err error) {
	// TODO: Your code here...
	service.RuleServiceImpl{}.Query(req.RuleCode)
	return
}

// ModifyRule implements the CRUDImpl interface.
func (s *CRUDImpl) ModifyRule(ctx context.Context, req *api.ModifyRuleRequest) (resp *api.ModifyRuleResponse, err error) {
	// TODO: Your code here...
	rule := rule_dao.Rule{}
	copier.Copy(&rule, req.Rule)
	service.RuleServiceImpl{}.Modify(rule)
	return
}

// AddTask implements the CRUDImpl interface.
func (s *CRUDImpl) AddTask(ctx context.Context, req *api.AddTaskRequest) (resp *api.AddTaskResponse, err error) {
	task := task_dao.Task{}
	copier.Copy(&task, &req.Task)
	task.NextTime, _ = time.ParseInLocation("2006-01-02 15:04:05", req.Task.NextTime, time.Local)
	task.Frequency = time.Duration(req.Task.Frequency)
	res := service.TaskServiceImpl{}.Add(task)
	return &api.AddTaskResponse{Success: res}, nil
}

// DeleteTask implements the CRUDImpl interface.
func (s *CRUDImpl) DeleteTask(ctx context.Context, req *api.DeleteTaskRequest) (resp *api.DeleteTaskResponse, err error) {
	res := service.TaskServiceImpl{}.Delete(req.TaskCode)
	return &api.DeleteTaskResponse{Success: res}, nil
}

// QueryTask implements the CRUDImpl interface.
func (s *CRUDImpl) QueryTask(ctx context.Context, req *api.QueryTaskRequest) (resp *api.QueryTaskResponse, err error) {
	res := service.TaskServiceImpl{}.Query(req.TaskCode)
	task := api.Task{}
	copier.Copy(&task, &res)
	task.NextTime = res.NextTime.Format("2006-01-02 15:04:05")
	task.Frequency = int64(int(res.Frequency))
	return &api.QueryTaskResponse{Task: &task}, nil
}

// ModifyTask implements the CRUDImpl interface.
func (s *CRUDImpl) ModifyTask(ctx context.Context, req *api.ModifyTaskRequest) (resp *api.ModifyTaskResponse, err error) {
	task := task_dao.Task{}
	copier.Copy(&task, &req.Task)
	task.NextTime, _ = time.ParseInLocation("2006-01-02 15:04:05", req.Task.NextTime, time.Local)
	task.Frequency = time.Duration(req.Task.Frequency)
	res := service.TaskServiceImpl{}.Modify(task)
	return &api.ModifyTaskResponse{Success: res}, nil
}
