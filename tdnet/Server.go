package tdnet

import (
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
}

func (s Server) Start() {
	fmt.Printf("[START] server listenner at IP: %s, port %d, is starting\n", s.IP, s.Port)

	go func() {
		// 1、获取一个 ip 地址
		addr, _error := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if _error != nil {
			fmt.Println("resolve tcp addr err:", _error)
			return
		}

		// 2、监听服务器地址
		listener, _error := net.ListenTCP(s.IPVersion, addr)
		if _error != nil {
			fmt.Println("listen:", s.IPVersion, "err:", _error)
			return
		}
		// 已经监听成功
		fmt.Println("start tomdog server", s.Name, "success, now listening...")

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

			// 3.3 todo Server.Start() 处理该信链接请求的业务方法

			// 这里暂时做一个最大 512 字节的回显服务
			go func() {
				// 不断循环，从客户端获取数据
				for {
					buf := make([]byte, 512)
					cnt, err := connection.Read(buf)
					if err != nil {
						fmt.Println("receive buf error")
						continue
					}
					// 回显
					if _, err := connection.Write(buf[:cnt]); err != nil {
						fmt.Println("write back error", err)
						continue
					}
				}
			}()

		}
		//	end for
	}()

}

func (s Server) Stop() {
	fmt.Println("[STOP] tomdog server ,name ", s.Name)

	// todo 关闭资源
}

func (s Server) Serve() {
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
	}

	return s
}
