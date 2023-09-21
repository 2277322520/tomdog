package tdface

type IMessage interface {
	
	// GetDataLen 获取消息数据长度
	GetDataLen() uint32
	
	// GetMsgId 获取消息 ID
	GetMsgId() uint32
	
	// GetData 获取消息内容
	GetData() []byte
	
	// SetDataLen 设置消息数据长度
	SetDataLen(uint32)
	
	// SetData 设置数据
	SetData([]byte)
}
