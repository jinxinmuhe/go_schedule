package task

import (
	"context"
	"fmt"

	"go_schedule/client"
	"go_schedule/dao/mongodb"
	"go_schedule/logic/consistent_hash"
	"go_schedule/logic/schedule"
	pb "go_schedule/task_management"
	"go_schedule/util/config"
	"go_schedule/util/data_schema"
	"go_schedule/util/log"
	"go_schedule/util/tool"

	"go.mongodb.org/mongo-driver/bson"
)

// DeleteTask 创建任务
func DeleteTask(ctx context.Context, req *pb.DeleteTaskReq) (*pb.DeleteTaskResp, error) {
	resp := pb.DeleteTaskResp{}
	// 若不是当前节点负责的task，则进行转发
	if !schedule.TaskExists(req.GetTaskId()) {
		return redirectDeleteReq(ctx, req)
	}
	task := data_schema.TaskInfo{
		TaskID: req.GetTaskId(),
	}
	schedule.ScheLock()
	defer schedule.ScheUnLock()
	if schedule.LenDeleteChan() >= config.Viper.GetInt("schedule.delete_size") {
		log.Errorf("delete task too frequency")
		resp.Code = pb.RespCode_FAIL
		resp.Msg = fmt.Sprintf("delete task too frequency")
		return &resp, fmt.Errorf("delete task too frequency")
	}

	// 写mongodb
	filter := bson.D{
		{"task_id", req.GetTaskId()},
	}
	if _, err := mongodb.DeleteTask(ctx, filter); err != nil {
		log.Errorf("delete task from mongodb fail, error:%v, filter:%+v", err, filter)
		resp.Code = pb.RespCode_FAIL
		resp.Msg = fmt.Sprintf("delete task from mongodb fail")
		return &resp, fmt.Errorf("delete task from mongodb fail")
	}

	// 写zk
	// pathTask := fmt.Sprintf("/go_schedule/task/%s", req.GetTaskId())
	// zookeeper.DeleteNode(pathTask)

	schedule.DeleteTask(task)
	resp.Code = pb.RespCode_SUCC
	resp.Msg = "succ"
	return &resp, nil
}

// func deleteZkNode(taskID string) {
// 	pathIP := fmt.Sprintf("/go_schedule/schedule/%s/%s", tool.IP, taskID)
// 	if err := zookeeper.DeleteNode(pathIP); err != nil {
// 		log.Errorf("delete node fail, path:%s, error:%+v", pathIP, err)
// 		return err
// 	}
// 	pathTask := fmt.Sprintf("/go_schedule/task/%s", taskID)
// 	if err := zookeeper.DeleteNode(pathTask); err != nil {
// 		log.Errorf("delete node fail, path:%s, error:%+v", pathTask, err)
// 		return err
// 	}
// 	return nil
// }

func redirectDeleteReq(ctx context.Context, req *pb.DeleteTaskReq) (*pb.DeleteTaskResp, error) {
	resp := pb.DeleteTaskResp{
		TaskId: req.GetTaskId(),
	}
	// path := fmt.Sprintf("/go_schedule/task/%s", req.GetTaskId())
	// ipByte, err := zookeeper.GetData(path)
	// if err != nil {
	// 	log.Errorf("delete task %s fail, cannt find the specific ip, detail msg:%+v", req.GetTaskId(), err)
	// 	resp.Code = pb.RespCode_FAIL
	// 	resp.Msg = "delete task fail"
	// 	return &resp, err
	// }

	taskMd5, err := consistent_hash.HashCalculation(req.GetTaskId())
	if err != nil {
		log.Errorf("delete task %s fail, cannt find the specific ip, detail msg:%+v", req.GetTaskId(), err)
		resp.Code = pb.RespCode_FAIL
		resp.Msg = "delete task fail"
		return &resp, err
	}
	ip := consistent_hash.SelectIP(taskMd5)
	if ip == "" {
		log.Errorf("delete task %s fail, the found ip is empty", req.GetTaskId())
		resp.Code = pb.RespCode_FAIL
		resp.Msg = "delete task fail"
		return &resp, fmt.Errorf("ip is empty")
	}
	if ip == tool.IP {
		log.Errorf("delete task %s fail, the ip registered in zk is same with current node", req.GetTaskId())
		resp.Code = pb.RespCode_FAIL
		resp.Msg = "delete task fail"
		return &resp, err
	}

	return client.DeleteTaskClient(ctx, req, fmt.Sprintf("%s:%d", ip, config.Viper.GetInt("port")))
}
