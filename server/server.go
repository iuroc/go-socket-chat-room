package server

import (
	"log"
	"net"
)

func Server() {
	listener, _ := net.Listen("tcp", ":8080")
	defer listener.Close()
	clientList := []net.Conn{}
	for {
		client, _ := listener.Accept()
		go clientConnect(client, &clientList)
	}
}

// 发现客户端连接
func clientConnect(client net.Conn, clientList *[]net.Conn) {
	// 当前客户端，在客户端列表中的索引
	index := len(*clientList)
	// 将当前客户端，加入到客户端列表中
	*clientList = append(*clientList, client)
	for {
		message := make([]byte, 1024)
		// 获取当前客户端的请求消息数据
		if _, err := client.Read(message); err != nil {
			// 发生错误，从客户端列表中移除当前客户端
			*clientList = append((*clientList)[:index], (*clientList)[index+1:]...)
			break
		}
		// 将当前客户端的消息，广播给所有的客户端
		for _, item := range *clientList {
			if _, err := item.Write(message); err != nil {
				log.Println(err.Error())
			}
		}
	}
}
