package logic

import (
	"context"

	"worker/internal/svc"
	"worker/worker"

	"github.com/zeromicro/go-zero/core/logx"
)

type FinishWorkerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFinishWorkerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FinishWorkerLogic {
	return &FinishWorkerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FinishWorkerLogic) FinishWorker(in *worker.TaskFinishRequest) (*worker.TaskFinishResponse, error) {
	logx.Infof("msg from worker: %v", in.Finish)

	return &worker.TaskFinishResponse{
		Ack: "Ok!!!",
	}, nil
}
