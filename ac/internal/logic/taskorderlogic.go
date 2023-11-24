package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"ac/internal/svc"
	"ac/internal/types"

	"github.com/zeromicro/go-zero/core/logc"
	"github.com/zeromicro/go-zero/core/logx"
)

type TaskOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskOrderLogic {
	return &TaskOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskOrderLogic) TaskOrder(req *types.TaskOrderRequest) (resp *types.TaskOrderResponse, err error) {
	// pushing TaskOrder to kafka
	body, e := json.Marshal(req)
	if e != nil {
		logc.Errorf(l.ctx, "Could not marshal this message: %v", req)
	}
	if err := l.svcCtx.KqPusherClient.Push(string(body)); err != nil {
		logc.Errorf(l.ctx, "KqPusherClient Push Error , err :%v", err)
	}
	logc.Infof(l.ctx, "Pushed to kafka: 'Computation:%v'", req.Computation)
	resp = &types.TaskOrderResponse{}
	resp.Ack = fmt.Sprintf("Order %v accepted.", req)
	return
}
