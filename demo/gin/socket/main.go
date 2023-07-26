package main

import (
	"fmt"
	"socket/router"
	"socket/wsManager"
)

func main() {
	//挂载路由
	router.Include(wsManager.Routers)

	//初始化路由器,加载中间件等
	r := router.InitRoute()
	if err := r.Run(":9999"); err != nil {
		fmt.Errorf("端口占用,err:%v\n", err)
	}
}
