package test

import (
	"fmt"
	"net"
	"testing"
	"time"
	"tomdog/tdnet"
	"tomdog/utils"
)

func TestDataPackClient(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial err:", err)
	}

	doRequest(err, conn)

	doneChannel := make(chan bool)

	go ParseResponse(doneChannel, conn)

	// 客户端阻塞
	select {
	case done := <-doneChannel:
		// 处理接收到的数据
		if done {
			utils.Logging("client 001 got response!!!")
		}
	case <-time.After(5 * time.Second):
		// 超时处理，五秒钟内其他分支没有管道返回则认为请求超时
		utils.Logging("client 001 time out!!!")
	}
}

func doRequest(err error, conn net.Conn) bool {
	// 创建一个封包对象
	dp := tdnet.NewDataPack()
	message_1 := &tdnet.Message{
		RouterId: 10001,
		DataLen:  5,
		Data:     []byte{'h', 'e', 'l', 'l', 'o'},
	}

	sendMessage_1, err := dp.Pack(message_1)
	if err != nil {
		fmt.Println("pack message err,", sendMessage_1)
		return true
	}

	message_2 := &tdnet.Message{
		RouterId: 10002,
		DataLen:  7,
		Data:     []byte{'w', 'o', 'r', 'l', 'd', '!', '!'},
	}

	sendMessage_2, err := dp.Pack(message_2)
	if err != nil {
		fmt.Println("pack message err,", sendMessage_2)
		return true
	}

	data := append(sendMessage_1, sendMessage_2...)

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("write data err:", data)
	}
	return false
}
