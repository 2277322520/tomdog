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
	HandleAPI tdface.HandFunc
	
	ExitBuffChan chan bool
	
	// 当前连接的处理方法 Router
	Router tdface.IRouter
}

// NewConnection 构造函数
func NewConnection(conn *net.TCPConn, connID uint32, callbackApi tdface.HandFunc, router tdface.IRouter) *Connection {
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		HandleAPI:    callbackApi,
		ExitBuffChan: make(chan bool, 1),
		Router:       router,
	}
	
	return c
}

// StartReader 处理 conn 读数据的协程
// 在 Go 语言中，成员方法的接收者（Receiver）用于确定该方法属于哪个类型，它是通过在函数名前添加一个参数来实现的。这个参数通常是一个接收
// 者类型的变量，它可以是类型的值或指针。在你的代码中，StartReader 方法的接收者是 c *Connection，这表示它是属于 *Connection 类型的。
// 在 Go 中，如果你想要在方法内部修改接收者的状态，通常会使用指针作为接收者。因为指针传递可以修改原始对象的状态，而值传递只会在
// 方法内部操作副本，不会影响原始对象。在你的代码中，StartReader 方法可能需要修改 Connection 对象的状态，所以使用指针是更合适的选择。
// 这也是 Go 语言中方法声明的一种习惯用法，以便在方法内部可以修改接收者的状态。如果你使用值接收者，那么在方法内部修改的将会是
// 接收者的副本，而不是原始对象，这可能不是你想要的行为。因此，通常建议使用指针接收者来声明方法，以便可以修改接收者的状态。
func (c *Connection) StartReader() {
	
	fmt.Println("reader goroutine is running")
	defer fmt.Println(c.RemoteAddr().String(), " conn reader exit!")
	defer c.Stop()
	
	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		
		if err != nil {
			fmt.Println("receive buf error, ", err)
			c.ExitBuffChan <- true
			continue
		}
		
		req := Request{
			conn: c,
			data: buf,
		}
		
		go func(request tdface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.AfterHandle(request)
			//	如果传递 req 而不是 &req，那么你传递给 PreHandle、Handle 和 AfterHandle 方法的将不再是接口类型，
			//	而是 Request 结构体的值。这可能会导致类型不匹配的错误，因为方法期望的是接口类型。
		}(&req)
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
