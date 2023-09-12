package tdnet

import (
	"fmt"
	"net"
	"tomdog/tdface"
)

type Connection struct {

	// 当前连接的 socket TCP 套接字
	Conn *net.TCPConn

	// 当前连接 ID，也可以称为 SessionID,ID 全局唯一
	ConnID uint32

	// 当前连接是否是关闭状态
	isClosed bool

	// 处理当前连接的 API
	handleAPI tdface.HandFunc

	ExitBuffChan chan bool
}

// NewConnection 构造函数
func NewConnection(conn *net.TCPConn, connID uint32, callbackApi tdface.HandFunc) *Connection {
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		handleAPI:    callbackApi,
		ExitBuffChan: make(chan bool, 1),
	}

	return c
}

// StartReader 处理 conn 读数据的协程
func (c *Connection) StartReader() {

	fmt.Println("reader goroutine is running")
	defer fmt.Println(c.RemoteAddr().String(), " conn reader exit!")
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)

		if err != nil {
			fmt.Println("receive buf error, ", err)
			c.ExitBuffChan <- true
			continue
		}

		if err := c.handleAPI(c.Conn, buf, cnt); err != nil {
			fmt.Println()
			c.ExitBuffChan <- true
			return
		}
	}
}

func (c *Connection) Start() {
	go c.StartReader()

	for true {
		// select 语句用于处理多个通道操作。它类似于 switch 语句，但专门用于通道操作。
		// select 允许你在多个通道之间进行非阻塞的选择，从而实现并发控制。
		select {
		case <-c.ExitBuffChan:
			// 如果成功从 c.ExitBuffChan 通道中接收到数据，那么 return 语句将会退出当前函数，也就是 Start() 函数。
			// 这意味着当从 c.ExitBuffChan 中接收到数据时，整个 Start() 函数的执行都会结束，协程将退出。
			return
		}
	}
}

func (c *Connection) Stop() {
	if c.isClosed {
		return
	}

	c.isClosed = true

	// todo Connection Stop() 如果用户注册了该连接的回调业务，则应该在此处调用

	// 关闭连接
	err := c.Conn.Close()
	if err != nil {
		fmt.Println("close connection error")
		return
	}

	// 向管道发送通知，告知主协程当前连接已经成功关闭
	c.ExitBuffChan <- true
	// 关闭管道
	close(c.ExitBuffChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
