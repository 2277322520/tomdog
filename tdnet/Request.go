package tdnet

import "tomdog/tdface"

type Request struct {
	
	// 已经和客户端建立的连接
	conn tdface.IConnection
	
	// 请求的数据
	data []byte
}

func (r Request) GetConnection() tdface.IConnection {
	return r.conn
}

func (r Request) GetData() []byte {
	return r.data
}
