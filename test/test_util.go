package test

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"time"
	"tomdog/tdnet"
	"tomdog/utils"
)

// ParseResponse conn 入参显然不是单例的，无需施加同步互斥手段
func ParseResponse(doneChannel chan bool, conn net.Conn) {
	fmt.Println("client test ... start")
	time.Sleep(3 * time.Second)

	// 创建拆包封包对象
	dp := tdnet.NewDataPack()

	for {
		//	先读取出流中的 head 部分
		header := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(conn, header)
		if err != nil {
			fmt.Println("read head error")
			continue
		}

		// 将 header 字节流拆包到 msg 中
		msg, err := dp.Unpack(header)
		if err != nil {
			fmt.Println("server unpack err:", err)
			return
		}

		// 根据 dataLen 将数据部分读取到 msg
		if msg.GetDataLen() > 0 {

			//这行代码 msg := msg.(*tdnet.Message) 是类型断言的用法。在Go中，接口可以容纳不同类型的值。
			//当你从接口中获取一个值时，需要将其转换为特定的类型才能访问其原始类型的方法和字段。
			//在这个代码中，msg 最初是一个 tdface.IMessage 接口类型的变量，它可能包含任何实现 tdface.IMessage 接口的类型的值。
			//但在后续的代码中，你想要访问 Message 类型特有的方法和字段，因此你需要将 msg 转换为 *tdnet.Message 类型，
			//这是 tdnet.Message 类型的指针。
			//具体来说，msg.(*tdnet.Message) 这行代码尝试将 msg 转换为 *tdnet.Message 类型的指针。如果 msg 实际上不是 *tdnet.Message 类型的指针，
			//而是其他类型，或者根本不是指针，那么这个操作会引发运行时的 panic。
			msg := msg.(*tdnet.Message)
			msg.Data = make([]byte, msg.GetDataLen())
			//msg.Data = make([]byte, msg.GetDataLen()*2+8)
			_, err := io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("server unpack err:", err)
				return
			}

			utils.Logging("==> Recv Msg:ID=" + strconv.Itoa(int(msg.RouterId)) + ",Len=" + strconv.Itoa(int(msg.DataLen)) + ",data=" + string(msg.Data))

			// 相应处理完毕，客户端应该关闭连接，通过管道告知主协程
			doneChannel <- true
		}

	}
}
