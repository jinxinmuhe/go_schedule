package client

import (
	"context"

	"time"

	pb "go_schedule/task_management"
	"go_schedule/util/log"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

// CreateTaskClient 创建任务
func CreateTaskClient(ctx context.Context, req *pb.CreateTaskReq, address string) (*pb.CreateTaskResp, error) {
	resp := pb.CreateTaskResp{}
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Errorf("connect to %s fail, error:%+v", address, err)
		resp.Code = pb.RespCode_FAIL
		resp.Msg = "connect fail"
		return &resp, err
	}
	defer conn.Close()
	c := pb.NewTaskManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateTask(ctx, req)
	if err != nil {
		log.Errorf("request %s for creating task fail, error:%+v", address, err)
	}
	return r, err
}

// UpdateTaskClient 创建任务
func UpdateTaskClient(ctx context.Context, req *pb.UpdateTaskReq, address string) (*pb.UpdateTaskResp, error) {
	resp := pb.UpdateTaskResp{}
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Errorf("connect to %s fail, error:%+v", address, err)
		resp.Code = pb.RespCode_FAIL
		resp.Msg = "connect fail"
		return &resp, err
	}
	defer conn.Close()
	c := pb.NewTaskManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UpdateTask(ctx, req)
	if err != nil {
		log.Errorf("request %s for updating task fail, error:%+v", address, err)
	}
	return r, err
}

// DeleteTaskClient 创建任务
func DeleteTaskClient(ctx context.Context, req *pb.DeleteTaskReq, address string) (*pb.DeleteTaskResp, error) {
	resp := pb.DeleteTaskResp{}
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Errorf("connect to %s fail, error:%+v", address, err)
		resp.Code = pb.RespCode_FAIL
		resp.Msg = "connect fail"
		return &resp, err
	}
	defer conn.Close()
	c := pb.NewTaskManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.DeleteTask(ctx, req)
	if err != nil {
		log.Errorf("request %s for deleting task fail, error:%+v", address, err)
	}
	return r, err
}
