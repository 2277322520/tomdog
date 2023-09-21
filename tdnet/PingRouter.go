package tdnet

import (
	"fmt"
	"tomdog/tdface"
)

type PingRouter struct {
	// PingRouter 继承自 BaseRouter
	BaseRouter
}

func (b *PingRouter) Handle(request tdface.IRequest) {
	fmt.Println("call ping router handle")
	fmt.Println("recv from client: msgid = ", request.GetData().GetMsgId(), ", data=", string(request.GetData().GetData()))
	
	err := request.GetConnection().SendMsg(1001, []byte("pong...pong..pong..."))
	if err != nil {
		fmt.Println("send response msg error ", err)
	}
}

func (b *PingRouter) AfterHandle(request tdface.IRequest) {
	fmt.Println("call ping router after handle")
	
	//_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping router\n"))
	//
	//if err != nil {
	//	fmt.Println("after handle  ping router error")
	//}
}
