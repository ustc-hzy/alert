package main

import (
	"alert/core/service"
	api "alert/kitex_gen/api/schedule"
	"log"
)

func main() {
	svr := api.NewServer(new(service.ScheduleImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
