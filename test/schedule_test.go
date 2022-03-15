package test

import (
	"alert/core/service"
	"testing"
)

func TestSchedule(t *testing.T) {
	service.ScheduleImpl{}.Schedule(50)
}
