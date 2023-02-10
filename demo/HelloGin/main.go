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

func main() {
	routers.Include(user.Routers, upload.Routers, async.Routers, admin.Routers, ws.Routers)
	r := routers.InitRoute()
	if err := r.Run(global.PORT); err != nil {
		fmt.Errorf("端口占用,err:%v\n", err)
	}
}
