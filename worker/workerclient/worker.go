// Code generated by goctl. DO NOT EDIT.
// Source: worker.proto

package workerclient

import (
	"context"

	"worker/worker"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	TaskFinishRequest  = worker.TaskFinishRequest
	TaskFinishResponse = worker.TaskFinishResponse

	Worker interface {
		FinishWorker(ctx context.Context, in *TaskFinishRequest, opts ...grpc.CallOption) (*TaskFinishResponse, error)
	}

	defaultWorker struct {
		cli zrpc.Client
	}
)

func NewWorker(cli zrpc.Client) Worker {
	return &defaultWorker{
		cli: cli,
	}
}

func (m *defaultWorker) FinishWorker(ctx context.Context, in *TaskFinishRequest, opts ...grpc.CallOption) (*TaskFinishResponse, error) {
	client := worker.NewWorkerClient(m.cli.Conn())
	return client.FinishWorker(ctx, in, opts...)
}