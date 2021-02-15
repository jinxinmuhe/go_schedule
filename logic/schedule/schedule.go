package schedule

import (
	"context"
	"fmt"
	"sort"
	"time"

	"go_schedule/dao/mongodb"
	"go_schedule/logic/execute"
	"go_schedule/util/config"
	"go_schedule/util/data_schema"
	"go_schedule/util/log"

	"github.com/robfig/cron"
)

var sd data_schema.Schedule

// InitSchedule 初始化调度
func InitSchedule() {
	sd.Create = make(chan data_schema.TaskInfo, config.Viper.GetInt("schedule.create_size"))
	sd.Update = make(chan data_schema.TaskInfo, config.Viper.GetInt("schedule.update_size"))
	sd.Delete = make(chan data_schema.TaskInfo, config.Viper.GetInt("schedule.delete_size"))
	sd.Entries = make([]data_schema.Entry, 0, 4)
}

func LenCreateChan() int {
	return len(sd.Create)
}

func LenUpdateChan() int {
	return len(sd.Update)
}

func LenDeleteChan() int {
	return len(sd.Delete)
}

// TaskExists 判断是不是存在对应的任务
func TaskExists(taskID string) bool {
	for _, t := range sd.Entries {
		if t.Task.TaskID == taskID {
			return true
		}
	}
	return false
}

// CreateTask 添加任务
func CreateTask(task data_schema.TaskInfo) {
	sd.Create <- task
}

// UpdateTask 更新任务
func UpdateTask(task data_schema.TaskInfo) {
	sd.Update <- task
}

// DeleteTask 删除任务
func DeleteTask(task data_schema.TaskInfo) {
	sd.Delete <- task
}

// Start 启动调度
func Start(ctx context.Context) {
	for {
		d := 5 * time.Minute
		sort.Sort(sd.Entries)
		if len(sd.Entries) > 0 && !sd.Entries[0].Next.IsZero() {
			d = sd.Entries[0].Next.Sub(time.Now())
		}
		t := time.NewTimer(d)
		select {
		case <-t.C:
			executeTask(ctx)
		case task := <-sd.Create:
			doCreate(ctx, task)
		case task := <-sd.Delete:
			doDelete(ctx, task)
		case task := <-sd.Update:
			doUpdate(ctx, task)
		}
	}
}

func executeTask(ctx context.Context) {
	now := time.Now()
	for i, entry := range sd.Entries {
		if entry.Next.Before(now) {
			doExecute(ctx, entry.Task)
			sd.Entries[i].Prev = now
			sd.Entries[i].Next = sd.Entries[i].Schedule.Next(now)
			// insert schedule log to mongodb
			s := data_schema.ScheduleHistory{
				TaskID:     entry.Task.TaskID,
				ScheduleID: now.Unix(),
				NextTime:   sd.Entries[i].Next,
				Variables:  entry.Task.Variables,
				CreateTime: now,
				UpdateTime: now,
			}
			id, err := mongodb.InsertOneSchedule(ctx, s)
			if err != nil {
				log.Errorf("insert schedule into mongo fail, schedule:%+v, task:%+v, error:%+v", s, entry.Task, err)
			} else {
				log.Infof("scheduleID:%+v", id)
			}

		}
	}
}

func doExecute(ctx context.Context, task data_schema.TaskInfo) {
	if task.KafkaTopic != "" {
		execute.MqExecute(ctx, task)
	}
}

func doCreate(ctx context.Context, task data_schema.TaskInfo) {
	s, err := cron.Parse(task.CronInfo)
	if err != nil {
		log.Errorf("parse task cron info fail, taskID:%d, cronInfo:%s, error:%+v", task.TaskID, task.CronInfo, err)
		return
	}
	e := data_schema.Entry{
		Task:     task,
		Schedule: s,
		Next:     s.Next(time.Now()),
	}
	ScheLock()
	defer ScheUnLock()
	sd.Entries = append(sd.Entries, e)
}

func doUpdate(ctx context.Context, task data_schema.TaskInfo) error {
	for i, t := range sd.Entries {
		if t.Task.TaskID == task.TaskID {
			s, err := cron.Parse(task.CronInfo)
			if err != nil {
				log.Errorf("parse cron info fail, error:%+v, cron info:%s", err, task.CronInfo)
				return err
			}
			sd.Entries[i].Next = s.Next(time.Now())
			sd.Entries[i].Schedule = s
			task.CreateTime = t.Task.CreateTime
			sd.Entries[i].Task = task
		}
	}
	return fmt.Errorf("cannt find the task")
}

func doDelete(ctx context.Context, task data_schema.TaskInfo) error {
	index := -1
	for i, t := range sd.Entries {
		if t.Task.TaskID == task.TaskID {
			index = i
		}
	}
	if index == -1 {
		return fmt.Errorf("cannt find the task")
	}
	ScheLock()
	defer ScheUnLock()
	sd.Entries = append(sd.Entries[:index], sd.Entries[index+1:]...)
	return nil
}

func ScheLock() {
	sd.Lock.Lock()
}

func ScheUnLock() {
	sd.Lock.Unlock()
}
