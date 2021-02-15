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
		TaskId: "1234",
		Name: "test",
		ScheduleTime: "*/5 * * * * *",
		KafkaTopic: "kafkatopic",
		AlarmEmail: "xxx@qq.com",
		Owner: "jasonjinxin",
	}
	if _, err := UpdateTaskClient(ctx, &req, fmt.Sprintf("%s:12343", tool.IP)); err != nil {
		t.Fatalf("update fail, error:%+v", err)
	}
}
