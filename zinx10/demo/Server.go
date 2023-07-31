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
type HelloRouter struct {
	znet.BaseRouter
}

//Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")

	//先读取客户端的数据，再回写ping
	fmt.Println("recv from client: msgId=", request.GetMsgId(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(200, []byte("ping...ping..."))
	if err != nil {
		fmt.Println("SendMsg error:", err)
	}
}

func (this *HelloRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call HelloRouter Handle...")

	//先读取客户端的数据，再回写ping
	fmt.Println("recv from client: msgId=", request.GetMsgId(), ", data=", string(request.GetData()))

	err := request.GetConnection().SendMsg(201, []byte("HelloRouter welcome"))
	if err != nil {
		fmt.Println("SendMsg error:", err)
	}
}

//创建连接之后执行钩子函数
func DoConnectionBegin(conn ziface.IConnectioin) {
	fmt.Println("=============================> DoConnectionBegin is Called...")
	if err := conn.SendMsg(202, []byte("DoConnection Begin...")); err != nil {
		fmt.Println("DoConnection Begin error:", err)
	}

	//给当前的连接设置一些属性
	fmt.Println("------------------------------> set conn name, home...")
	conn.SetProperty("name", "go go go")
	conn.SetProperty("baidu", "www.baidu.com")
	conn.SetProperty("home", "home home home")
}

func DoConnectionLost(conn ziface.IConnectioin) {
	fmt.Println("=============================> DoConnectionLost is Called...")
	fmt.Println("conn Id = ", conn.GetConnID(), " is lost...")

	//获取连接属性
	if name, err := conn.GetProperty("name"); err == nil {
		fmt.Println("name:=", name)
	}
	if baidu, err := conn.GetProperty("baidu"); err == nil {
		fmt.Println("baidu:=", baidu)
	}
	if home, err := conn.GetProperty("home"); err == nil {
		fmt.Println("home:=", home)
	}

}

func main() {
	//1、创建一个server句柄，使用Zinx的api
	s := znet.NewServer()

	//2、注册连接Hook钩子函数
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)

	//3、给当前框架添加自定义的router
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})

	//4、启动server
	s.Serve()
}
