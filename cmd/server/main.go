// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"GoIM/internal/app/server"
	"flag"
	"fmt"
)

func main() {
	var rpcPort int
	flag.IntVar(&rpcPort, "r", 8081, "rpc port")
	var port int
	flag.IntVar(&port, "w", 8080, "http port")
	flag.Parse()
	fmt.Printf("rpc port:%v\nweb port:%v\n",
		rpcPort, port)
	server.Init(rpcPort, port)
	server.Run()
}
