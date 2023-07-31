package tnet

import (
	"fmt"
	"net"
	"tcp/tiface"
	"tcp/utils"
)

type Server struct {
	IPVersion  string              //ip版本
	IP         string              //ip
	Port       int                 //端口
	Name       string              //服务器名称
	Version    string              //服务器版本号
	MsgHandler tiface.IMsgHandle   //当前的Server的消息管理模块，用来绑定MsgId和对应的处理业务api关系
	ConnMgr    tiface.IConnManager //该server的连接管理器
}

// NewServer 初始化server模块的方法
func NewServer() tiface.IServer {
	s := &Server{
		IPVersion:  utils.GlobalObject.IPVersion,
		IP:         utils.GlobalObject.IP,
		Port:       utils.GlobalObject.Port,
		Name:       utils.GlobalObject.Name,
		Version:    utils.GlobalObject.Version,
		MsgHandler: NewMsgHandle(),
		ConnMgr:    NewConnManager(),
	}

	return s
}

// Start 启动服务器方法
func (s *Server) Start() {
	fmt.Printf("server name: %s, ip:%s, port:%d is starting\n", utils.GlobalObject.Name, utils.GlobalObject.IP, utils.GlobalObject.Port)

	//1、开启消息队列及worker工作池
	s.MsgHandler.StartWorkerPool()

	go func() { //协程，不会阻塞，相当于另类线程---相当于异步
		//2、获取一个tcp的addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error:", err)
			return
		}
		//3、监听服务器的地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}

		fmt.Println("start zinx server succ, ", s.Name, "succ, Listening...")

		var cid uint32
		cid = 0

		//4、阻塞的等待客户端连接，处理客户端连接业务（读写）
		for {
			//如果有客户端连接过来，阻塞会返回
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			//设置贵大连接个数的判断，如果超过最大连接，那么则关闭此新的的连接
			if s.ConnMgr.ConnTotal() >= utils.GlobalObject.MaxConn {
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

// Stop 停止服务器方法
func (s *Server) Stop() {
	fmt.Printf("[stop server] ip:%s, port:%d=============", s.IP, s.Port)
	s.ConnMgr.ClearConn()
}

// Serve 开启业务服务方法
func (s *Server) Serve() {
	s.Start()

	//阻塞状态
	select {}
}

func (s *Server) GetConnMgr() tiface.IConnManager {
	return s.ConnMgr
}

//路由功能：给当前的服务注册一个路由方法，供客户端的链接处理使用
func (s *Server) AddRouter(msgId uint32, router tiface.IRouter) {
	s.MsgHandler.AddRouter(msgId, router)
	fmt.Println("add Router success")
}
