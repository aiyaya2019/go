package pnet

import (
	"demo1/paperlesstcp/piface"
	"demo1/paperlesstcp/utils"
	"fmt"
)

type Server struct {
	//服务器名称
	Name string
	//服务器绑定的ip版本
	IpVersion string
	//服务器监听的ip
	Ip string
	//服务器监听的端口
	Port int
}

//初始化server模块的方法
func NewServer() piface.IServer {
	s := &Server{
		//服务器名称
		Name: utils.GlobalObject.Name,
		//服务器绑定的ip版本
		IpVersion: utils.GlobalObject.IpVersion,
		//服务器监听的ip
		Ip: utils.GlobalObject.Ip,
		//服务器监听的端口
		Port: utils.GlobalObject.Port,
	}
	return s
}

//运行服务器
func (s *Server) Server() {
	s.Start()

}

//启动服务器
func (s *Server) Start() {
	fmt.Println(utils.GlobalObject.Name)
	fmt.Printf("server name:#{s.Name}, ip:#{s.Ip}, port:#{s.Port} is starting")

}

//停止服务器
func (s *Server) Stop() {

}
