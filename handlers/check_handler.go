package handlers

import (
	"alert/kitex_gen/api"
	"context"
)

type CheckImpl struct{}

// Check implements the CheckImpl interface.
func (s *CheckImpl) Check(ctx context.Context, req *api.CheckRequest) (resp *api.CheckResponse, err error) {
	// TODO: Your code here...
	return
}
