package client

import (
	"fmt"
	"go-socket-chat-room/util"
	"log"
	"net"
	"sync"
)

// 启动客户端
func Client() {
	server, err := net.Dial("tcp", ":7758")
	if err != nil {
		log.Fatal(err.Error())
	}
	var username string
	fmt.Print("请输入昵称：")
	fmt.Scanln(&username)
	var wg sync.WaitGroup
	wg.Add(2)
	go sender(server, username)
	go receiver(server)
	wg.Wait()
}

// 消息发送线程
func sender(server net.Conn, username string) {
	for {
		var text string
		fmt.Scanln(&text)
		header := make(map[string]string)
		header["Username"] = username
		message := util.MakeMessage(header, text)
		if _, err := server.Write(message); err != nil {
			log.Fatal(err.Error())
		}
	}
}

// 消息接收线程
func receiver(server net.Conn) {
	for {
		message := make([]byte, 1024)
		if _, err := server.Read(message); err != nil {
			log.Fatal(err.Error())
		}
		response := util.ParseMessage(message)
		username := response.Header["Username"]
		text := response.Body
		fmt.Printf("【%s】%s\n", username, text)
	}
}
