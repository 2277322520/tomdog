package tdnet

import (
	"tomdog/tdface"
)

var multiClientRouterId uint32 = 10004

type MultiClientRouter struct {
	BaseRouter
}

func (b *MultiClientRouter) Handle(request tdface.IRequest) {

}

func (router *MultiClientRouter) GetRouterId() uint32 {
	return multiClientRouterId
}
