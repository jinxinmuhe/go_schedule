package zookeeper

import (
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"strconv"
	"time"

	"go_schedule/util/config"
	"go_schedule/util/log"
	"go_schedule/util/tool"

	"github.com/go-zookeeper/zk"
)

var zkConn *zk.Conn

func InitZookeeper() {
	conn, _, err := zk.Connect(config.Viper.GetStringSlice("zookeeper.hosts"), config.Viper.GetDuration("zookeeper.timeout")*time.Second)
	if err != nil {
		panic(err)
	}
	zkConn = conn
	if exist, err := ExistNode("/go_schedule"); err != nil {
		panic(err)
	} else if !exist {
		if _, err := CreateNode("/go_schedule", nil); err != nil {
			panic(err)
		}
	}

	if exist, err := ExistNode("/go_schedule/schedule"); err != nil {
		panic(err)
	} else if !exist {
		if _, err := CreateNode("/go_schedule/schedule", nil); err != nil {
			panic(err)
		}
	}

	hash := md5.New()
	if _, err := hash.Write([]byte(tool.IP)); err != nil {
		panic(err)
	}
	result := hash.Sum(nil)
	m := binary.BigEndian.Uint32(result)
	md5Str := strconv.FormatUint(uint64(m), 10)
	if _, err := CreateTemplateNode(fmt.Sprintf("/go_schedule/schedule/%s", tool.IP), []byte(md5Str)); err != nil {
		panic(err)
	}
}

// CreateTemplateNode 创建临时节点
func CreateTemplateNode(path string, data []byte) (result string, err error) {
	for i := 0; i < 3; i++ {
		if result, err = zkConn.Create(path, data, zk.FlagEphemeral, zk.WorldACL(zk.PermAll)); err == nil {
			break
		}
	}
	if err != nil {
		log.Errorf("create template node fail, path:%s, data:%s, error:%+v", path, string(data), err)
		return "", err
	}
	return result, nil
}

// CreateSequenceNode 创建顺序节点
func CreateSequenceNode(path string, data []byte) (string, error) {
	result, err := zkConn.Create(path, data, zk.FlagSequence, nil)
	if err != nil {
		log.Errorf("create sequence node fail, path:%s, data:%s, error:%+v", path, string(data), err)
		return "", err
	}
	return result, nil
}

// CreateSeqTempNode 创建临时顺序节点
func CreateSeqTempNode(path string, data []byte) (string, error) {
	result, err := zkConn.Create(path, data, zk.FlagSequence|zk.FlagEphemeral, nil)
	if err != nil {
		log.Errorf("create sequence template node fail, path:%s, data:%s, error:%+v", path, string(data), err)
		return "", err
	}
	return result, nil
}

// CreateNode 创建持久节点
func CreateNode(path string, data []byte) (string, error) {
	var result string
	var err error
	i := 0
	for i = 0; i < 3; i++ {
		if result, err = zkConn.Create(path, data, 0, zk.WorldACL(zk.PermAll)); err == nil {
			break
		}
	}
	if i >= 3 {
		log.Errorf("create node fail, path:%s, data:%s, error:%+v", path, string(data), err)
		return "", err
	}
	return result, nil
}

// SetData 写数据
func SetData(path string, data []byte, version int32) error {
	var stat *zk.Stat
	var err error
	for i := 0; i < 3; i++ {
		if stat, err = zkConn.Set(path, data, version); err == nil {
			break
		}
	}

	if err != nil {
		log.Errorf("set data to zk fail, error:%+v, path:%s, data:%s, version:%d", err, path, data, version)
		return err
	}

	log.Infof("set data status:%+v", *stat)
	return nil
}

// GetData 获取数据
func GetData(path string) ([]byte, error) {
	var data []byte
	var err error
	for i := 0; i < 3; i++ {
		if data, _, err = zkConn.Get(path); err == nil {
			break
		}
	}

	if err != nil {
		log.Errorf("get data from zk fail, error:%+v, path:%s", err, path)
		return nil, err
	}
	return data, nil
}

// DeleteNode 删除节点
func DeleteNode(path string) (err error) {
	for i := 0; i < 3; i++ {
		if err = zkConn.Delete(path, -1); err == nil {
			break
		}
	}
	if err != nil {
		log.Errorf("delete zk node fail, error:%+v", err)
	}
	return err
}

// ExistNode 判断节点是否存在
func ExistNode(path string) (exist bool, err error) {
	for i := 0; i < 3; i++ {
		if exist, _, err = zkConn.Exists(path); err == nil {
			break
		}
	}
	if err != nil {
		log.Errorf("request zk exist fail, path:%s, error:%+v", path, err)
	}
	return exist, err
}

// ChildrenNodes 获取子节点
func ChildrenNodes(path string) (list []string, err error) {
	for i := 0; i < 3; i++ {
		if list, _, err = zkConn.Children(path); err == nil {
			break
		}
	}
	if err != nil {
		log.Errorf("get children node fail, path:%s, error:%+v", path, err)
	}
	return list, err
}
