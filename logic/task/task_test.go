package task

import (
	"context"
	"testing"

	"go_schedule/dao/mongodb"
	"go_schedule/dao/zookeeper"
	"go_schedule/logic/schedule"
	"go_schedule/task_management"
	"go_schedule/util/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	config.InitConfig("")
	mongodb.InitMongodb()
	zookeeper.InitZookeeper()
	schedule.InitSchedule()
	go schedule.Start(context.Background())
}

func TestSearchTask(t *testing.T) {
	ctx := context.Background()
	req := task_management.SearchTaskReq{
		TaskId: "",
	}
	if resp, err := SearchTask(ctx, &req); err != nil {
		t.Fatalf("error:%+v", err)
	} else {
		t.Logf("resp:%+v", *resp)
	}
}

func TestCreateTask(t *testing.T) {
	ctx := context.Background()
	req := task_management.CreateTaskReq{
		TaskName:     "test",
		ScheduleTime: "*/2 * * * * *",
		KafkaTopic:   "kafka_topic",
		AlarmEmail:   "xx@qq.com",
		Owner:        "jasonjinxin",
	}
	resp, err := CreateTask(ctx, &req)
	if err != nil {
		t.Fatalf("error:%+v", err)
	}
	t.Logf("resp:%+v", *resp)
	// time.Sleep(30 * time.Second)
}

func TestUpdateTask(t *testing.T) {
	ctx := context.Background()
	req := task_management.UpdateTaskReq{
		TaskId:       "6027bfe60a1e48befd73ffd1",
		Name:         "test_up",
		ScheduleTime: "*/3 * * * * *",
		KafkaTopic:   "kafka_topic",
		AlarmEmail:   "xx@qq.com",
		Owner:        "jasonjinxin",
	}
	resp, err := UpdateTask(ctx, &req)
	if err != nil {
		t.Fatalf("error:%+v", err)
	}
	t.Logf("resp:%+v", *resp)
}

func TestDeleteTask(t *testing.T) {
	ctx := context.Background()
	req := task_management.DeleteTaskReq{
		TaskId: "000000000000000000000000",
	}
	resp, err := DeleteTask(ctx, &req)
	if err != nil {
		t.Fatalf("error:%+v", err)
	}
	t.Logf("resp:%+v", *resp)
}

func TestObjectID(t *testing.T) {
	o1,_ := primitive.ObjectIDFromHex("6027bfe60a1e48befd73ffd1")
	o2,_ := primitive.ObjectIDFromHex("6027bfe60a1e48befd73ffd1")
	if o1 == o2 {
		t.Logf("true----")
	}
}
