package main

import (
	"alert/core"
	"alert/core/dao/indicator_dao"
	"alert/core/dao/task_dao"
	"alert/core/service"
	"alert/core/utils"
	"alert/kitex_gen/api"
	"context"
	"github.com/jinzhu/copier"
	"time"
)

// CRUDImpl implements the last service interface defined in the IDL.
type CRUDImpl struct{}

// AddIndicator implements the CRUDImpl interface.
func (s *CRUDImpl) AddIndicator(ctx context.Context, req *api.AddIndicatorRequest) (resp *api.AddIndicatorResponse, err error) {
	indicator, indicatorJson := utils.ConvertIndicator(req)
	res := service.IndicatorServiceImpl{}.Add(indicator, indicatorJson)
	return &api.AddIndicatorResponse{Success: res}, nil
}

// DeleteIndicator implements the CRUDImpl interface.
func (s *CRUDImpl) DeleteIndicator(ctx context.Context, req *api.DeleteIndicatorRequest) (resp *api.DeleteIndicatorResponse, err error) {
	res := service.IndicatorServiceImpl{}.Delete(req.IndicatorCode)
	return &api.DeleteIndicatorResponse{Success: res}, nil
}

// QueryIndicator implements the CRUDImpl interface.
func (s *CRUDImpl) QueryIndicator(ctx context.Context, req *api.QueryIndicatorRequest) (resp *api.QueryIndicatorResponse, err error) {
	res := service.IndicatorServiceImpl{}.Query(req.IndicatorCode)
	indicator := api.Indicator{}
	copier.Copy(&indicator, res)
	indicator.TimeCreate = res.CreateTime.Format(core.TIMETAMPLATE)
	indicator.TimeUpdate = res.UpdateTime.Format(core.TIMETAMPLATE)
	return &api.QueryIndicatorResponse{Indicator: &indicator}, nil
}

// ModifyIndicator implements the CRUDImpl interface.
func (s *CRUDImpl) ModifyIndicator(ctx context.Context, req *api.ModifyIndicatorRequest) (resp *api.ModifyIndicatorResponse, err error) {
	indicator := indicator_dao.Indicator{}
	copier.Copy(&indicator, req.Indicator)
	res := service.IndicatorServiceImpl{}.Modify(indicator)
	return &api.ModifyIndicatorResponse{Success: res}, nil
}

// AddRule implements the CRUDImpl interface.
func (s *CRUDImpl) AddRule(ctx context.Context, req *api.AddRuleRequest) (resp *api.AddRuleResponse, err error) {
	rule := utils.ConvertRule(req)
	ruleJson := utils.ConvertRuleJson(req)
	res := service.RuleServiceImpl{}.Add(rule, ruleJson)
	return &api.AddRuleResponse{Success: res}, nil
}

// DeleteRule implements the CRUDImpl interface.
func (s *CRUDImpl) DeleteRule(ctx context.Context, req *api.DeleteRuleRequest) (resp *api.DeleteRuleResponse, err error) {
	res := service.RuleServiceImpl{}.Delete(req.RuleCode)
	return &api.DeleteRuleResponse{Success: res}, nil
}

// QueryRule implements the CRUDImpl interface.
func (s *CRUDImpl) QueryRule(ctx context.Context, req *api.QueryRuleRequest) (resp *api.QueryRuleResponse, err error) {
	res := service.RuleServiceImpl{}.Query(req.RuleCode)
	rule := api.Rule{}
	copier.Copy(&rule, res)
	rule.TimeCreate = res.CreateTime.Format(core.TIMETAMPLATE)
	rule.TimeUpdate = res.UpdateTime.Format(core.TIMETAMPLATE)
	return &api.QueryRuleResponse{Rule: &rule}, nil
}

// ModifyRule implements the CRUDImpl interface.
func (s *CRUDImpl) ModifyRule(ctx context.Context, req *api.ModifyRuleRequest) (resp *api.ModifyRuleResponse, err error) {
	rule := utils.ConvertRule2(req)
	res := service.RuleServiceImpl{}.Modify(rule)
	return &api.ModifyRuleResponse{Success: res}, nil
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
