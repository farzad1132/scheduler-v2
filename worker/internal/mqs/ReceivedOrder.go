package mqs

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
	"worker/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReceivedOrder struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReceivedOrder(ctx context.Context, svcCtx *svc.ServiceContext) *ReceivedOrder {
	return &ReceivedOrder{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type TaskOrderRequest struct {
	Computation float64 `form:"computation"`
}

func (l *ReceivedOrder) Consume(key, val string) error {
	logx.Infof("Received Order key :%s , val :%s", key, val)

	var request TaskOrderRequest

	err := json.Unmarshal([]byte(val), &request)
	if err != nil {
		logx.Error(err)
	}

	// Start Worker
	app := "./workerrunner"

	arg0 := fmt.Sprintf("Start-%v", time.Now().Format(time.TimeOnly))
	arg1 := fmt.Sprintf("%v", request.Computation)

	cmd := exec.Command(app, arg0, arg1)

	if err := cmd.Start(); err != nil {
		logx.Errorf("Error running task: %v", err)
	} else {
		logx.Info("Worker started successfully.")
	}
	return nil
}
