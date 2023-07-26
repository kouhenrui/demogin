package wsManager

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

/**
 * @ClassName ws
 * @Description TODO
 * @Author khr
 * @Date 2023/6/8 15:08
 * @Version 1.0
 */

// 设置websocket
// CheckOrigin防止跨站点的请求伪造
var upGrader = &websocket.Upgrader{
	//设置读取写入字节大小
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		//可以添加验证信息
		return true
	},
}
var manager = NewWebSocketManager()

// 定义一个 Data 结构体，用于保存用户的信息
type Data struct {
	Ip       string `json:"ip"`
	UserName string `json:"username"`
	Type     string `json:"type"`
	Content  string `json:"content"`
}

// 定义一个 connection 结构体，用于保存每个连接的信息
type Connection struct {
	ws        *websocket.Conn // WebSocket 连接
	data      *Data           // 用户数据
	sc        chan []byte     // 发送消息的通道
	isPrivate bool
}

func Routers(e *gin.Engine) {
	wsGroup := e.Group("/api/ws")
	{
		wsGroup.GET("/con", WSConn)
		wsGroup.GET("/ping", getcon)
	}
}

func getcon(c *gin.Context) {
	connections := manager.GetConnections()
	for _, y := range connections {
		fmt.Println("所有的连接", y.data)
	}
	c.Set("res", "success")
	//return
}
func WSConn(c *gin.Context) {
	//升级get请求为webSocket协议
	ws, wserr := upGrader.Upgrade(c.Writer, c.Request, nil)
	if wserr != nil {
		fmt.Println("websocket连接错误")
		c.Error(errors.New("ws连接错误"))
		return
	}

	fmt.Println(c.ClientIP(), "请求的ip地址")
	name := c.Query("name")
	//name := c.Param("name")
	fmt.Println("name", name)
	conn := &Connection{ws: ws, data: &Data{
		Ip: c.ClientIP(),
	}, sc: make(chan []byte, 1024), isPrivate: false}
	// 将连接添加到连接管理器

	manager.AddConnection(conn)

	fmt.Println(conn)
	//fmt.Println(manager.GetConnections(), "获取所有的连接")
	conn.conMessages()
	go conn.handleConnect()
	fmt.Println("websocket连接成功")

}

/*
 * @MethodName handleConnect
 * @Description 并发处理消息发送接收
 * @Author khr
 * @Date 2023/6/9 15:52
 */

func (con *Connection) handleConnect() {

	//defer con.ws.Close()
	//defer manager.Unregister(con)
	//启动异步读取和发送消息的 goroutine
	go con.readMessages()
	go con.writeMessages()
	//go manager.StartBroadcasting()
}

/*
 * @MethodName readMessages
 * @Description 读取ws信息
 * @Author khr
 * @Date 2023/6/9 14:32
 */

func (con *Connection) readMessages() {
	for {
		_, message, err := con.ws.ReadMessage()
		if err != nil {
			log.Println("Failed to read message from WebSocket:", err)
			break
		}
		if con.isPrivate {
			fmt.Println("说到信息了,message:", string(message))
			//name := "ws"
			//global.Producer(name, string(message))
			con.sc <- message
		} else {
			//golbal.Producer(string(message), name)
			fmt.Println("全局信道信息")
			manager.BroadcastToAll(message)
			//.broadcast <- message
			//conn.bc <- message
		}

	}
}

/*
 * @MethodName conMessages
 * @Description 连接成功发送标识
 * @Author khr
 * @Date 2023/6/9 15:48
 */

func (con *Connection) conMessages() {
	var msg = []byte("ws连接成功")
	err := con.ws.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		con.HandleErrorMessage(err)
	}
}

/*
 * @MethodName writeMessages
 * @Description 消息写入mq
 * @Author khr
 * @Date 2023/6/9 14:35
 */

func (con *Connection) writeMessages() {
	for message := range con.sc {
		con.sendToConsumer(message)
		//conn.fc <- message
		//err := conn.ws.WriteMessage(websocket.TextMessage, message)
		//if err != nil {
		//	log.Println("Failed to write message to WebSocket:", err)
		//	conn.HandleErrorMessage(err)
		//}
	}
}

func (con *Connection) HandleErrorMessage(err error) {
	fmt.Println("error:", err)
	con.ws.Close()
	close(con.sc)
	errors.Join(err)
	//con.ws.WriteMessage(websocket.TextMessage, []byte("Connection file!n"))

}

func (con *Connection) sendToConsumer(out []byte) {
	err := con.ws.WriteMessage(websocket.TextMessage, out)
	if err != nil {
		con.HandleErrorMessage(err)
	}
}
