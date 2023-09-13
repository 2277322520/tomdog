package tdnet

import "tomdog/tdface"

// BaseRouter 实现接口类之前，先嵌入这个基类
// 有的 XxxRouterImpl 不希望有 PreHandle 或者 AfterHandle
// 所以所有的 XxxRouterImpl 全部继承CaseRouter 的好处就是不实现这两个方法也可以实例化
type BaseRouter struct {
}

func (b *BaseRouter) PreHandle(request tdface.IRequest) {

}

func (b *BaseRouter) Handle(request tdface.IRequest) {

}

func (b *BaseRouter) AfterHandle(request tdface.IRequest) {

}
