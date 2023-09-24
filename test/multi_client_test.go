package test

import (
	"fmt"
	"net"
	"testing"
	"time"
	"tomdog/tdnet"
	"tomdog/utils"
)

//TestMultiClient 第一个客户端
func TestMultiClient(t *testing.T) {
	go dotestclient001()
}

func dotestclient001() {
	conn, err := buildConn()

	data, err := buildRequest()
	if err != nil {
		utils.Logging("building request data error")
	}

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("write data err:", data)
	}

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

func buildRequest() ([]byte, error) {
	// 创建一个封包对象
	dp := tdnet.NewDataPack()
	message := &tdnet.Message{
		RouterId: 10005,
		DataLen:  11,
		Data:     []byte{'h', 'e', 'l', 'l', 'o', ' ', 'i', ' ', 'a', 'm', ' ', '0', '0', '1'},
	}

	sendMessage, err := dp.Pack(message)
	if err != nil {
		fmt.Println("pack message err,", sendMessage)
		return nil, err
	}

	return sendMessage, nil
}

func buildConn() (net.Conn, error) {
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial err:", err)
	}
	return conn, err
}
