package wsManager

import (
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
)

/**
 * @ClassName wsManager
 * @Description TODO
 * @Author khr
 * @Date 2023/6/9 10:45
 * @Version 1.0
 */

type ConnectionManager struct {
	connections     map[*Connection]bool
	lock            sync.RWMutex
	userConnections map[string][]*Connection
	userLock        sync.RWMutex
	broadcast       chan []byte
}

// 创建 WebSocket 连接管理器
func NewWebSocketManager() *ConnectionManager {
	return &ConnectionManager{
		connections:     make(map[*Connection]bool),
		lock:            sync.RWMutex{},
		userConnections: make(map[string][]*Connection),
		userLock:        sync.RWMutex{},
		broadcast:       make(chan []byte),
	}
}

/*
 * @MethodName AddConnection
 * @Description 注册websocketmanager连接
 * @Author khr
 * @Date 2023/6/9 10:37
 */

func (cm *ConnectionManager) AddConnection(conn *Connection) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	cm.connections[conn] = true
	cm.userLock.Lock()
	defer cm.userLock.Unlock()
	cm.userConnections[conn.data.UserName] = append(cm.userConnections[conn.data.UserName], conn)
}

/*
 * @MethodName Unregister
 * @Description 删除websocketmanager连接
 * @Author khr
 * @Date 2023/6/9 10:38
 */

func (cm *ConnectionManager) Unregister(conn *Connection) {
	cm.lock.Lock()
	defer cm.lock.Unlock()
	delete(cm.connections, conn)
	cm.userLock.Lock()
	defer cm.userLock.Unlock()

	connections := cm.userConnections[conn.data.UserName]
	for i, conns := range connections {
		if conns == conn {
			cm.userConnections[conn.data.UserName] = append(connections[:i], connections[i+1:]...)
			break
		}
	}
	close(conn.sc)
}

/*
 * @MethodName GetConnections
 * @Description 获取所有连接
 * @Author khr
 * @Date 2023/6/9 10:55
 */

func (cm *ConnectionManager) GetConnections() []*Connection {
	cm.lock.RLock()
	defer cm.lock.RUnlock()
	connections := make([]*Connection, 0, len(cm.connections))
	for connection := range cm.connections {
		connections = append(connections, connection)
	}

	return connections
}

/*
 * @MethodName BroadcastToUser
 * @Description  广播消息给指定用户
 * @Author khr
 * @Date 2023/6/9 10:47
 */

func (cm *ConnectionManager) BroadcastToUser(user string, message []byte) {
	cm.userLock.RLock()
	defer cm.userLock.RUnlock()

	connections := cm.userConnections[user]
	for _, conn := range connections {
		err := conn.ws.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			// 处理写入错误
			HandleWsErr(err)
		}
	}
}

/*
 * @MethodName BroadcastToAll
 * @Description  广播消息给所有连接
 * @Author khr
 * @Date 2023/6/9 10:47
 */

func (cm *ConnectionManager) BroadcastToAll(message []byte) {
	cm.lock.RLock()
	defer cm.lock.RUnlock()
	for connections := range cm.connections {

		fmt.Println(connections.data.UserName)
		fmt.Println(connections.data.Ip)
		err := connections.ws.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			// 处理写入错误
			HandleWsErr(err)
		}
	}
}

/*
 * @MethodName
 * @Description 错误处理
 * @Author khr
 * @Date 2023/6/9 10:56
 */

func HandleWsErr(err error) {
	errors.Join(err)
}
