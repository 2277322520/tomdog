package tdface

type IMessage interface {
	
	// GetDataLen 获取消息数据长度
	GetDataLen() uint32
	
	// GetMsgId 获取消息 ID
	GetMsgId() uint32
	
	// GetData 获取消息内容
	GetData() []byte
	
	// 设置消息数据长度
	setDataLen(uint32)
}


