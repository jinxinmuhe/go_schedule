package zookeeper

import (
	"testing"

	"go_schedule/dao/mongodb"
	"go_schedule/util/config"

	"github.com/go-zookeeper/zk"
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

func TestCreatTemplateNode(t *testing.T) {
	for i := 0; i < 3; i++ {
		result, err := CreateTemplateNode("/jason_test/test", []byte("apple"))
		if err == zk.ErrNodeExists {
			t.Logf("exists error:%+v", err)
		} else {
			t.Errorf("error:%+v", err)
		}
		t.Logf("resp:%s", result)
	}
}

func TestChildrenNodes(t *testing.T) {
	result, err := ChildrenNodes("/brokers")
	if err != nil {
		t.Fatalf("error:%+v", err)
	}
	t.Logf("children:%+v", result)
}

func TestChildrenWatch(t *testing.T) {
	list, wchann, err := ChildrenWatch("/go_schedule/schedule")
	if err != nil {
		t.Errorf("error:%+v", err)
	}
	t.Logf("list:%+v", list)
	select {
	case c := <-wchann:
		t.Logf("chann:%+v", c)
	}
}
