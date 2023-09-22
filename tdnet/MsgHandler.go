package tdnet

import (
	"fmt"
	"strconv"
	"tomdog/tdface"
)

type MsgHandler struct {
	
	// 接口 map
	apis map[uint32]tdface.IRouter
}

func (m *MsgHandler) DoMsgHandler(request tdface.IRequest) {
	msgid := request.GetData().GetMsgId()
	if _, ok := m.apis[msgid]; !ok {
		panic("no suitable optional route selector")
	}
	
	m.apis[msgid].PreHandle(request)
	m.apis[msgid].Handle(request)
	m.apis[msgid].AfterHandle(request)
}

func (m *MsgHandler) AddRouter(msgid uint32, router tdface.IRouter) {
	// 重复添加路由
	if _, ok := m.apis[msgid]; ok {
		panic("repeat router, msgid =" + strconv.Itoa(int(msgid)))
	}
	
	m.apis[msgid] = router
	fmt.Println("router add success,msgid =", msgid)
}

// NewMsgHandler 构造函数
func NewMsgHandler() *MsgHandler {
	return &MsgHandler{
		apis: make(map[uint32]tdface.IRouter),
	}
}
