package znet

import (
	"fmt"
	"net"
	"zinx/utils"
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
	//当前的Server的消息管理模块，用来绑定MsgId和对应的处理业务api关系
	MsgHandler ziface.IMsgHandle

	//该server的连接管理器
	ConnMgr ziface.IConnManager

	//该Server创建连接之后自动调用Hook函数--OnConnStart
	OnConnStart func(conn ziface.IConnectioin)

	//该Server销毁连接之前自动调用Hook函数--OnConnStop
	OnConnStop func(conn ziface.IConnectioin)
}

//启动服务器
func (s *Server) Start() {
	fmt.Printf("[zinx]server Name:%s, ip:%s, port:%d is starting\n", utils.GlobalObject.Name, utils.GlobalObject.Host, utils.GlobalObject.TcpPort)
	fmt.Printf("[zinx]Version:%s, MaxConn:%d, MaxPackageSize:%d\n", utils.GlobalObject.Version, utils.GlobalObject.MaxConn, utils.GlobalObject.MaxPackageSize)

	//0、开启消息队列及worker工作池
	s.MsgHandler.StartWorkerPool()

	go func() { //协程，不会阻塞，相当于另类线程---相当于异步

		//1、获取一个tcp的addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error:", err)
			return
		}

		//2、监听服务器的地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}

		fmt.Println("start zinx server succ, ", s.Name, "succ, Listening...")

		var cid uint32
		cid = 0

		//3、阻塞的等待客户端连接，处理客户端连接业务（读写）
		for {
			//如果有客户端连接过来，阻塞会返回
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			//设置贵大连接个数的判断，如果超过最大连接，那么则关闭此新的的连接
			if s.ConnMgr.Len() >= utils.GlobalObject.MaxConn {
				//todo 给客户端响应一个超出最大连接的错误包
				fmt.Println("================================>too many connections maxConn = ", utils.GlobalObject.MaxConn)
				conn.Close()
				continue
			}

			//将处理新连接的业务方法 和 conn 进行绑定 得到链接模块
			dealConn := NewConnection(s, conn, cid, s.MsgHandler)
			cid++

			//启动当前的链接业务处理
			go dealConn.Start()
		}
	}()

}

//停止服务器
func (s *Server) Stop() {
	//将一些服务器的资源、状态或者一些已开辟的连接信息 进行停止或回收
	fmt.Println("[stop] zinx server name ", s.Name)
	s.ConnMgr.ClearConn()
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
func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.MsgHandler.AddRouter(msgId, router)
	fmt.Println("add Router success")
}

func (s *Server) GetConnMgr() ziface.IConnManager {
	return s.ConnMgr
}

//初始化server模块的方法
func NewServer() ziface.IServer {
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		MsgHandler: NewMsgHandle(),
		ConnMgr:    NewCConnManager(),
	}
	return s
}

//注册OnConnStart钩子函数的方法
func (s *Server) SetOnConnStart(hookFunc func(connection ziface.IConnectioin)) {
	s.OnConnStart = hookFunc
}

//注册OnConnStop钩子函数的方法
func (s *Server) SetOnConnStop(hookFunc func(connection ziface.IConnectioin)) {
	s.OnConnStop = hookFunc
}

//调用OnConnStart钩子函数的方法
func (s *Server) CallOnConnStart(conn ziface.IConnectioin) {
	if s.OnConnStart != nil {
		fmt.Println("-------------->call OnConnStart()...")
		s.OnConnStart(conn)
	}
}

//调用OnConnStop钩子函数的方法
func (s *Server) CallOnConnStop(conn ziface.IConnectioin) {
	if s.OnConnStop != nil {
		fmt.Println("-------------->call OnConnStop()...")
		s.OnConnStop(conn)
	}

}
