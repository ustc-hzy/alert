package handlers

import (
	"alert/kitex_gen/api"
	"context"
)

// ScheduleImpl implements the last service interface defined in the IDL.
type ScheduleImpl struct{}

// Schedule implements the ScheduleImpl interface.
func (s *ScheduleImpl) Schedule(ctx context.Context, req *api.ScheduleRequest) (resp *api.ScheduleResponse, err error) {
	// TODO: Your code here...
	return
}
