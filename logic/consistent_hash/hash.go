package consistent_hash

import (
	"crypto/md5"
	"encoding/binary"
	"go_schedule/dao/zookeeper"
	"go_schedule/util/log"
	"sort"
)

var IPMd5 map[uint32]string

// HashCalculation 计算节点或者task_id的哈希
func HashCalculation(key string) (uint32, error) {
	hash := md5.New()
	if _, err := hash.Write([]byte(key)); err != nil {
		log.Errorf("hash write fail, key:%s, error:%+v", key, err)
		return 0, err
	}
	result := hash.Sum(nil)
	return binary.BigEndian.Uint32(result), nil
}

// SelectIP 选择一个节点
func SelectIP(taskMD5 uint32) string {
	md5s := make([]uint32, 0, len(IPMd5))
	for k := range IPMd5 {
		md5s = append(md5s, k)
	}
	sort.Slice(md5s, func(i, j int) bool { return md5s[i] < md5s[j] })
	if len(md5s) >= 1 && taskMD5 >= md5s[len(md5s)-1] {
		return IPMd5[md5s[len(md5s)-1]]
	}
	for i := range md5s {
		if taskMD5 < md5s[i] {
			return IPMd5[md5s[i]]
		}
	}
	return ""
}

// InitIPMd5List 获取集群各节点的md5
func InitIPMd5List() error {
	IPMd5Temp := make(map[uint32]string)
	basicPath := "/go_schedule/schedule"
	list, err := zookeeper.ChildrenNodes(basicPath)
	if err != nil {
		log.Errorf("children nodes fail, error:%+v", err)
		return err
	}

	for _, ip := range list {
		if code, err := HashCalculation(ip); err == nil {
			IPMd5Temp[code] = ip
		}
	}
	IPMd5 = IPMd5Temp
	return nil
}
