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
		client.Client()
	} else {
		server.Server()
	}
}
