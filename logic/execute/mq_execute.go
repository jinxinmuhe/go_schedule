package execute

import (
	"context"
	"go_schedule/util/data_schema"
	"go_schedule/util/log"
)

// MqExecute 通过kafka将任务通知给执行器
func MqExecute(ctx context.Context, task data_schema.TaskInfo) {
	log.InfoLogger.Printf("execute succ, task name:%s", task.TaskName)
}
