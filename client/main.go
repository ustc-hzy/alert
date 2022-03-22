package main

import (
	"alert/kitex_gen/api"
	"alert/kitex_gen/api/crud"
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"log"
	"time"
)

func main() {
	cli, err := crud.NewClient("task", client.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		log.Fatal(err)
	}

	//testAdd
	addReq := &api.AddTaskRequest{
		Task: &api.Task{
			TaskCode:  "taskTest1",
			TaskName:  "sum",
			RuleCode:  "test12",
			Frequency: 30,
			NextTime:  "2006-01-02 15:04:05",
			Status:    true,
		},
	}

	addResp, err := cli.AddTask(context.Background(), addReq, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	if addResp.Success {
		log.Println("Added Successfully")
	} else {
		log.Println("Add Failed")
	}

	//test delete
	deleteReq := &api.DeleteTaskRequest{
		TaskCode: "taskTest1",
	}
	deleteResp, err := cli.DeleteTask(context.Background(), deleteReq, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	if deleteResp.Success {
		log.Println("Deleted Successfully")
	} else {
		log.Println("Delete Failed")
	}

	//test query
	queryReq := &api.QueryTaskRequest{
		TaskCode: "taskTest",
	}
	queryResp, err := cli.QueryTask(context.Background(), queryReq, callopt.WithRPCTimeout(3*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(queryResp)

	//test modify
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
		log.Fatal(err)
	}
	if modifyResp.Success {
		log.Println("Modified Successfully")
	} else {
		log.Println("Modify Failed")
	}

}
