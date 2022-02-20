package handlers

import (
	"alert/kitex_gen/api"
	"context"
)

type ComputeImpl struct{}

// Compute implements the ComputeImpl interface.
func (s *ComputeImpl) Compute(ctx context.Context, req *api.ComputeRequest) (resp *api.ComputeResponse, err error) {
	// TODO: Your code here...
	return
}
