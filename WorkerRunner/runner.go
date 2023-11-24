package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
	"workerrunner/worker"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func sendRPC(msg string) {
	logx.Infof("Sending RPC")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "127.0.0.1:8080", grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	//conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logx.Error(err)
		return
	} else {
		logx.Info("Dialed successfully")
	}
	client := worker.NewWorkerClient(conn)
	resp, err := client.FinishWorker(context.Background(), &worker.TaskFinishRequest{
		Finish: msg,
	})
	if err != nil {
		logx.Error(err)
		return
	}
	logx.Info(resp)

}

func main() {
	args := os.Args
	if len(args) != 3 {
		logx.Error("Not enough arguments")
		panic(1)
	}

	name := args[1]
	str_comp := args[2]

	logx.Infof("Starting Task: %v\n", str_comp)

	str_flot, err := strconv.ParseFloat(str_comp, 64)
	if err != nil {
		logx.Error("Could not convert computation")
		panic(1)
	}

	// fail with probability of 0.3
	if rand.Float64() <= 0.3 {
		<-time.After(time.Duration(3 * rand.Float64() * float64(time.Second)))
		sendRPC(fmt.Sprintf("Task %v failed.\n", name))
	} else {
		<-time.After(time.Duration(str_flot * float64(time.Second)))
		sendRPC(fmt.Sprintf("Task %v completed.\n", name))
	}
}
