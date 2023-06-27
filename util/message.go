package util

import "bytes"

type Message struct {
	// 消息头
	Header map[string]string
	// 消息体
	Body string
}

// 解析消息内容
func ParseMessage(message []byte) Message {
	data := Message{}
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
func MakeMessage(header map[string]string, body string) []byte {
	var message bytes.Buffer
	for key, value := range header {
		message.WriteString(key + ": " + value + "\r\n")
	}
	message.WriteString("\r\n" + body)
	return message.Bytes()
}
