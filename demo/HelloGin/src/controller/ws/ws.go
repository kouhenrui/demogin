package ws

import (
	"HelloGin/src/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

// 设置websocket
// CheckOrigin防止跨站点的请求伪造
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Routers(e *gin.Engine) {

	wsGroup := e.Group("/ws")
	{
		wsGroup.GET("/connect", ws)
	}

}

// websocket实现
func ws(c *gin.Context) {
	fmt.Print("进入ws连接")
	result := global.NewResult(c)
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	fmt.Println(err)
	if err != nil {
		//return
		result.Err("连接错误1")
		return
	}
	//fmt.Println(ws)
	defer ws.Close() //返回前关闭
	for {
		//读取ws中的数据

		fmt.Println("收取信息2")
		mt, message, err := ws.ReadMessage()
		fmt.Println(mt, string(message), "打印获取的信息")
		if err != nil {
			result.Err("连接错误")
			return
		}
		//fmt.Print(string(message), "message")
		//写入ws数据
		err = ws.WriteMessage(mt, []byte("welcome to ws"))
		if err != nil {
			result.Err("连接错误3")
			return
		}
	}
	fmt.Println("我们结束聊天啦")
}

// func WSCONNECT() {
//	defer c.Close()
//	defer close(done)
//	 upGrader.Upgrade(c.Writer, c.Request, nil)
//	for {
//		c.SetReadDeadline(time.Now().Add(timeoutDuration))
//		_, message, err := c.ReadMessage()
//		if err != nil {
//			log.Println("read:", err)
//			return
//		}
//		if len(message) >= 2 {
//			message = message[2:]
//		}
//		log.Printf("recv: %s", message)
//	}
//}()
