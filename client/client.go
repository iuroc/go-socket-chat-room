package client

import (
	"fmt"
	"log"
	"net"
	"sync"
	"bytes"
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
		message := makeMessage(header, text)
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
		response := parseMessage(message)
		username := response.Header["Username"]
		text := response.Body
		fmt.Printf("【%s】%s\n", username, text)
	}
}


type messageData struct {
	// 消息头
	Header map[string]string
	// 消息体
	Body string
}

// 解析消息内容
func parseMessage(message []byte) messageData {
	data := messageData{}
	data.Header = make(map[string]string)
	parts := bytes.SplitN(message, []byte("\r\n\r\n"), 2)
	// 解析消息头
	if len(parts) > 0 {
		lines := bytes.Split(parts[0], []byte("\r\n"))
		for _, line := range lines {
			header := bytes.SplitN(line, []byte(":"), 2)
			if len(header) == 2 {
				key := string(bytes.TrimSpace(header[0]))
				value := string(bytes.TrimSpace(header[1]))
				data.Header[key] = value
			}
		}
	}
	// 解析消息体
	if len(parts) > 1 {
		data.Body = string(parts[1])
	}
	return data
}

// 生成消息
func makeMessage(header map[string]string, body string) []byte {
	var message bytes.Buffer
	for key, value := range header {
		message.WriteString(key + ": " + value + "\r\n")
	}
	message.WriteString("\r\n" + body)
	return message.Bytes()
}
