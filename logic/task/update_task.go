package task

import (
	"context"
	"fmt"
	"time"

	"go_schedule/client"
	"go_schedule/dao/mongodb"
	"go_schedule/logic/consistent_hash"
	"go_schedule/logic/schedule"
	pb "go_schedule/task_management"
	"go_schedule/util/config"
	"go_schedule/util/data_schema"
	"go_schedule/util/log"
	"go_schedule/util/tool"

	"github.com/robfig/cron"
	"go.mongodb.org/mongo-driver/bson"
)

// UpdateTask 创建任务
func UpdateTask(ctx context.Context, req *pb.UpdateTaskReq) (*pb.UpdateTaskResp, error) {
	// 若不是当前节点负责的task，则进行转发
	if !schedule.TaskExists(req.GetTaskId()) {
		return redirectUpdateReq(ctx, req)
	}

	resp := pb.UpdateTaskResp{}
	if isvalid, err := paramValidUpdate(ctx, req); !isvalid {
		log.Errorf("parameters is invalid, request:%+v, error:%+v", *req, err)
		resp.Code = pb.RespCode_FAIL
		resp.Msg = "parameter is invalid"
		return &resp, err
	}
	schedule.ScheLock()
	defer schedule.ScheUnLock()
	if schedule.LenUpdateChan() >= config.Viper.GetInt("schedule.update_size") {
		log.Errorf("update task too frequency")
		resp.Code = pb.RespCode_FAIL
		resp.Msg = "parameter is invalid"
		return &resp, fmt.Errorf("update task too frequency")
	}
	taskB := bson.D{
		{"$set", bson.D{
			{"task_name", req.GetName()},
			{"cron_info", req.GetScheduleTime()},
			{"kafka_topic", req.GetKafkaTopic()},
			{"alarm_mail", req.GetAlarmEmail()},
			{"variables", req.GetEnvVariable()},
			{"update_time", time.Now()}},
		},
	}

	// 更新db
	filter := bson.D{{"task_id", req.GetTaskId()}}
	if _, err := mongodb.UpdateOneTask(ctx, filter, taskB); err != nil {
		log.Errorf("update task fail, write mongodb fail, task:%+v, error:%+v", taskB, err)
		resp.Code = pb.RespCode_FAIL
		resp.Msg = "write mongodb fail"
		return &resp, err
	}
	log.Infof("write mongodb succ, taskID:%s", req.GetTaskId())

	task := data_schema.TaskInfo{
		TaskID:     req.GetTaskId(),
		TaskName:   req.GetName(),
		CronInfo:   req.GetScheduleTime(),
		KafkaTopic: req.GetKafkaTopic(),
		AlarmMail:  req.GetAlarmEmail(),
		Variables:  req.GetEnvVariable(),
		UpdateTime: time.Now().In(tool.TimeLocal),
	}
	schedule.UpdateTask(task)
	resp.Code = pb.RespCode_SUCC
	resp.Msg = "succ"
	return &resp, nil
}

func paramValidUpdate(ctx context.Context, req *pb.UpdateTaskReq) (bool, error) {
	if req.Name == "" || req.KafkaTopic == "" {
		return false, fmt.Errorf("miss required parameters")
	}

	_, err := cron.Parse(req.ScheduleTime)
	if err != nil {
		log.Errorf("schedule time is illegal, taskInfo:%+v, error:%+v", *req, err)
		return false, err
	}

	return true, nil
}

func redirectUpdateReq(ctx context.Context, req *pb.UpdateTaskReq) (*pb.UpdateTaskResp, error) {
	resp := pb.UpdateTaskResp{}
	// path := fmt.Sprintf("/go_schedule/task/%s", req.GetTaskId())
	// ipByte, err := zookeeper.GetData(path)
	// if err != nil {
	// 	log.Errorf("update task %s fail, cannt find the specific ip, detail msg:%+v", req.GetTaskId(), err)
	// 	resp.Code = pb.RespCode_FAIL
	// 	resp.Msg = "update task fail"
	// 	return &resp, err
	// }

	taskMd5, err := consistent_hash.HashCalculation(req.GetTaskId())
	if err != nil {
		log.Errorf("update task %s fail, cannt find the specific ip, detail msg:%+v", req.GetTaskId(), err)
		resp.Code = pb.RespCode_FAIL
		resp.Msg = "update task fail"
		return &resp, err
	}
	ip := consistent_hash.SelectIP(taskMd5)
	if ip == "" {
		log.Errorf("update task %s fail, the found ip is empty", req.GetTaskId())
		resp.Code = pb.RespCode_FAIL
		resp.Msg = "update task fail"
		return &resp, fmt.Errorf("ip is empty")
	}
	if ip == tool.IP {
		log.Errorf("update task %s fail, the ip registered in zk is same with current node", req.GetTaskId())
		resp.Code = pb.RespCode_FAIL
		resp.Msg = "update task fail"
		return &resp, err
	}

	return client.UpdateTaskClient(ctx, req, fmt.Sprintf("%s:%d", ip, config.Viper.GetInt("port")))
}
