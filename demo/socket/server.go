package main

/**
* @program: work_space
*
* @description:
*
* @author: khr
*
* @create: 2023-02-09 10:06
**/
import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	CoonTypeTCP int8 = 1 // tcp连接
	ConnTypeWS  int8 = 2 // websocket连接
)

func main() {
	router := mux.NewRouter()
	go h.run()
	router.HandleFunc("/ws", WSHandler)
	if err := http.ListenAndServe("127.0.0.1:8080", router); err != nil {
		fmt.Println("err:", err)
	}
}
