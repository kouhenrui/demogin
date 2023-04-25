package ws

import (
	"HelloGin/src/global"
	"HelloGin/src/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

//type hub struct {
//	c map[*connection]bool
//	b chan []byte
//	r chan *connection
//	u chan *connection
//}
//
//var h = hub{
//	c: make(map[*connection]bool),
//	u: make(chan *connection),
//	b: make(chan []byte),
//	r: make(chan *connection),
//}

type connection struct {
	ws   *websocket.Conn
	sc   chan []byte
	data *Data
}
type Data struct {
	Ip       string   `json:"ip"`
	User     string   `json:"user"`
	From     string   `json:"from"`
	Type     string   `json:"type"`
	Content  string   `json:"content"`
	UserList []string `json:"user_list"`
}

var user_list = []string{}
var name = "test"

func Routers(e *gin.Engine) {
	wsGroup := e.Group("/api/ws")
	{
		wsGroup.GET("/connect", wsconnect)
	}
}

// 设置websocket
// CheckOrigin防止跨站点的请求伪造
var upGrader = &websocket.Upgrader{
	//设置读取写入字节大小
	ReadBufferSize:  512,
	WriteBufferSize: 512,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		//可以添加验证信息
		return true
	},
}

// websocket实现
func wsconnect(c *gin.Context) {
	fmt.Print("进入ws连接")
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
	fmt.Println("开始决定进入那个线程", conn)
	// 异步处理 Websocket 消息
	go handleConnection(ws)
	//defer ws.Close()
	//go con.Writer()
	//go con.Reader()
	//defer func() {
	//	count := len(user_list)
	//	if count == 0 {
	//		ws.Close()
	//	}
	//
	//}()
	//defer ws.Close() //返回前关闭
	//for {
	//	//读取ws中的数据
	//
	//	fmt.Println("收取信息2")
	//	mt, message, err := ws.ReadMessage()
	//	fmt.Println(mt, string(message), "打印获取的信息")
	//	if err != nil {
	//		result.Err("连接错误")
	//		return
	//	}
	//	//fmt.Print(string(message), "message")
	//	//写入ws数据
	//	err = ws.WriteMessage(mt, []byte("welcome to ws"))
	//	if err != nil {
	//		result.Err("连接错误3")
	//		return
	//	}
	//}
	//fmt.Println("我们结束聊天啦")
}
func handleConnection(conn *websocket.Conn) {
	// 创建异步读取和发送消息的 channel
	readCh := make(chan []byte)
	writeCh := make(chan []byte)

	// 启动异步读取和发送消息的 goroutine
	go readMessages(conn, readCh)
	go writeMessages(conn, writeCh, writeCh)

	// 处理 Websocket 消息
	for {
		select {
		case message := <-readCh:
			if string(message) == "out" {
				conn.Close()
			}
			// 处理来自客户端的消息
			HandleMessage(string(message), conn)

		case message := <-writeCh:
			// 发送消息给客户端
			if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
				// 发送消息失败
				// 处理错误
				HandleErrorMessage(err, conn)
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

func readMessages(conn *websocket.Conn, ch chan<- []byte) {
	for {
		// 读取客户端发送的消息
		_, message, err := conn.ReadMessage()
		if err != nil {
			// 发生错误
			close(ch)
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

func writeMessages(conn *websocket.Conn, in <-chan []byte, out chan<- []byte) {
	//message := global.Consumer(name)
	//fmt.Println(message, "打印从rabbitmq队取得信息")
	//for message := range in {
	//	// 将消息发送给客户端
	out <- []byte("message")
	//}
}

/*
 * @MethodName HandleMessage
 * @Description 处理发送过来的信息，并返回
 * @Author khr
 * @Date 2023/4/12 11:05
 */

func HandleMessage(message string, conn *websocket.Conn) {
	fmt.Println(message)
	global.Producer(message, name)
	conn.WriteMessage(websocket.TextMessage, []byte("websocket已发送信息到rabbitmq"))
}

/*
 * @MethodName HandleErrorMessage
 * @Description 发送消息失败,处理错误
 * @Author khr
 * @Date 2023/4/12 11:08
 */

func HandleErrorMessage(err error, conn *websocket.Conn) {
	fmt.Println("error:", err)
	conn.WriteMessage(websocket.TextMessage, []byte("Connection file!n"))
	conn.Close()
}

//func (c *connection) Writer() {
//	fmt.Println("进入编写信息携程")
//	for message := range c.sc {
//		c.ws.WriteMessage(websocket.TextMessage, message)
//	}
//	c.ws.Close()
//}
//
//func (c *connection) Reader() {
//	fmt.Println("进入读取信息携程")
//	for {
//		_, message, err := c.ws.ReadMessage()
//		if err != nil {
//			h.r <- c
//			break
//		}
//		fmt.Println(string(message), "读取inxi")
//		fmt.Println(c.data.Type, "读取inxi")
//		json.Unmarshal(message, &c.data)
//		switch c.data.Type {
//		case "login":
//			c.register()
//		case "user":
//			c.data.Type = "user"
//			data_b, _ := json.Marshal(c.data)
//			fmt.Println(string(data_b), "读取的信息")
//			h.b <- data_b
//			fmt.Println(string(<-h.b))
//		case "logout":
//			fmt.Println("用户推出")
//			c.del()
//		default:
//			fmt.Print("========default================")
//		}
//	}
//}
//
//func (c *connection) del() {
//
//	count := len(user_list)
//	var n_slice = []string{}
//	for i := range user_list {
//		if user_list[i] == c.data.User && i == count {
//			n_slice = user_list[:count]
//		} else if user_list[i] == c.data.User {
//			n_slice = append(user_list[:i], user_list[i+1:]...)
//			break
//		}
//	}
//
//	if len(n_slice) <= 0 {
//		c.ws.Close()
//	}
//	data_b, _ := json.Marshal(c.data)
//	h.b <- data_b
//	h.r <- c
//}
//func (c *connection) register() {
//	c.data.User = c.data.Content
//	c.data.From = c.data.User
//	user_list = append(user_list, c.data.User)
//	c.data.UserList = user_list
//	data_b, _ := json.Marshal(c.data)
//	h.b <- data_b
//}

//type Connection struct {
//	wsConn *websocket.Conn
//	//读取websocket的channel
//	inChan chan []byte
//	//给websocket写消息的channel
//	outChan   chan []byte
//	closeChan chan byte
//	mutex     sync.Mutex
//	//closeChan 状态
//	isClosed bool
//}
//
////初始化长连接
//func InitConnection(wsConn *websocket.Conn) (conn *Connection, err error) {
//	conn = &Connection{
//		wsConn:    wsConn,
//		inChan:    make(chan []byte, 1000),
//		outChan:   make(chan []byte, 1000),
//		closeChan: make(chan byte, 1),
//	}
//	//启动读协程
//	go conn.readLoop()
//	//启动写协程
//	go conn.writeLoop()
//	return
//}
//
////读取websocket消息
//func (conn *Connection) ReadMessage() (data []byte, err error) {
//	select {
//	case data = <-conn.inChan:
//	case <-conn.closeChan:
//		err = errors.New("connection is closed")
//	}
//	return
//}
//
////发送消息到websocket
//func (conn *Connection) WriteMessage(data []byte) (err error) {
//	select {
//	case conn.outChan <- data:
//	case <-conn.closeChan:
//		err = errors.New("connection is closed")
//	}
//	return
//}
//
////关闭连接
//func (conn *Connection) Close() {
//	//线程安全的Close,可重入
//	conn.wsConn.Close()
//
//	//只执行一次
//	conn.mutex.Lock()
//	if !conn.isClosed {
//		close(conn.closeChan)
//		conn.isClosed = true
//	}
//	conn.mutex.Unlock()
//}
//
//func (conn *Connection) readLoop() {
//	var (
//		data []byte
//		err  error
//	)
//	for {
//		if _, data, err = conn.wsConn.ReadMessage(); err != nil {
//			goto ERR
//		}
//		//如果数据量过大阻塞在这里,等待inChan有空闲的位置！
//		select {
//		case conn.inChan <- data:
//		case <-conn.closeChan:
//			//closeChan关闭的时候
//			goto ERR
//
//		}
//	}
//ERR:
//	conn.Close()
//}
//
//func (conn *Connection) writeLoop() {
//	var (
//		data []byte
//		err  error
//	)
//	for {
//		select {
//		case data = <-conn.outChan:
//		case <-conn.closeChan:
//			goto ERR
//
//		}
//		if err = conn.wsConn.WriteMessage(websocket.TextMessage, data); err != nil {
//			goto ERR
//		}
//	}
//ERR:
//	conn.Close()
//}

//func (h *hub) run() {
//	for {
//		select {
//		case c := <-h.r:
//			h.c[c] = true
//			c.data.Ip = c.ws.RemoteAddr().String()
//			c.data.Type = "handshake"
//			c.data.UserList = user_list
//			data_b, _ := json.Marshal(c.data)
//			c.sc <- data_b
//		case c := <-h.u:
//			if _, ok := h.c[c]; ok {
//				delete(h.c, c)
//				close(c.sc)
//			}
//		case data := <-h.b:
//			for c := range h.c {
//				select {
//				case c.sc <- data:
//				default:
//					delete(h.c, c)
//					close(c.sc)
//				}
//			}
//		}
//	}
//}

//func Myws(w http.ResponseWriter, r *http.Request) {
//	ws, err := wu.Upgrade(w, r, nil)
//	if err != nil {
//		return
//	}
//	c := &connection{sc: make(chan []byte, 256), ws: ws, data: &Data{}}
//	h.r <- c
//	go c.writer()
//	c.reader()
//	defer func() {
//		c.data.Type = "logout"
//		user_list = del(user_list, c.data.User)
//		c.data.UserList = user_list
//		c.data.Content = c.data.User
//		data_b, _ := json.Marshal(c.data)
//		h.b <- data_b
//		h.r <- c
//	}()
//}
//
//func (c *connection) writer() {
//	for message := range c.sc {
//		c.ws.WriteMessage(websocket.TextMessage, message)
//	}
//	c.ws.Close()
//}
//
//var user_list = []string{}
//
//func (c *connection) reader() {
//	for {
//		_, message, err := c.ws.ReadMessage()
//		if err != nil {
//			h.r <- c
//			break
//		}
//		json.Unmarshal(message, &c.data)
//		switch c.data.Type {
//		case "login":
//			c.data.User = c.data.Content
//			c.data.From = c.data.User
//			user_list = append(user_list, c.data.User)
//			c.data.UserList = user_list
//			data_b, _ := json.Marshal(c.data)
//			h.b <- data_b
//		case "user":
//			c.data.Type = "user"
//			data_b, _ := json.Marshal(c.data)
//			h.b <- data_b
//		case "logout":
//			c.data.Type = "logout"
//			user_list = del(user_list, c.data.User)
//			data_b, _ := json.Marshal(c.data)
//			h.b <- data_b
//			h.r <- c
//		default:
//			fmt.Print("========default================")
//		}
//	}
//}
//
//func del(slice []string, user string) []string {
//	count := len(slice)
//	if count == 0 {
//		return slice
//	}
//	if count == 1 && slice[0] == user {
//		return []string{}
//	}
//	var n_slice = []string{}
//	for i := range slice {
//		if slice[i] == user && i == count {
//			return slice[:count]
//		} else if slice[i] == user {
//			n_slice = append(slice[:i], slice[i+1:]...)
//			break
//		}
//	}
//	fmt.Println(n_slice)
//	return n_slice
//}
