package client

import (
	"context"
	"fmt"
	pb "go_schedule/task_management"
	"go_schedule/util/tool"
	"testing"
)

func TestUpdateTask(t *testing.T) {
	ctx := context.Background()
	req := pb.UpdateTaskReq{
		TaskId: "1663dcf85c1742f0t1f25",
		Name: "test_1",
		ScheduleTime: "*/5 * * * * *",
		KafkaTopic: "kafkatopic",
		AlarmEmail: "xxx@qq.com",
		Owner: "jasonjinxin",
	}
	if _, err := UpdateTaskClient(ctx, &req, fmt.Sprintf("%s:12343", tool.IP)); err != nil {
		t.Fatalf("update fail, error:%+v", err)
	}
}

func TestCreateTask(t *testing.T) {
	ctx := context.Background()
	req := pb.CreateTaskReq{
		TaskName: "test_5",
		ScheduleTime: "*/5 * * * * *",
		KafkaTopic: "kafkatopic",
		AlarmEmail: "xxx@qq.com",
		Owner: "jasonjinxin",
	}
	if _, err := CreateTaskClient(ctx, &req, fmt.Sprintf("%s:12343", tool.IP)); err != nil {
		t.Fatalf("create fail, error:%+v", err)
	}
}
