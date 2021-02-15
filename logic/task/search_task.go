package task

import (
	"context"

	"go_schedule/dao/mongodb"
	pb "go_schedule/task_management"
	"go_schedule/util/log"

	"go.mongodb.org/mongo-driver/bson"
)

// SearchTask 查询任务
func SearchTask(ctx context.Context, in *pb.SearchTaskReq) (*pb.SearchTaskResp, error) {
	resp := pb.SearchTaskResp{}
	keys := make([]string, 0)
	values := make([]interface{}, 0)
	if in.GetKafkaTopic() != "" {
		keys = append(keys, "kafka_topic")
		values = append(values, in.GetKafkaTopic())
	}
	if in.GetOwner() != "" {
		keys = append(keys, "owner")
		values = append(values, in.GetOwner())
	}
	if in.GetTaskId() != "" {
		keys = append(keys, "_id")
		values = append(values, in.GetTaskId())
	}

	conditions := make([]bson.E, 0, len(keys))
	for i, k := range keys {
		conditions = append(conditions, bson.E{k, values[i]})
	}

	tasks, err := mongodb.SearchTask(ctx, conditions)
	if err != nil {
		resp.Code = pb.RespCode_FAIL
		resp.Msg = "inner error"
		log.ErrLogger.Printf("search task fail, error:%+v, conditions:%+v", err, conditions)
		return &resp, err
	}

	pbTasks := make([]*pb.TaskInfo, 0, len(tasks))
	for _, t := range tasks {
		pt := pb.TaskInfo{
			TaskId:       t.TaskID,
			TaskName:     t.TaskName,
			ScheduleTime: t.CronInfo,
			KafkaTopic:   t.KafkaTopic,
			AlarmEmail:   t.AlarmMail,
			Owner:        t.Owner,
			EnvVariable:  t.Variables,
		}
		pbTasks = append(pbTasks, &pt)
	}
	resp.Code = pb.RespCode_SUCC
	resp.Msg = "succ"
	resp.Tasks = pbTasks
	return &resp, nil
}
