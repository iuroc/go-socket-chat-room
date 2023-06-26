package client

import (
	"fmt"
	"go-socket-chat-room/util"
	"log"
	"net"
	"sync"
)

func Client() {
	server, _ := net.Dial("tcp", ":8080")
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Print("请输入昵称：")
	var username string
	fmt.Scanln(&username)
	go receiveMessage(server)
	go sendMessage(server, username)
	wg.Wait()
}

func sendMessage(server net.Conn, username string) {
	for {
		var body string
		fmt.Scanln(&body)
		header := make(map[string]string)
		header["Username"] = username
		server.Write(util.MakeMessage(header, body))
	}
}

func receiveMessage(server net.Conn) {
	for {
		message := make([]byte, 1024)
		_, err := server.Read(message)
		if err != nil {
			break
		}
		response := util.ParseMessage(message)
		username := response.Header["Username"]
		body := response.Body
		log.Printf("【%s】%s\n", username, body)
	}
}
