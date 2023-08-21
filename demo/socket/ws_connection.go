package main

/**
* @program: work_space
*
* @description:
*
* @author: khr
*
* @create: 2023-02-09 10:07
**/
import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type connection struct {
	ws   *websocket.Conn
	sc   chan []byte
	data *Data
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 65536,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WSHandler(w http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	conn := &Conn{
		CoonType: ConnTypeWS,
		WS:       wsConn,
	}
	DoConn(conn)
}

//func myws(w http.ResponseWriter, r *http.Request) {
//	ws, err := upgrader.Upgrade(w, r, nil)
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