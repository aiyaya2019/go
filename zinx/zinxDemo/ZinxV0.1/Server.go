package main

import "zinx/znet"

//基于zinx框架开发的服务器应用程序

func main() {
	//1、创建一个server句柄，使用Zinx的api
	s := znet.NewServer("[zinx v0.1]")
	//2、启动server
	s.Serve()
}
