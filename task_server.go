package main

import (
	"context"
	"fmt"

	"go_schedule/logic/task"
	pb "go_schedule/task_management"
	"go_schedule/util/log"
)

type taskServer struct {
	pb.UnimplementedTaskManagementServer
}

func (s *taskServer) CreateTask(ctx context.Context, in *pb.CreateTaskReq) (*pb.CreateTaskResp, error) {
	if in == nil {
		log.ErrLogger.Printf("create task request is nil")
		return nil, fmt.Errorf("request is nil")
	}
	log.InfoLogger.Printf("CreateTask Req: %v", *in)
	return task.CreateTask(ctx, in)
}

func (s *taskServer) UpdateTask(ctx context.Context, in *pb.UpdateTaskReq) (*pb.UpdateTaskResp, error) {
	if in == nil {
		log.ErrLogger.Printf("update task request is nil")
		return nil, fmt.Errorf("request is nil")
	}
	log.InfoLogger.Printf("UpdateTask Req: %v", *in)
	return task.UpdateTask(ctx, in)
}

func (s *taskServer) DeleteTask(ctx context.Context, in *pb.DeleteTaskReq) (*pb.DeleteTaskResp, error) {
	if in == nil {
		log.ErrLogger.Printf("delete task request is nil")
		return nil, fmt.Errorf("request is nil")
	}
	log.InfoLogger.Printf("DeleteTask Req: %v", *in)
	return task.DeleteTask(ctx, in)
}

func (s *taskServer) SearchTask(ctx context.Context, in *pb.SearchTaskReq) (*pb.SearchTaskResp, error) {
	return task.SearchTask(ctx, in)
}
