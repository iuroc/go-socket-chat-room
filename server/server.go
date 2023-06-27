package server

import (
	"log"
	"net"
)

// 启动服务端
func Server() {
	// 创建服务端连接
	listener, err := net.Listen("tcp", ":7758")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("服务端正在运行")
	// 创建客户端队列
	clientList := make(map[int]net.Conn)
	// 客户端 ID 累加生成器
	var clientId int
	// 循环监听客户端连接
	for {
		client, err := listener.Accept()
		if err != nil {
			log.Println(err.Error())
		}
		// 将客户端加入队列
		clientList[clientId] = client
		log.Printf("在线人数：%d\n", len(clientList))
		// 开始为客户端提供服务
		go handleClient(client, clientId, clientList)
		clientId++
	}
}

// 处理客户端连接
func handleClient(client net.Conn, clientId int, clientList map[int]net.Conn) {
	for {
		// 获取客户端消息内容
		message := make([]byte, 1024)
		_, err := client.Read(message)
		// 判断客户端离线
		if err != nil {
			delete(clientList, clientId)
			log.Printf("在线人数：%d\n", len(clientList))
			return
		}
		// 广播转发消息
		for _, item := range clientList {
			if _, err := item.Write(message); err != nil {
				log.Println(err.Error())
			}
		}
	}
}
