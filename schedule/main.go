package main

import (
	"alert/core/service"
	"time"
)

func main() {
	service.ScheduleImpl{}.Schedule(1 * time.Second)
}
