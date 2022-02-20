package main

import (
	"alert/core/handlers"
	api "alert/kitex_gen/api/schedule"
	"log"
)

func main() {
	svr := api.NewServer(new(handlers.ScheduleImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
