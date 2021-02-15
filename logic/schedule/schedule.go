package schedule

import (
	"context"
	"fmt"
	"sort"
	"time"

	"go_schedule/dao/mongodb"
	"go_schedule/dao/zookeeper"
	"go_schedule/logic/consistent_hash"
	"go_schedule/logic/execute"
	"go_schedule/util/config"
	"go_schedule/util/consts"
	"go_schedule/util/data_schema"
	"go_schedule/util/log"
	"go_schedule/util/tool"

	"github.com/go-zookeeper/zk"
	"github.com/robfig/cron"
	"go.mongodb.org/mongo-driver/bson"
)

var sd data_schema.Schedule

// InitSchedule 初始化调度
func InitSchedule(ctx context.Context) error {
	sd.Create = make(chan data_schema.TaskInfo, config.Viper.GetInt("schedule.create_size"))
	sd.Update = make(chan data_schema.TaskInfo, config.Viper.GetInt("schedule.update_size"))
	sd.Delete = make(chan data_schema.TaskInfo, config.Viper.GetInt("schedule.delete_size"))
	sd.Init = make(chan int, 10)
	sd.Entries = make([]data_schema.Entry, 0, 4)

	// 监听集群节点变化
	_, wchan, err := zookeeper.ChildrenWatch("/go_schedule/schedule")
	if err != nil {
		log.ErrLogger.Printf("wathch cluster node fail, error:%+v", err)
		return err
	}
	go ClusterChange(ctx, wchan)
	return nil
}

func doInitEntries(ctx context.Context) error {
	lockPath := fmt.Sprintf("%s/%s", consts.ZKLockPath, tool.IP)
	if _, err := zookeeper.CreateTemplateNode(lockPath, nil); err != nil {
		log.ErrLogger.Printf("lock fail, error:%+v", err)
		InitEntry()
		return err
	}
	defer zookeeper.DeleteNode(lockPath)
	time.Sleep(500 * time.Millisecond)
	consistent_hash.InitIPMd5List()
	tasks, err := mongodb.SearchTask(ctx, bson.D{})
	if err != nil {
		log.ErrLogger.Printf("search task fail, %+v", err)
		InitEntry()
		return err
	}
	for i, t := range tasks {
		taskMd5, err := consistent_hash.HashCalculation(t.TaskID)
		if err != nil {
			log.ErrLogger.Printf("task id hash calculation fail task_id:%s, error:%+v", t.TaskID, err)
			InitEntry()
			return err
		}
		if tool.IP == consistent_hash.SelectIP(taskMd5) && !TaskExists(t.TaskID) {
			log.InfoLogger.Printf("init entries, create index:%d, task:%+v", i, t)
			CreateTask(t)
		}
		if tool.IP != consistent_hash.SelectIP(taskMd5) && TaskExists(t.TaskID) {
			log.InfoLogger.Printf("init entries, delete index:%d, task:%+v", i, t)
			DeleteTask(t)
		}
	}
	return nil
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

func InitEntry() {
	sd.Init <- 1
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
		case <-sd.Init:
			doInitEntries(ctx)
		}
	}
}

func executeTask(ctx context.Context) {
	now := time.Now().In(tool.TimeLocal)
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
				log.ErrLogger.Printf("insert schedule into mongo fail, schedule:%+v, task:%+v, error:%+v", s, entry.Task, err)
			} else {
				log.InfoLogger.Printf("scheduleID:%+v", id)
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
		log.ErrLogger.Printf("parse task cron info fail, taskID:%d, cronInfo:%s, error:%+v", task.TaskID, task.CronInfo, err)
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
				log.ErrLogger.Printf("parse cron info fail, error:%+v, cron info:%s", err, task.CronInfo)
				return err
			}
			sd.Entries[i].Next = s.Next(time.Now())
			sd.Entries[i].Schedule = s
			task.CreateTime = t.Task.CreateTime
			sd.Entries[i].Task = task
		}
	}
	log.ErrLogger.Printf("[doUpdate]cannt find the task:%+v", task)
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
		log.ErrLogger.Printf("[doDelete]cannt find the task:%+v", task)
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

// ClusterChange 集群变化处理
func ClusterChange(ctx context.Context, c <-chan zk.Event) {
	select {
	case <-c:
		_, wchan, err := zookeeper.ChildrenWatch("/go_schedule/schedule")
		if err != nil {
			log.ErrLogger.Printf("wathch cluster node fail, error:%+v", err)
		} else {
			go ClusterChange(ctx, wchan)
		}
		InitEntry()
	}
}
