package server

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	rwLock              *sync.RWMutex
	validate            *validator.Validate
	hub                 *Hub
	rpcSvr              RPCServer
	webSvr              WebServer
	sendMsgAllCount     uint64
	sendMsgFailedCount  uint64
	sendMsgSuccessCount uint64
	userCount           uint64

	sendMsgAllCountLock sync.RWMutex
)

func Init(rpcPort, wsPort int) {
	rwLock = new(sync.RWMutex)
	validate = validator.New()
	hub = NewHub()
	webSvr.onInit(wsPort, rpcPort)
	rpcSvr.onInit(rpcPort)
}

func Run() {
	go hub.run()
	go rpcSvr.run()
	webSvr.run()

}
