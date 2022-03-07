package main

import (
	"alert/core/vo"
	"alert/handlers"
	api "alert/kitex_gen/api/schedule"
	"log"
)

var Tasklist []vo.TaskVO

func main() {
	svr := api.NewServer(new(handlers.ScheduleImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
