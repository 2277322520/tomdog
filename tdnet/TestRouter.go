package tdnet

import (
	"fmt"
	"strconv"
	"tomdog/tdface"
	"tomdog/utils"
)

var testRouterId uint32 = 10001

type TestRouter struct {
	// PingRouter 继承自 BaseRouter
	BaseRouter
}

func (b *TestRouter) Handle(request tdface.IRequest) {
	utils.Logging("call test router after handle")
	utils.Logging("recv from client: msgid = " + strconv.Itoa(int(request.GetData().GetRouterId())) + ", data=" + string(request.GetData().GetData()))

	err := request.GetConnection().SendMsg(20002, []byte("tested...tested..tested..."))
	if err != nil {
		fmt.Println("send response msg error ", err)
	}
}

func (b *TestRouter) AfterHandle(request tdface.IRequest) {
	utils.Logging("call test router after handle")
}

func (router *TestRouter) GetRouterId() uint32 {
	return testRouterId
}
