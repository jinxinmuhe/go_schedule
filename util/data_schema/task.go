package data_schema

import (
	"time"
)

type TaskInfo struct {
	TaskID     string    `json:"task_id" bson:"task_id,omitempty"`
	TaskName   string    `json:"task_name" bson:"task_name"`
	Owner      string    `json:"owner" bson:"owner"`
	CronInfo   string    `json:"cron_info" bson:"cron_info"`
	KafkaTopic string    `json:"kafka_topic" bson:"kafka_topic"`
	AlarmMail  string    `json:"alarm_mail" bson:"alarm_mail"`
	Variables  string    `json:"variables" bson:"variables"`
	CreateTime time.Time `json:"create_time" bson:"create_time"`
	UpdateTime time.Time `json:"update_time" bson:"update_time"`
}

type ScheduleHistory struct {
	TaskID     string    `json:"task_id" bson:"task_id"`
	ScheduleID int64     `json:"schedule_id" bson:"schedule_id"`
	NextTime   time.Time `json:"next_time" bson:"next_time"`
	Variables  string    `json:"variables" bson:"variables"`
	CreateTime time.Time `json:"create_time" bson:"create_time"`
	UpdateTime time.Time `json:"update_time" bson:"update_time"`
}

type ExecuteLog struct {
	TaskID     string    `json:"task_id" bson:"task_id"`
	ScheduleID int64     `json:"schedule_id" bson:"schedule_id"`
	ResultCode int       `json:"result_code" bson:"result_code"`
	DetailMsg  string    `json:"detail_msg" bson:"detail_msg"`
	CreateTime time.Time `json:"create_time" bson:"create_time"`
	UpdateTime time.Time `json:"update_time" bson:"update_time"`
}
