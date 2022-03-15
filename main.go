package main

import (
	api "alert/kitex_gen/api/crud"
	"log"
)

func main() {
	svr := api.NewServer(new(CRUDImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
