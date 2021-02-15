package mongodb

import (
	"context"
	"go_schedule/util/config"
	"go_schedule/util/data_schema"
	"reflect"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func init() {
	config.InitConfig("")
	InitMongodb()
}

func TestInsertOneTask(t *testing.T) {
	ctx := context.Background()
	cases := []data_schema.TaskInfo{
		{TaskID: "1234", TaskName: "tes2", CronInfo: "1 1 2 * *", CreateTime: time.Now()},
	}

	for _, c := range cases {
		id, err := InsertOneTask(ctx, c)
		if err != nil {
			t.Fatalf("err:%+v", err)
		}
		t.Logf("type:%+v, id:%+v", reflect.TypeOf(id), id)
	}
}

func TestUpdateOneTask(t *testing.T) {
	ctx := context.Background()
	filterCases := []bson.D{
		{{"taskid", 1234}},
	}
	updateCases := []bson.D{
		{{"$set", bson.D{{"taskname", "test1"}}}},
	}

	result, err := UpdateOneTask(ctx, filterCases[0], updateCases[0])
	if err != nil {
		t.Fatalf("error:%+v", err)
	}
	t.Logf("result:%+v", result)
}

func TestSearchTasks(t *testing.T) {
	ctx := context.Background()
	filterCases := []bson.D{
		{{"task_id", 1234}},
	}

	if tasks, err := SearchTask(ctx, filterCases[0]); err != nil {
		t.Fatalf("error:%+v", err)
	} else {
		t.Logf("result:%+v", tasks)
	}
}
