package test

import (
	"alert/kitex_gen/api"
	"alert/kitex_gen/api/schedule"
	"context"
	"github.com/cloudwego/kitex/client"
	"log"
	"testing"
)

func TestSchedule(t *testing.T) {
	cli, err := schedule.NewClient("add", client.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		log.Fatal(err)
	}
	req := &api.ScheduleRequest{10000}
	resp, err := cli.Schedule(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
