package rpc

import (
	"alert/kitex_gen/api"
	"alert/kitex_gen/api/crud"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"log"
	"testing"
	"time"
)

const ADDR = "127.0.0.1:8888"

var indicator = &api.Indicator{
	IndicatorCode: "test11",
	Name:          "test",
	Expression:    "test",
	Description:   "test",
	TimeCreate:    "2006-01-02 15:04:05",
	TimeUpdate:    "2006-01-02 15:04:05",
}
var indicatorJson = &api.IndicatorJson{
	Indicators: nil,
	Calculate:  -1,
	Value:      "select deal_amount from `deal_infos`",
}

var rule = &api.Rule{
	RuleCode:    "test20",
	RuleName:    "test",
	RoomId:      0,
	Expression:  "test",
	Description: "test",
	TimeStart:   "2022-01-02 00:00:00",
	TimeEnd:     "2022-01-02 00:00:00",
	TimeCreate:  "2006-01-02 15:04:05",
	TimeUpdate:  "2006-01-02 15:04:05",
}
var ruleJson = &api.RuleJson{
	Rules:         nil,
	Logic:         0,
	Op:            0,
	Value:         0,
	IndicatorCode: "test10",
}

func Connect() crud.Client {
	c, err := crud.NewClient("crud", client.WithHostPorts(ADDR))
	if err != nil {
		log.Fatalln(err)
	}
	return c
}

func TestRpcAddIndicator(t *testing.T) {
	c := Connect()
	addIndicatorReq := &api.AddIndicatorRequest{
		Indicator:     indicator,
		IndicatorJson: indicatorJson,
	}
	addIndicatorResp, err := c.AddIndicator(context.Background(), addIndicatorReq, callopt.WithConnectTimeout(3*time.Second))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(addIndicatorResp)
}

func TestRpcDeleteIndicator(t *testing.T) {
	c := Connect()
	deleteIndicatorReq := &api.DeleteIndicatorRequest{
		IndicatorCode: "test10",
	}
	deleteIndicatorResp, err := c.DeleteIndicator(context.Background(), deleteIndicatorReq, callopt.WithConnectTimeout(3*time.Second))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(deleteIndicatorResp)
}

func TestRpcModifyIndicator(t *testing.T) {
	c := Connect()
	modifyIndicatorReq := &api.ModifyIndicatorRequest{
		Indicator: indicator,
	}
	modifyIndicatorResp, err := c.ModifyIndicator(context.Background(), modifyIndicatorReq, callopt.WithConnectTimeout(3*time.Second))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(modifyIndicatorResp)
}

func TestRpcQueryIndicator(t *testing.T) {
	c := Connect()
	queryIndicatorReq := &api.QueryIndicatorRequest{
		IndicatorCode: "test11",
	}
	queryIndicatorResp, err := c.QueryIndicator(context.Background(), queryIndicatorReq, callopt.WithConnectTimeout(3*time.Second))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(queryIndicatorResp)
}

func TestRpcAddRule(t *testing.T) {
	c := Connect()
	addRuleReq := &api.AddRuleRequest{
		Rule:     rule,
		RuleJson: ruleJson,
	}
	addRuleResp, err := c.AddRule(context.Background(), addRuleReq, callopt.WithConnectTimeout(3*time.Second))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(addRuleResp)
}

func TestRpcDeleteRule(t *testing.T) {
	c := Connect()
	deleteRuleReq := &api.DeleteRuleRequest{
		RuleCode: "test20",
	}
	deleteRuleResp, err := c.DeleteRule(context.Background(), deleteRuleReq, callopt.WithConnectTimeout(3*time.Second))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(deleteRuleResp)
}

func TestRpcModifyRule(t *testing.T) {
	c := Connect()
	modifyRuleReq := &api.ModifyRuleRequest{
		Rule: rule,
	}
	modifyRuleResp, err := c.ModifyRule(context.Background(), modifyRuleReq, callopt.WithConnectTimeout(3*time.Second))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(modifyRuleResp)
}

func TestRpcQueryRule(t *testing.T) {
	c := Connect()
	queryRuleReq := &api.QueryRuleRequest{
		RuleCode: "test20",
	}
	queryRuleResp, err := c.QueryRule(context.Background(), queryRuleReq, callopt.WithConnectTimeout(3*time.Second))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(queryRuleResp)
}

func TestRpcAddTask(t *testing.T) {
	cli := Connect()
	addReq := &api.AddTaskRequest{
		Task: &api.Task{
			TaskCode:  "taskTest2",
			TaskName:  "test",
			RuleCode:  "testCaseRule",
			Frequency: 30,
			NextTime:  "2006-01-02 15:04:05",
			Status:    true,
		},
	}

	addResp, err := cli.AddTask(context.Background(), addReq, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatalln(err)
	}
	if addResp.Success {
		log.Println("Added Successfully")
	} else {
		log.Println("Add Failed")
	}
}

func TestRpcDeleteTask(t *testing.T) {
	cli := Connect()
	deleteReq := &api.DeleteTaskRequest{
		TaskCode: "taskTest1",
	}
	deleteResp, err := cli.DeleteTask(context.Background(), deleteReq, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatalln(err)
	}
	if deleteResp.Success {
		log.Println("Deleted Successfully")
	} else {
		log.Println("Delete Failed")
	}
}

func TestRpcQueryTask(t *testing.T) {
	cli := Connect()
	queryReq := &api.QueryTaskRequest{
		TaskCode: "taskTest",
	}
	queryResp, err := cli.QueryTask(context.Background(), queryReq, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(queryResp)
}

func TestRpcModifyTask(t *testing.T) {
	cli := Connect()
	modifyReq := &api.ModifyTaskRequest{
		Task: &api.Task{
			TaskCode:  "taskTest",
			TaskName:  "sum",
			RuleCode:  "test12",
			Frequency: 90,
			NextTime:  "2006-01-02 15:04:05",
			Status:    true,
		},
	}

	modifyResp, err := cli.ModifyTask(context.Background(), modifyReq, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatalln(err)
	}
	if modifyResp.Success {
		log.Println("Modified Successfully")
	} else {
		log.Println("Modify Failed")
	}
}
