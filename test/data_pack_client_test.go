package test

import (
	"fmt"
	"net"
	"testing"
	"tomdog/tdnet"
)

func TestDataPackClient(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client dial err:", err)
	}
	
	// 创建一个封包对象
	dp := tdnet.NewDataPack()
	message_1 := &tdnet.Message{
		Id:      0,
		DataLen: 5,
		Data:    []byte{'h', 'e', 'l', 'l', 'o'},
	}
	
	sendMessage_1, err := dp.Pack(message_1)
	if err != nil {
		fmt.Println("pack message err,", sendMessage_1)
		return
	}
	
	message_2 := &tdnet.Message{
		Id:      0,
		DataLen: 7,
		Data:    []byte{'w', 'o', 'r', 'l', 'd', '!'},
	}
	
	sendMessage_2, err := dp.Pack(message_2)
	if err != nil {
		fmt.Println("pack message err,", sendMessage_2)
		return
	}
	
	data := append(sendMessage_1, sendMessage_2...)
	
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("write data err:", data)
	}
	
	// 客户端阻塞
	select {}
}
