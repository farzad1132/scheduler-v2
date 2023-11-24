package main

import (
	"context"
	"flag"
	"fmt"

	"worker/internal/config"
	"worker/internal/mqs"
	"worker/internal/server"
	"worker/internal/svc"
	"worker/worker"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/worker.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	group := service.NewServiceGroup()

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		worker.RegisterWorkerServer(grpcServer, server.NewWorkerServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	group.Add(s)

	defer group.Stop()

	for _, mq := range mqs.Consumers(c, context.Background(), ctx) {
		group.Add(mq)
	}

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	group.Start()
}
