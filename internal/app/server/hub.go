// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package server

import (
	"GoIM/pkg/auth"
	"GoIM/pkg/gateway"
	"GoIM/pkg/utils"
	"encoding/json"
	"fmt"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	auths map[string]*Client
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		auths:      make(map[string]*Client),
	}
}

type UserMessage struct {
	Msg string
	To  string
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			fmt.Println("client register")
			h.clients[client] = true
		case client := <-h.unregister:
			fmt.Println("client unregister")
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			var umsg UserMessage
			err := json.Unmarshal(message, &umsg)
			if err != nil {
				fmt.Println("err:", err)
				return
			}
			authUser := auth.GetUserAuthInfoByUid(umsg.To)
			if authUser != nil {
				client := h.auths[authUser.AccessToken]
				if client != nil {
					client.send <- []byte(umsg.Msg)
				} else {
					fmt.Println(umsg.To + "不在同一台服务器")
					clientMsg := gateway.NewClientMsg("text", utils.Int64ToString(authUser.Id), umsg.Msg)
					err := json.Unmarshal(message, clientMsg)
					if err != nil {
						fmt.Println("json解码异常：", err.Error())
					} else {
						fmt.Println("通过rpc转发", clientMsg)
						err = gateway.SendMsg(clientMsg)
						if err != nil {
							fmt.Println("发送失败", err.Error())
						}
					}
				}
			}

		}
	}
}
