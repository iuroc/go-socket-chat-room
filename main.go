package main

import (
	"flag"
	"go-socket-chat-room/client"
	"go-socket-chat-room/server"
)

func main() {
	clientMode := flag.Bool("client", false, "客户端模式")
	flag.Parse()
	if *clientMode {
		// 客户端
		client.Client()
	} else {
		// 服务端
		server.Server()
	}
}
