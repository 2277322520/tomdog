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
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping...\n"))
	
	if err != nil {
		fmt.Println(" handle ping router error")
	}
}

func (b *PingRouter) AfterHandle(request tdface.IRequest) {
	fmt.Println("call ping router after handle")
	
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping router\n"))
	
	if err != nil {
		fmt.Println("after handle  ping router error")
	}
}
