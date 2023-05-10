package ws

import (
	"HelloGin/src/global"
	"HelloGin/src/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// 定义一个 Data 结构体，用于保存用户的信息
type Data struct {
	Ip       string   `json:"ip"`
	User     string   `json:"user"`
	UserName string   `json:username`
	From     string   `json:"from"`
	Type     string   `json:"type"`
	Content  string   `json:"content"`
	UserList []string `json:"user_list"`
}

// 定义一个 connection 结构体，用于保存每个连接的信息
type connection struct {
	ws   *websocket.Conn // WebSocket 连接
	data *Data           // 用户数据
	sc   chan []byte     // 发送消息的通道
}

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

// 定义一个全局的 WebSocket 连接池
var connections = make(map[*connection]bool)

// 全局消息通道，用于接收所有连接的消息
var broadcast = make(chan []byte)

// 连接管理器
type Manager struct {
	connections map[*connection]bool
	register    chan *connection
	unregister  chan *connection
}

var user_list = []string{}

func Routers(e *gin.Engine) {
	wsGroup := e.Group("/api/ws")
	{
		wsGroup.GET("/connect", wsconnect)

	}
}

// websocket实现
func wsconnect(c *gin.Context) {
	// 初始化连接管理器
	//manager := Manager{
	//	connections: make(map[*connection]bool),
	//	register:    make(chan *connection),
	//	unregister:  make(chan *connection),
	//}

	//fmt.Print("进入ws连接")
	result := global.NewResult(c)
	//升级get请求为webSocket协议
	ws, wserr := upGrader.Upgrade(c.Writer, c.Request, nil)
	if wserr != nil {
		fmt.Println("websocket连接错误")
		result.Err(util.WEBSOCKET_CONNECT_ERROR)
		return
	}
	//fmt.Println("将get请求转换为ws请求")
	conn := &connection{ws: ws, data: &Data{}, sc: make(chan []byte, 1024)}
	connections[conn] = true
	fmt.Println("进行注册")
	// 将连接对象注册到连接管理器中
	//manager.register <- conn
	//fmt.Println(<-manager.register, "打印注册的")
	// 异步处理 Websocket 消息
	go conn.handleConnection()

}

// 从连接池中删除连接
func (conn *connection) close() {
	delete(connections, conn)
	close(conn.sc)
}
func (conn *connection) handleConnection() {
	defer conn.close()
	fmt.Println("开始异步手法信息")
	// 创建异步读取和发送消息的 channel
	readCh := make(chan []byte)
	writeCh := make(chan []byte)
	// 启动异步读取和发送消息的 goroutine
	go conn.readMessages(readCh)
	go conn.writeMessages(writeCh, writeCh)

	// 处理 Websocket 消息
	for {
		select {
		case message := <-readCh:
			if string(message) == "out" {
				conn.close()
			}
			//if string(message) == "login" {
			//	_ = conn.WriteMessage(websocket.TextMessage, []byte("欢迎加入ws连接"))
			//}
			fmt.Println(string(message))
			// 处理来自客户端的消息
			conn.HandleMessage(message)

		case message := <-writeCh:
			// 发送消息给客户端
			if err := conn.ws.WriteMessage(websocket.TextMessage, message); err != nil {
				// 发送消息失败
				// 处理错误
				conn.HandleErrorMessage(err)
			}
		}
	}
}

/*
 * @MethodName readMessages
 * @Description 协程读取信道
 * @Author khr
 * @Date 2023/4/12 11:08
 */

func (conn *connection) readMessages(ch chan<- []byte) {
	//fmt.Println("读取ws里面的信息")
	for {
		// 读取客户端发送的消息
		_, message, err := conn.ws.ReadMessage()
		if err != nil {
			// 发生错误
			log.Println(err)
			return
		}
		// 将消息发送给处理函数
		ch <- message
	}
}

/*
 * @MethodName writeMessages
 * @Description 协程写入信道
 * @Author khr
 * @Date 2023/4/12 11:09
 */

func (conn *connection) writeMessages(in <-chan []byte, out chan<- []byte) {
	message := <-in

	//	// 将消息发送给客户端
	out <- message
	//}
}

/*
 * @MethodName HandleMessage
 * @Description 处理发送过来的信息，并返回
 * @Author khr
 * @Date 2023/4/12 11:05
 */

func (conn *connection) HandleMessage(in []byte) {
	//fmt.Println(in)
	//global.Producer(message, name)
	conn.ws.WriteMessage(websocket.TextMessage, in)
}

/*
 * @MethodName HandleErrorMessage
 * @Description 发送消息失败,处理错误
 * @Author khr
 * @Date 2023/4/12 11:08
 */

func (conn *connection) HandleErrorMessage(err error) {
	fmt.Println("error:", err)
	conn.ws.WriteMessage(websocket.TextMessage, []byte("Connection file!n"))
	conn.close()
}

// 广播信息
func (manager *Manager) broadcastMessage(message []byte) {
	for conn := range manager.connections {
		select {
		case conn.sc <- message: // 发送消息到连接的通道中
		default:
			// 如果连接的通道已满，则从连接管理器中删除该连接
			delete(manager.connections, conn)
			close(conn.sc)
		}
	}
}
