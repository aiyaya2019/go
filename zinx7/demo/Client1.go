package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx/znet"
)

//模拟客户端
func main() {
	fmt.Println("client1 start...")
	time.Sleep(1 * time.Second)

	//1、直接连接远程服务器，得到一个conn连接
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client1 start err, exit:", err)
		return
	}

	for {
		//发送封包的message消息
		dp := znet.NewDataPack()

		binaryMsg, err := dp.Pack(znet.NewMsgPackage(1, []byte("zinxV0.6 client1 test message")))
		if err != nil {
			fmt.Println("pack error:", err)
			return
		}

		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println("write error:", err)
			return
		}

		//服务器给客户端回复一个message数据,

		//1、先读取流中的head部分，得到id和dataLen
		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head error:", err)
			break
		}
		//将二进制的head拆包到msg结构体中
		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("client1 unpack msgHead error:", err)
			break
		}
		if msgHead.GetMsgLen() > 0 {
			//2、在根据dataLen进行第二次读取，将data读出来

			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msgHead.GetMsgLen())
			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("client1 read msg data error:", err)
				break
			}
			fmt.Println("--->recv server msg: id=", msg.Id, ", dataLen=", msg.DataLen, ", data=", string(msg.Data))
		}

		//cpu阻塞
		time.Sleep(1 * time.Second)

	}

}
