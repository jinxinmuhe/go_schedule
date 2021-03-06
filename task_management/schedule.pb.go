// Code generated by protoc-gen-go. DO NOT EDIT.
// source: schedule.proto

/*
Package task_management is a generated protocol buffer package.

It is generated from these files:
	schedule.proto

It has these top-level messages:
	CreateTaskReq
	CreateTaskResp
	UpdateTaskReq
	UpdateTaskResp
	DeleteTaskReq
	DeleteTaskResp
	SearchTaskReq
	TaskInfo
	SearchTaskResp
*/
package task_management

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RespCode int32

const (
	RespCode_UNKNOWN RespCode = 0
	RespCode_SUCC    RespCode = 1
	RespCode_FAIL    RespCode = 2
)

var RespCode_name = map[int32]string{
	0: "UNKNOWN",
	1: "SUCC",
	2: "FAIL",
}
var RespCode_value = map[string]int32{
	"UNKNOWN": 0,
	"SUCC":    1,
	"FAIL":    2,
}

func (x RespCode) String() string {
	return proto.EnumName(RespCode_name, int32(x))
}
func (RespCode) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CreateTaskReq struct {
	TaskName     string `protobuf:"bytes,1,opt,name=task_name,json=taskName" json:"task_name,omitempty"`
	ScheduleTime string `protobuf:"bytes,2,opt,name=schedule_time,json=scheduleTime" json:"schedule_time,omitempty"`
	KafkaTopic   string `protobuf:"bytes,3,opt,name=kafka_topic,json=kafkaTopic" json:"kafka_topic,omitempty"`
	AlarmEmail   string `protobuf:"bytes,4,opt,name=alarm_email,json=alarmEmail" json:"alarm_email,omitempty"`
	Owner        string `protobuf:"bytes,5,opt,name=owner" json:"owner,omitempty"`
	TaskId       string `protobuf:"bytes,6,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	EnvVariable  string `protobuf:"bytes,20,opt,name=env_variable,json=envVariable" json:"env_variable,omitempty"`
}

func (m *CreateTaskReq) Reset()                    { *m = CreateTaskReq{} }
func (m *CreateTaskReq) String() string            { return proto.CompactTextString(m) }
func (*CreateTaskReq) ProtoMessage()               {}
func (*CreateTaskReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateTaskReq) GetTaskName() string {
	if m != nil {
		return m.TaskName
	}
	return ""
}

func (m *CreateTaskReq) GetScheduleTime() string {
	if m != nil {
		return m.ScheduleTime
	}
	return ""
}

func (m *CreateTaskReq) GetKafkaTopic() string {
	if m != nil {
		return m.KafkaTopic
	}
	return ""
}

func (m *CreateTaskReq) GetAlarmEmail() string {
	if m != nil {
		return m.AlarmEmail
	}
	return ""
}

func (m *CreateTaskReq) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *CreateTaskReq) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *CreateTaskReq) GetEnvVariable() string {
	if m != nil {
		return m.EnvVariable
	}
	return ""
}

type CreateTaskResp struct {
	TaskId string   `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	Code   RespCode `protobuf:"varint,2,opt,name=code,enum=task_management.RespCode" json:"code,omitempty"`
	Msg    string   `protobuf:"bytes,3,opt,name=msg" json:"msg,omitempty"`
}

func (m *CreateTaskResp) Reset()                    { *m = CreateTaskResp{} }
func (m *CreateTaskResp) String() string            { return proto.CompactTextString(m) }
func (*CreateTaskResp) ProtoMessage()               {}
func (*CreateTaskResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateTaskResp) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *CreateTaskResp) GetCode() RespCode {
	if m != nil {
		return m.Code
	}
	return RespCode_UNKNOWN
}

func (m *CreateTaskResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type UpdateTaskReq struct {
	TaskId       string `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	ScheduleTime string `protobuf:"bytes,3,opt,name=schedule_time,json=scheduleTime" json:"schedule_time,omitempty"`
	KafkaTopic   string `protobuf:"bytes,4,opt,name=kafka_topic,json=kafkaTopic" json:"kafka_topic,omitempty"`
	AlarmEmail   string `protobuf:"bytes,5,opt,name=alarm_email,json=alarmEmail" json:"alarm_email,omitempty"`
	Owner        string `protobuf:"bytes,6,opt,name=owner" json:"owner,omitempty"`
	EnvVariable  string `protobuf:"bytes,20,opt,name=env_variable,json=envVariable" json:"env_variable,omitempty"`
}

func (m *UpdateTaskReq) Reset()                    { *m = UpdateTaskReq{} }
func (m *UpdateTaskReq) String() string            { return proto.CompactTextString(m) }
func (*UpdateTaskReq) ProtoMessage()               {}
func (*UpdateTaskReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UpdateTaskReq) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *UpdateTaskReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateTaskReq) GetScheduleTime() string {
	if m != nil {
		return m.ScheduleTime
	}
	return ""
}

func (m *UpdateTaskReq) GetKafkaTopic() string {
	if m != nil {
		return m.KafkaTopic
	}
	return ""
}

func (m *UpdateTaskReq) GetAlarmEmail() string {
	if m != nil {
		return m.AlarmEmail
	}
	return ""
}

func (m *UpdateTaskReq) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *UpdateTaskReq) GetEnvVariable() string {
	if m != nil {
		return m.EnvVariable
	}
	return ""
}

type UpdateTaskResp struct {
	Code RespCode `protobuf:"varint,1,opt,name=code,enum=task_management.RespCode" json:"code,omitempty"`
	Msg  string   `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *UpdateTaskResp) Reset()                    { *m = UpdateTaskResp{} }
func (m *UpdateTaskResp) String() string            { return proto.CompactTextString(m) }
func (*UpdateTaskResp) ProtoMessage()               {}
func (*UpdateTaskResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UpdateTaskResp) GetCode() RespCode {
	if m != nil {
		return m.Code
	}
	return RespCode_UNKNOWN
}

func (m *UpdateTaskResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type DeleteTaskReq struct {
	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
}

func (m *DeleteTaskReq) Reset()                    { *m = DeleteTaskReq{} }
func (m *DeleteTaskReq) String() string            { return proto.CompactTextString(m) }
func (*DeleteTaskReq) ProtoMessage()               {}
func (*DeleteTaskReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeleteTaskReq) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

type DeleteTaskResp struct {
	TaskId string   `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	Code   RespCode `protobuf:"varint,2,opt,name=code,enum=task_management.RespCode" json:"code,omitempty"`
	Msg    string   `protobuf:"bytes,3,opt,name=msg" json:"msg,omitempty"`
}

func (m *DeleteTaskResp) Reset()                    { *m = DeleteTaskResp{} }
func (m *DeleteTaskResp) String() string            { return proto.CompactTextString(m) }
func (*DeleteTaskResp) ProtoMessage()               {}
func (*DeleteTaskResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *DeleteTaskResp) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *DeleteTaskResp) GetCode() RespCode {
	if m != nil {
		return m.Code
	}
	return RespCode_UNKNOWN
}

func (m *DeleteTaskResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type SearchTaskReq struct {
	TaskId     string `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	KafkaTopic string `protobuf:"bytes,2,opt,name=kafka_topic,json=kafkaTopic" json:"kafka_topic,omitempty"`
	Owner      string `protobuf:"bytes,3,opt,name=owner" json:"owner,omitempty"`
}

func (m *SearchTaskReq) Reset()                    { *m = SearchTaskReq{} }
func (m *SearchTaskReq) String() string            { return proto.CompactTextString(m) }
func (*SearchTaskReq) ProtoMessage()               {}
func (*SearchTaskReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SearchTaskReq) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *SearchTaskReq) GetKafkaTopic() string {
	if m != nil {
		return m.KafkaTopic
	}
	return ""
}

func (m *SearchTaskReq) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type TaskInfo struct {
	TaskId       string `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
	TaskName     string `protobuf:"bytes,2,opt,name=task_name,json=taskName" json:"task_name,omitempty"`
	ScheduleTime string `protobuf:"bytes,3,opt,name=schedule_time,json=scheduleTime" json:"schedule_time,omitempty"`
	KafkaTopic   string `protobuf:"bytes,4,opt,name=kafka_topic,json=kafkaTopic" json:"kafka_topic,omitempty"`
	AlarmEmail   string `protobuf:"bytes,5,opt,name=alarm_email,json=alarmEmail" json:"alarm_email,omitempty"`
	Owner        string `protobuf:"bytes,6,opt,name=owner" json:"owner,omitempty"`
	EnvVariable  string `protobuf:"bytes,20,opt,name=env_variable,json=envVariable" json:"env_variable,omitempty"`
}

func (m *TaskInfo) Reset()                    { *m = TaskInfo{} }
func (m *TaskInfo) String() string            { return proto.CompactTextString(m) }
func (*TaskInfo) ProtoMessage()               {}
func (*TaskInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TaskInfo) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *TaskInfo) GetTaskName() string {
	if m != nil {
		return m.TaskName
	}
	return ""
}

func (m *TaskInfo) GetScheduleTime() string {
	if m != nil {
		return m.ScheduleTime
	}
	return ""
}

func (m *TaskInfo) GetKafkaTopic() string {
	if m != nil {
		return m.KafkaTopic
	}
	return ""
}

func (m *TaskInfo) GetAlarmEmail() string {
	if m != nil {
		return m.AlarmEmail
	}
	return ""
}

func (m *TaskInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *TaskInfo) GetEnvVariable() string {
	if m != nil {
		return m.EnvVariable
	}
	return ""
}

type SearchTaskResp struct {
	Tasks []*TaskInfo `protobuf:"bytes,1,rep,name=tasks" json:"tasks,omitempty"`
	Code  RespCode    `protobuf:"varint,2,opt,name=code,enum=task_management.RespCode" json:"code,omitempty"`
	Msg   string      `protobuf:"bytes,3,opt,name=msg" json:"msg,omitempty"`
}

func (m *SearchTaskResp) Reset()                    { *m = SearchTaskResp{} }
func (m *SearchTaskResp) String() string            { return proto.CompactTextString(m) }
func (*SearchTaskResp) ProtoMessage()               {}
func (*SearchTaskResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *SearchTaskResp) GetTasks() []*TaskInfo {
	if m != nil {
		return m.Tasks
	}
	return nil
}

func (m *SearchTaskResp) GetCode() RespCode {
	if m != nil {
		return m.Code
	}
	return RespCode_UNKNOWN
}

func (m *SearchTaskResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateTaskReq)(nil), "task_management.CreateTaskReq")
	proto.RegisterType((*CreateTaskResp)(nil), "task_management.CreateTaskResp")
	proto.RegisterType((*UpdateTaskReq)(nil), "task_management.UpdateTaskReq")
	proto.RegisterType((*UpdateTaskResp)(nil), "task_management.UpdateTaskResp")
	proto.RegisterType((*DeleteTaskReq)(nil), "task_management.DeleteTaskReq")
	proto.RegisterType((*DeleteTaskResp)(nil), "task_management.DeleteTaskResp")
	proto.RegisterType((*SearchTaskReq)(nil), "task_management.SearchTaskReq")
	proto.RegisterType((*TaskInfo)(nil), "task_management.TaskInfo")
	proto.RegisterType((*SearchTaskResp)(nil), "task_management.SearchTaskResp")
	proto.RegisterEnum("task_management.RespCode", RespCode_name, RespCode_value)
}

func init() { proto.RegisterFile("schedule.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x95, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xb3, 0xb6, 0x93, 0xa6, 0x93, 0xda, 0x44, 0xab, 0x4a, 0x18, 0x2a, 0x68, 0x31, 0x97,
	0x08, 0x44, 0x2a, 0x95, 0x27, 0x80, 0x00, 0x52, 0x04, 0xa4, 0x22, 0x4d, 0x40, 0xe2, 0x62, 0x6d,
	0xe3, 0x69, 0x6a, 0xe2, 0x7f, 0xd8, 0x26, 0xdc, 0x79, 0x46, 0x8e, 0x1c, 0xb8, 0xf0, 0x1e, 0x68,
	0xd7, 0x71, 0xec, 0x4d, 0x1c, 0x0c, 0x42, 0x48, 0xdc, 0xd6, 0xb3, 0x5f, 0x3e, 0xef, 0xfc, 0x3e,
	0xcf, 0x06, 0x8c, 0x64, 0x76, 0x8d, 0xce, 0x27, 0x0f, 0xfb, 0x51, 0x1c, 0xa6, 0x21, 0xbd, 0x91,
	0xb2, 0x64, 0x61, 0xfb, 0x2c, 0x60, 0x73, 0xf4, 0x31, 0x48, 0xad, 0x1f, 0x04, 0xf4, 0x41, 0x8c,
	0x2c, 0xc5, 0x09, 0x4b, 0x16, 0x63, 0xfc, 0x48, 0x8f, 0x60, 0x5f, 0x88, 0x02, 0xe6, 0xa3, 0x49,
	0x4e, 0x48, 0x6f, 0x7f, 0xdc, 0xe6, 0x85, 0x11, 0xf3, 0x91, 0xde, 0x07, 0x3d, 0x77, 0xb4, 0x53,
	0xd7, 0x47, 0x53, 0x11, 0x82, 0x83, 0xbc, 0x38, 0x71, 0x7d, 0xa4, 0xc7, 0xd0, 0x59, 0xb0, 0xab,
	0x05, 0xb3, 0xd3, 0x30, 0x72, 0x67, 0xa6, 0x2a, 0x24, 0x20, 0x4a, 0x13, 0x5e, 0xe1, 0x02, 0xe6,
	0xb1, 0xd8, 0xb7, 0xd1, 0x67, 0xae, 0x67, 0x6a, 0x99, 0x40, 0x94, 0x9e, 0xf3, 0x0a, 0x3d, 0x84,
	0x66, 0xf8, 0x39, 0xc0, 0xd8, 0x6c, 0x8a, 0xad, 0xec, 0x81, 0xde, 0x84, 0x3d, 0x71, 0x32, 0xd7,
	0x31, 0x5b, 0xa2, 0xde, 0xe2, 0x8f, 0x43, 0x87, 0xde, 0x83, 0x03, 0x0c, 0x96, 0xf6, 0x92, 0xc5,
	0x2e, 0xbb, 0xf4, 0xd0, 0x3c, 0x14, 0xbb, 0x1d, 0x0c, 0x96, 0x6f, 0x57, 0x25, 0xeb, 0x03, 0x18,
	0xe5, 0x36, 0x93, 0xa8, 0xec, 0x46, 0x24, 0xb7, 0x47, 0xa0, 0xcd, 0x42, 0x27, 0x6b, 0xcd, 0x38,
	0xbb, 0xd5, 0xdf, 0x40, 0xd6, 0xe7, 0xbf, 0x1e, 0x84, 0x0e, 0x8e, 0x85, 0x8c, 0x76, 0x41, 0xf5,
	0x93, 0xf9, 0xaa, 0x4b, 0xbe, 0xb4, 0xbe, 0x11, 0xd0, 0xa7, 0x91, 0x53, 0x62, 0xba, 0xf3, 0x5d,
	0x14, 0x34, 0xc1, 0x39, 0xc3, 0x28, 0xd6, 0xdb, 0x8c, 0xd5, 0x7a, 0xc6, 0x5a, 0x1d, 0xe3, 0xe6,
	0x6e, 0xc6, 0xad, 0x32, 0xe3, 0xdf, 0x40, 0xf9, 0x06, 0x8c, 0x72, 0x77, 0x49, 0xb4, 0x26, 0x46,
	0xfe, 0x88, 0x98, 0x52, 0x10, 0xeb, 0x81, 0xfe, 0x0c, 0x3d, 0xac, 0x07, 0xc6, 0x73, 0x2c, 0x2b,
	0xff, 0x69, 0x8e, 0x36, 0xe8, 0x17, 0xc8, 0xe2, 0xd9, 0x75, 0x6d, 0x8c, 0x1b, 0x69, 0x28, 0x5b,
	0x69, 0xac, 0x61, 0xab, 0x25, 0xd8, 0xd6, 0x77, 0x02, 0x6d, 0xee, 0x3d, 0x0c, 0xae, 0xc2, 0xdd,
	0xe6, 0xd2, 0x40, 0x2a, 0x75, 0x03, 0xf9, 0x9f, 0x7d, 0x2c, 0x5f, 0x08, 0x18, 0x65, 0x88, 0x49,
	0x44, 0x4f, 0xa1, 0xc9, 0x8f, 0x9f, 0x98, 0xe4, 0x44, 0xed, 0x75, 0x2a, 0x82, 0xc9, 0x91, 0x8c,
	0x33, 0xdd, 0x5f, 0x07, 0xf9, 0xe0, 0x21, 0xb4, 0x73, 0x0d, 0xed, 0xc0, 0xde, 0x74, 0xf4, 0x72,
	0x74, 0xfe, 0x6e, 0xd4, 0x6d, 0xd0, 0x36, 0x68, 0x17, 0xd3, 0xc1, 0xa0, 0x4b, 0xf8, 0xea, 0xc5,
	0x93, 0xe1, 0xab, 0xae, 0x72, 0xf6, 0x55, 0x01, 0x83, 0x9f, 0xe0, 0xf5, 0xfa, 0x05, 0xf4, 0x1c,
	0xa0, 0xb8, 0x3c, 0xe8, 0xdd, 0xad, 0x03, 0x48, 0x17, 0xe8, 0xed, 0xe3, 0x5f, 0xee, 0x27, 0x91,
	0xd5, 0xe0, 0x86, 0xc5, 0x08, 0x55, 0x18, 0x4a, 0xb7, 0x47, 0x85, 0xa1, 0x3c, 0x7f, 0x99, 0x61,
	0x31, 0x16, 0x15, 0x86, 0xd2, 0x74, 0x55, 0x18, 0xca, 0x33, 0x95, 0x19, 0x16, 0xb1, 0x55, 0x18,
	0x4a, 0x83, 0x51, 0x61, 0x28, 0x67, 0x6e, 0x35, 0x9e, 0xde, 0x79, 0x7f, 0x34, 0x0f, 0xed, 0xfc,
	0xb3, 0x3c, 0xdd, 0xd0, 0x5f, 0xb6, 0xc4, 0xff, 0xd3, 0xe3, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x46, 0x3c, 0xa9, 0xf0, 0xb1, 0x06, 0x00, 0x00,
}
