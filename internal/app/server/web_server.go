package server

import (
	"GoIM/pkg/auth"
	"GoIM/pkg/utils"
	"fmt"
	"log"
	"net/http"
)

type WebServer struct {
	Addr    string
	Port    int
	RpcPort int
}

func (ws *WebServer) onInit(wsPort int, rpcPort int) {
	ws.Addr = "0.0.0.0"
	ws.Port = wsPort
	ws.RpcPort = rpcPort
	fmt.Println("web server on init:", wsPort)
}

func (ws *WebServer) run() {
	r := utils.HttpRouter()
	r.Use(auth.AccessTokenMiddleware) // 添加需要使用的中间件
	r.Exclude("/", http.HandlerFunc(serveHome))
	r.Exclude("/login", http.HandlerFunc(login))
	r.Add("/ws", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ServeWs(hub, ws.Addr, ws.RpcPort, w, r)
	}))
	host := ws.Addr + ":" + utils.IntToString(ws.Port)
	fmt.Println("http.ListenAndServe:", host)
	err := http.ListenAndServe(host, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
