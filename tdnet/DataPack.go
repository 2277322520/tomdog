package tdnet

import "tomdog/tdface"

// DataPack 封包拆包实现类
type DataPack struct {
}

// NewPack 构造函数
func NewPack() *DataPack {
	return &DataPack{}
}

// GetHeadLen 获取包头长度
func (d DataPack) GetHeadLen() uint32 {
	// 包头长度=len(Message.id)+len(Message.DataLen)
	return HeadLen
}

func (d DataPack) Pack(msg tdface.IMessage) ([]byte, error) {
	//TODO implement me
	panic("implement me")
}

func (d DataPack) Unpack(bytes []byte) (tdface.IMessage, error) {
	//TODO implement me
	panic("implement me")
}
