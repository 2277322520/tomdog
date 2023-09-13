package tdnet

import (
	"errors"
	"fmt"
	"net"
	"tomdog/tdface"
)

type Server struct {
	// 服务器的名称
	Name string
	// IPV4或者其他
	IPVersion string
	// 服务器绑定的 IP 地址，点分十进制表示
	IP string
	// 服务器绑定的端口
	Port int
	// 当前 Server 由用户绑定回调的 router
	Router tdface.IRouter
}

func (s *Server) AddRouter(router tdface.IRouter) {
	s.Router = router
	
	fmt.Println("Add Router success!!!")
}

// CallBackToClient 定义当前客户端的 Handle API
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[conn handle] call back to client")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write buf error ", err)
		return errors.New("call back error")
	}
	
	return nil
}

func (s *Server) Start() {
	fmt.Printf("[START] server listenner at IP: %s, port %d, is starting\n", s.IP, s.Port)
	
	go func() {
		// 1、获取一个 ip 地址
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err:", err)
			return
		}
		
		// 2、监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen:", s.IPVersion, "err:", err)
			return
		}
		// 已经监听成功
		fmt.Println("start tomdog server", s.Name, "success, now listening...")
		
		// sever.go 应该有一个自动生成 connID 的方法，并且生成的ID应该满足要求
		var cid uint32
		cid = 0
		
		// 3、启动网络连接
		for {
			// 3.1 阻塞等待客户端建立连接请求
			connection, err := listener.AcceptTCP()
			if err != nil {
				// 获取连接失败
				fmt.Println("accept error")
				continue
			}
			
			// 3.2 todo Server.Start() 设置服务器最大连接控制，如果超过最大连接，则关闭最新的链接
			
			// 3.3 Server.Start() 处理该信链接请求的业务方法
			dealConn := NewConnection(connection, cid, CallBackToClient, s.Router)
			cid++
			
			go dealConn.Start()
		}
		//	end for
	}()
	
}

func (s *Server) Stop() {
	fmt.Println("[STOP] tomdog server ,name ", s.Name)
	
	// todo 关闭资源
}

func (s *Server) Serve() {
	s.Start()
	
	// todo Server.Serve
	select {}
}

func NewServer(name string) tdface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      7077,
		Router:    nil,
	}
	
	return s
}
