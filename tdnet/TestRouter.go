package tdnet

import (
	"fmt"
	"tomdog/tdface"
)

type TestRouter struct {
	// PingRouter 继承自 BaseRouter
	BaseRouter
}

func (b *PingRouter) Handle(request tdface.IRequest) {
	fmt.Println("call ping router handle")
	fmt.Println("recv from client: msgid = ", request.GetData().GetRouterId(), ", data=", string(request.GetData().GetData()))
	
	err := request.GetConnection().SendMsg(20002, []byte("pong...pong..pong..."))
	if err != nil {
		fmt.Println("send response msg error ", err)
	}
}

func (b *PingRouter) AfterHandle(request tdface.IRequest) {
	fmt.Println("call ping router after handle")
}
