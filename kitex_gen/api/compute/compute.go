// Code generated by Kitex v0.0.8. DO NOT EDIT.

package compute

import (
	"alert/kitex_gen/api"
	"context"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return computeServiceInfo
}

var computeServiceInfo = newServiceInfo()

func newServiceInfo() *kitex.ServiceInfo {
	serviceName := "Compute"
	handlerType := (*api.Compute)(nil)
	methods := map[string]kitex.MethodInfo{
		"compute": kitex.NewMethodInfo(computeHandler, newComputeComputeArgs, newComputeComputeResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "api",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.0.8",
		Extra:           extra,
	}
	return svcInfo
}

func computeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.ComputeComputeArgs)
	realResult := result.(*api.ComputeComputeResult)
	success, err := handler.(api.Compute).Compute(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newComputeComputeArgs() interface{} {
	return api.NewComputeComputeArgs()
}

func newComputeComputeResult() interface{} {
	return api.NewComputeComputeResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Compute(ctx context.Context, req *api.ComputeRequest) (r *api.ComputeResponse, err error) {
	var _args api.ComputeComputeArgs
	_args.Req = req
	var _result api.ComputeComputeResult
	if err = p.c.Call(ctx, "compute", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}