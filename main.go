package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	"go_schedule/dao/mongodb"
	"go_schedule/dao/zookeeper"
	"go_schedule/logic/consistent_hash"
	"go_schedule/logic/schedule"
	pb "go_schedule/task_management"
	"go_schedule/util/config"
	"go_schedule/util/log"

	"google.golang.org/grpc"
)

var configPath = flag.String("config", "./conf/app.toml", "配置文件地址")

func main() {
	flag.Parse()
	log.InfoLogger.Printf("start")
	ctx := context.Background()
	initServer(ctx, *configPath)
	go schedule.Start(ctx)
	schedule.InitEntry()

	// 初始化grpc
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Viper.GetInt("port")))
	if err != nil {
		log.ErrLogger.Printf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTaskManagementServer(s, &taskServer{})
	if err := s.Serve(lis); err != nil {
		log.ErrLogger.Printf("failed to serve: %v", err)
	}
}

func initServer(ctx context.Context, path string) {
	if err := config.InitConfig(path); err != nil {
		panic(err)
	}
	if err := mongodb.InitMongodb(); err != nil {
		panic(err)
	}
	if err := zookeeper.InitZookeeper(); err != nil {
		panic(err)
	}
	if err := consistent_hash.InitIPMd5List(); err != nil {
		panic(err)
	}
	if err := schedule.InitSchedule(ctx); err != nil {
		panic(err)
	}
}
