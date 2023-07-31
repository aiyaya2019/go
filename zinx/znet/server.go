package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

//iServer的接口实现，定义一个Server的服务器模块
type Server struct {
	//服务器名称
	Name string
	//服务器绑定的ip版本
	IPVersion string
	//服务器监听的ip
	IP string
	//服务器监听的端口
	Port int
	//当前的Server添加一个router，server注册的链接对应的处理业务
	Router ziface.IRouter
}

//启动服务器
func (s *Server) Start() {
	fmt.Printf("[Start] Server Listenner at IP:%s, Port:%d is starting\n", s.IP, s.Port)
	//1、获取一个tcp的addr
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("resolve tcp addr error:", err)
		return
	}

	//2、监听服务器的地址
	listenner, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println("listen", s.IPVersion, "err", err)
		return
	}

	fmt.Println("start zinx server succ, ", s.Name, "succ, Listenning...")

	var cid uint32
	cid = 0

	//3、阻塞的等待客户端连接，处理客户端连接业务（读写）
	for {
		//如果有客户端连接过来，阻塞会返回
		conn, err := listenner.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err", err)
			continue
		}

		//将处理新连接的业务方法 和 conn 进行绑定 得到链接模块
		dealConn := NewConnection(conn, cid, s.Router)
		cid++

		//启动当前的链接业务处理
		go dealConn.Start()
	}

}

//停止服务器
func (s *Server) Stop() {
	//todo 将一些服务器的资源、状态或者一些已开辟的连接信息 进行停止或回收

}

//运行服务器
func (s *Server) Serve() {
	//启动server的服务功能
	s.Start()

	//todo 做启动服务器之后的业务

	//阻塞状态
	select {}

}

//路由功能：给当前的服务注册一个路由方法，供客户端的链接处理使用
func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
	fmt.Println("add Router succ")

}

//初始化server模块的方法
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "127.0.0.1",
		Port:      8999,
		Router:    nil,
	}
	return s
}
