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
)

//var (
//	upgrade = websocket.Upgrader{
//		//允许跨域
//		CheckOrigin: func(r *http.Request) bool {
//			//可以添加用户认证，错误返回false
//
//			return true
//		},
//	}
//)

func main() {
	routers.Include(user.Routers, upload.Routers, async.Routers, admin.Routers, ws.Routers)
	r := routers.InitRoute()

	//http.HandleFunc("/ws", ws.Myws)
	//http.ListenAndServe(global.PORT, nil)
	if err := r.Run(global.PORT); err != nil {
		fmt.Errorf("端口占用,err:%v\n", err)
	}
}
