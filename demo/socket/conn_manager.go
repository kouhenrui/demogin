package main

import (
	"sync"
)

/**
 * @ClassName conn_manager
 * @Description TODO
 * @Author khr
 * @Date 2023/5/29 14:29
 * @Version 1.0
 */
var ConnsManager = sync.Map{}

// SetConn 存储
func SetConn(deviceId int64, conn *Conn) {
	ConnsManager.Store(deviceId, conn)
}

// GetConn 获取
func GetConn(deviceId int64) *Conn {
	value, ok := ConnsManager.Load(deviceId)
	if ok {
		return value.(*Conn)
	}
	return nil
}

// DeleteConn 删除
func DeleteConn(deviceId int64) {
	ConnsManager.Delete(deviceId)
}

// PushAll 全服推送
//func PushAll(message *pb.Message) {
//	ConnsManager.Range(func(key, value interface{}) bool {
//		conn := value.(*Conn)
//		conn.Send(pb.PackageType_PT_MESSAGE, 0, message, nil)
//		return true
//	})
//}
