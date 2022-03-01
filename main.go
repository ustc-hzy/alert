package main

import (
	"alert/core/dto"
	"alert/handlers"
	api "alert/kitex_gen/api/schedule"
	"log"
)

var Tasklist []dto.TaskVO

func main() {
	svr := api.NewServer(new(handlers.ScheduleImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
