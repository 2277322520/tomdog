package tdnet

import (
	"fmt"
	"tomdog/tdface"
)

type PingRouter struct {
	// PingRouter 继承自 BaseRouter
	BaseRouter
}

func (b *TestRouter) Handle(request tdface.IRequest) {
	fmt.Println("call test router handle")
	fmt.Println("recv from client: msgid = ", request.GetData().GetRouterId(), ", data=", string(request.GetData().GetData()))
	
	err := request.GetConnection().SendMsg(20001, []byte("test...test..test..."))
	if err != nil {
		fmt.Println("send response msg error ", err)
	}
}

func (b *TestRouter) AfterHandle(request tdface.IRequest) {
	fmt.Println("call test router after handle")
	
	//_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping router\n"))
	//
	//if err != nil {
	//	fmt.Println("after handle  ping router error")
	//}
}
