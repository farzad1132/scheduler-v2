// Code generated by goctl. DO NOT EDIT.
// Source: worker.proto

package server

import (
	"context"

	"worker/internal/logic"
	"worker/internal/svc"
	"worker/worker"
)

type WorkerServer struct {
	svcCtx *svc.ServiceContext
	worker.UnimplementedWorkerServer
}

func NewWorkerServer(svcCtx *svc.ServiceContext) *WorkerServer {
	return &WorkerServer{
		svcCtx: svcCtx,
	}
}

func (s *WorkerServer) FinishWorker(ctx context.Context, in *worker.TaskFinishRequest) (*worker.TaskFinishResponse, error) {
	l := logic.NewFinishWorkerLogic(ctx, s.svcCtx)
	return l.FinishWorker(in)
}
