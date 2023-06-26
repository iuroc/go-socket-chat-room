package util

import (
	"bytes"
	"fmt"
	"net"
)

func ParseMessageByConn(conn net.Conn) messageData {
	message := make([]byte, 1024)
	conn.Read(message)
	return ParseMessage(message)
}

// 解析消息内容
func ParseMessage(message []byte) messageData {
	data := messageData{
		Header: make(map[string]string),
		Body:   "",
	}
	parts := bytes.SplitN(message, []byte("\r\n\r\n"), 2)
	if len(parts) > 0 {
		headerLines := bytes.Split(parts[0], []byte("\r\n"))
		for _, line := range headerLines {
			header := bytes.SplitN(line, []byte(":"), 2)
			if len(header) == 2 {
				key := string(bytes.TrimSpace(header[0]))
				value := string(bytes.TrimSpace(header[1]))
				data.Header[key] = value
			}
		}
	}
	if len(parts) > 1 {
		data.Body = string(parts[1])
	}
	return data
}

// 消息内容数据
type messageData struct {
	Header map[string]string
	Body   string
}

// 生成消息数据
func MakeMessage(header map[string]string, body string) []byte {
	var message bytes.Buffer
	for key, value := range header {
		message.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}
	message.WriteString("\r\n" + body)
	return message.Bytes()
}
