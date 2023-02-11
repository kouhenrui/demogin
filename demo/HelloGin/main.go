package main

import (
	"HelloGin/src/controller/admin"
	"HelloGin/src/controller/async"
	"HelloGin/src/controller/upload"
	"HelloGin/src/controller/user"
	"HelloGin/src/controller/ws"
	"HelloGin/src/global"
	"HelloGin/src/routers"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

func main() {
	routers.Include(user.Routers, upload.Routers, async.Routers, admin.Routers, ws.Routers)
	r := routers.InitRoute()
	ws, _, ws_err := websocket.DefaultDialer.Dial(global.WSADDRESS, nil)
	if ws_err != nil {
		log.Fatal("dial:", ws_err)
	} else {
		fmt.Println("Connected to server")
	}
	ws.SetPongHandler(func(str string) error {
		fmt.Println("pong received", str)
		return nil
	})

	defer ws.Close()
	//go
	if err := r.Run(global.PORT); err != nil {
		fmt.Errorf("端口占用,err:%v\n", err)
	}
}
