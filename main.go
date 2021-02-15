package main

import (
	"context"
	"flag"
	"fmt"
	"net"

	"go_schedule/dao/mongodb"
	"go_schedule/dao/zookeeper"
	"go_schedule/logic/schedule"
	pb "go_schedule/task_management"
	"go_schedule/util/config"
	"go_schedule/util/log"

	"google.golang.org/grpc"
)

var configPath = flag.String("config", "./conf/app.toml", "配置文件地址")

func main() {
	flag.Parse()
	log.Infof("start")
	initServer(*configPath)
	go schedule.Start(context.Background())

	// 初始化grpc
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Viper.GetInt("port")))
	if err != nil {
		log.Errorf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTaskManagementServer(s, &taskServer{})
	if err := s.Serve(lis); err != nil {
		log.Errorf("failed to serve: %v", err)
	}
}

func initServer(path string) {
	config.InitConfig(path)
	mongodb.InitMongodb()
	zookeeper.InitZookeeper()
	schedule.InitSchedule()
}
