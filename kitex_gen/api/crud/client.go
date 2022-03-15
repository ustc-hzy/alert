// Code generated by Kitex v0.2.0. DO NOT EDIT.

package crud

import (
	"alert/kitex_gen/api"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	AddIndicator(ctx context.Context, req *api.AddIndicatorRequest, callOptions ...callopt.Option) (r *api.AddIndicatorResponse, err error)
	DeleteIndicator(ctx context.Context, req *api.DeleteIndicatorRequest, callOptions ...callopt.Option) (r *api.DeleteIndicatorResponse, err error)
	QueryIndicator(ctx context.Context, req *api.QueryIndicatorRequest, callOptions ...callopt.Option) (r *api.QueryIndicatorResponse, err error)
	ModifyIndicator(ctx context.Context, req *api.ModifyIndicatorRequest, callOptions ...callopt.Option) (r *api.ModifyIndicatorResponse, err error)
	AddRule(ctx context.Context, req *api.AddRuleRequest, callOptions ...callopt.Option) (r *api.AddRuleResponse, err error)
	DeleteRule(ctx context.Context, req *api.DeleteRuleRequest, callOptions ...callopt.Option) (r *api.DeleteRuleResponse, err error)
	QueryRule(ctx context.Context, req *api.QueryRuleRequest, callOptions ...callopt.Option) (r *api.QueryRuleResponse, err error)
	ModifyRule(ctx context.Context, req *api.ModifyRuleRequest, callOptions ...callopt.Option) (r *api.ModifyRuleResponse, err error)
	AddTask(ctx context.Context, req *api.AddTaskRequest, callOptions ...callopt.Option) (r *api.AddTaskResponse, err error)
	DeleteTask(ctx context.Context, req *api.DeleteTaskRequest, callOptions ...callopt.Option) (r *api.DeleteTaskResponse, err error)
	QueryTask(ctx context.Context, req *api.QueryTaskRequest, callOptions ...callopt.Option) (r *api.QueryTaskResponse, err error)
	ModifyTask(ctx context.Context, req *api.ModifyTaskRequest, callOptions ...callopt.Option) (r *api.ModifyTaskResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kCRUDClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kCRUDClient struct {
	*kClient
}

func (p *kCRUDClient) AddIndicator(ctx context.Context, req *api.AddIndicatorRequest, callOptions ...callopt.Option) (r *api.AddIndicatorResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddIndicator(ctx, req)
}

func (p *kCRUDClient) DeleteIndicator(ctx context.Context, req *api.DeleteIndicatorRequest, callOptions ...callopt.Option) (r *api.DeleteIndicatorResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteIndicator(ctx, req)
}

func (p *kCRUDClient) QueryIndicator(ctx context.Context, req *api.QueryIndicatorRequest, callOptions ...callopt.Option) (r *api.QueryIndicatorResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryIndicator(ctx, req)
}

func (p *kCRUDClient) ModifyIndicator(ctx context.Context, req *api.ModifyIndicatorRequest, callOptions ...callopt.Option) (r *api.ModifyIndicatorResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ModifyIndicator(ctx, req)
}

func (p *kCRUDClient) AddRule(ctx context.Context, req *api.AddRuleRequest, callOptions ...callopt.Option) (r *api.AddRuleResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddRule(ctx, req)
}

func (p *kCRUDClient) DeleteRule(ctx context.Context, req *api.DeleteRuleRequest, callOptions ...callopt.Option) (r *api.DeleteRuleResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteRule(ctx, req)
}

func (p *kCRUDClient) QueryRule(ctx context.Context, req *api.QueryRuleRequest, callOptions ...callopt.Option) (r *api.QueryRuleResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryRule(ctx, req)
}

func (p *kCRUDClient) ModifyRule(ctx context.Context, req *api.ModifyRuleRequest, callOptions ...callopt.Option) (r *api.ModifyRuleResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ModifyRule(ctx, req)
}

func (p *kCRUDClient) AddTask(ctx context.Context, req *api.AddTaskRequest, callOptions ...callopt.Option) (r *api.AddTaskResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddTask(ctx, req)
}

func (p *kCRUDClient) DeleteTask(ctx context.Context, req *api.DeleteTaskRequest, callOptions ...callopt.Option) (r *api.DeleteTaskResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteTask(ctx, req)
}

func (p *kCRUDClient) QueryTask(ctx context.Context, req *api.QueryTaskRequest, callOptions ...callopt.Option) (r *api.QueryTaskResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryTask(ctx, req)
}

func (p *kCRUDClient) ModifyTask(ctx context.Context, req *api.ModifyTaskRequest, callOptions ...callopt.Option) (r *api.ModifyTaskResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ModifyTask(ctx, req)
}
