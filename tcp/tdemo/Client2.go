package main

import (
	"fmt"
	"io"
	"net"
	"tcp/tnet"
	"time"
)

func main() {
	fmt.Println("client start...")

	//建立网络连接并返回一个通用的 Conn 接口
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("connection error: ", err)
		return
	}
	fmt.Println("conn: ", conn)

	for {
		//发送封包的message消息
		dp := tnet.NewDataPack()

		binaryMsg, err := dp.Pack(tnet.NewMsgPackage(0, []byte("client2 msg...")))
		if err != nil {
			fmt.Println("pack error:", err)
			return
		}

		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println("write error:", err)
			return
		}

		//服务器给客户端回复一个message数据

		//1、先读取流中的head部分，得到id和msgLen
		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head error:", err)
			break
		}
		//将二进制的head拆包到msg结构体中
		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("client unpack msgHead error:", err)
			break
		}

		if msgHead.GetMsgLen() > 0 {
			//2、在根据msgLen进行第二次读取，将msg读出来
			msg := msgHead.(*tnet.Message)
			msg.Msg = make([]byte, msgHead.GetMsgLen())
			if _, err := io.ReadFull(conn, msg.Msg); err != nil {
				fmt.Println("client read msg data error:", err)
				break
			}
			fmt.Println("--->recv server msg: id=", msg.Id, ", dataLen=", msg.MsgLen, ", data=", string(msg.Msg))
		}

		time.Sleep(1 * time.Second)
	}

}
