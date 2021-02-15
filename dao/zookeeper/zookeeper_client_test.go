package zookeeper

import (
	"testing"

	"go_schedule/dao/mongodb"
	"go_schedule/util/config"
)

func init() {
	config.InitConfig("")
	mongodb.InitMongodb()
	InitZookeeper()
}

func TestSetData(t *testing.T) {
	err := SetData("/jasonjinxin0000000000", []byte("jasonjinxin"), -1)
	if err != nil {
		t.Fatalf("set data fail:%+v", err)
	}
}

func TestCreateNode(t *testing.T) {
	result, err := CreateNode("/jason_test", []byte("apple"))
	if err != nil {
		t.Fatalf("error:%+v", err)
	}
	t.Logf("resp:%s", result)
}

func TestChildrenNodes(t *testing.T) {
	result, err := ChildrenNodes("/brokers")
	if err != nil {
		t.Fatalf("error:%+v", err)
	}
	t.Logf("children:%+v", result)
}
