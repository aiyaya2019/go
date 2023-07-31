package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

//基于zinx框架开发的服务器应用程序

//ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

//Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")

	//先读取客户端的数据，再回写ping
	fmt.Println("recv from client: msgId=", request.GetMsgId(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(1, []byte("ping...ping..."))
	if err != nil {
		fmt.Println("SendMsg error:", err)
	}

}

func main() {
	//1、创建一个server句柄，使用Zinx的api
	s := znet.NewServer("[zinx v0.5.1]")

	s.AddRouter(&PingRouter{})

	//2、启动server
	s.Serve()
}
