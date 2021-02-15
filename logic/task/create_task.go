package task

import (
	"context"
	"fmt"
	"math/rand"
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
)

// CreateTask 创建任务
func CreateTask(ctx context.Context, req *pb.CreateTaskReq) (*pb.CreateTaskResp, error) {
	resp := pb.CreateTaskResp{}
	taskID := req.GetTaskId()
	if taskID == "" {
		taskID = fmt.Sprintf("%xt%x", time.Now().UnixNano(), rand.Int31n(8192))
		req.TaskId = taskID
	}
	goalIP, err := selectIP(taskID)
	if goalIP != tool.IP {
		// 转发任务创建
		return client.CreateTaskClient(ctx, req, fmt.Sprintf("%s:%d", goalIP, config.Viper.GetInt("port")))
	}

	// 参数及内存环境校验，防止内存任务创建失败
	if isvalid, err := paramValid(ctx, req); !isvalid {
		log.Errorf("parameters is invalid, task:%+v, error:%+v", *req, err)
		resp.Code = pb.RespCode_FAIL
		resp.Msg = "parameter is invalid"
		return &resp, err
	}
	schedule.ScheLock()
	defer schedule.ScheUnLock()
	if schedule.LenCreateChan() >= config.Viper.GetInt("schedule.create_size") {
		log.Errorf("create task too frequency")
		resp.Code = pb.RespCode_FAIL
		resp.Msg = "parameter is invalid"
		return &resp, fmt.Errorf("create task too frequency")
	}
	task := data_schema.TaskInfo{
		TaskID:     taskID,
		TaskName:   req.GetTaskName(),
		Owner:      req.GetOwner(),
		CronInfo:   req.GetScheduleTime(),
		KafkaTopic: req.GetKafkaTopic(),
		AlarmMail:  req.GetAlarmEmail(),
		Variables:  req.GetEnvVariable(),
		CreateTime: time.Now().In(tool.TimeLocal),
		UpdateTime: time.Now().In(tool.TimeLocal),
	}

	// 写mongodb
	id, err := mongodb.InsertOneTask(ctx, task)
	if err != nil {
		log.Errorf("create task fail, write mongodb fail, task:%+v, error:%+v", task, err)
		resp.Code = pb.RespCode_FAIL
		resp.Msg = "write mongodb fail"
		return &resp, err
	}
	log.Infof("write mongodb succ, taskID:%s, mongoID:%s", task.TaskID, id)

	// 写zookeeper
	// pathTask := fmt.Sprintf("/go_schedule/task/%s", task.TaskID)
	// zookeeper.CreateTemplateNode(pathTask, []byte(tool.IP))
	// log.Infof("write zk succ, taskID: %s", task.TaskID)

	// 写内存
	schedule.CreateTask(task)
	resp.Code = pb.RespCode_SUCC
	resp.Msg = "succ"
	return &resp, nil
}

// func writeZk(taskID string) error {
// 	pathIPParent := fmt.Sprintf("/go_schedule/schedule/%s", tool.IP)
// 	if exist, err := zookeeper.ExistNode(pathIPParent); err != nil {
// 		log.Errorf("cannt judge zk node is exist, path:%s, error:%+v", pathIPParent, err)
// 		return err
// 	} else if !exist {
// 		if _, err := zookeeper.CreateNode(pathIPParent, nil); err != nil {
// 			log.Errorf("cannt create parent zk node, path:%s, error:%+v", pathIPParent, err)
// 			return err
// 		}
// 	}

// 	pathIP := fmt.Sprintf("/go_schedule/schedule/%s/%s", tool.IP, taskID)
// 	if _, err := zookeeper.CreateNode(pathIP, []byte(taskID)); err != nil {
// 		log.Errorf("create task fail, write zk fail, path:%s, data:%s, error:%s", pathIP, taskID, err.Error())
// 		return err
// 	}
// 	pathTask := fmt.Sprintf("/go_schedule/task/%s", taskID)
// 	if _, err := zookeeper.CreateNode(pathTask, []byte(tool.IP)); err != nil {
// 		log.Errorf("create task fail, write zk fail, path:%s, data:%s, error:%s", pathTask, tool.IP, err.Error())
// 		return err
// 	}
// 	return nil
// }

func paramValid(ctx context.Context, req *pb.CreateTaskReq) (bool, error) {
	if req.TaskName == "" || req.Owner == "" || req.KafkaTopic == "" {
		return false, fmt.Errorf("miss required parameters")
	}

	_, err := cron.Parse(req.ScheduleTime)
	if err != nil {
		log.Errorf("schedule time is illegal, taskInfo:%+v, error:%+v", *req, err)
		return false, err
	}

	return true, nil
}

func selectIP(taskID string) (ip string, err error) {
	taskMd5, err := consistent_hash.HashCalculation(taskID)
	if err != nil {
		log.Errorf("create task, calculate taskID:%s md5 fail, error:%+v", taskID, err)
		return "", err
	}
	ip = consistent_hash.SelectIP(taskMd5)
	if ip == "" {
		log.Errorf("create task:%s, the found ip is empty", taskID)
		return "", fmt.Errorf("ip is empty")
	}
	return ip, nil
}
