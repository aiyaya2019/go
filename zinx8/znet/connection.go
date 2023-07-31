package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

//连接模块
type Connection struct {
	//当前链接的socker TCP套接字
	Conn *net.TCPConn

	//链接的id
	ConnID uint32

	//当前的链接状态
	isClosed bool

	//告知当前链接已经退出的/停止channel(由reader告知writer退出)
	ExitChan chan bool

	//无缓冲管道，用于读、写goroutine之间的消息通讯
	msgChan chan []byte

	//消息的管理MsgId和对应的处理业务api关系
	MsgHandler ziface.IMsgHandle
}

//初始化链接模块的方法
func NewConnection(conn *net.TCPConn, connID uint32, msgHandler ziface.IMsgHandle) *Connection {
	c := &Connection{
		Conn:       conn,
		ConnID:     connID,
		MsgHandler: msgHandler,
		isClosed:   false,
		msgChan:    make(chan []byte),
		ExitChan:   make(chan bool, 1),
	}
	return c
}

//链接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("[Reader Goroutine is running]")

	defer fmt.Println("[reader is exit], connID = ", c.ConnID, "[Reader is exit], remote addr is", c.RemoteAddr().String())
	defer c.Stop()

	for {
		//读取客户端数据到buf中

		//创建一个拆包对象
		dp := NewDataPack()

		//读取客户端的msg head 二进制流 8个字节
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head error:", err)
			break
		}

		//拆包，得到MsgId和msgDataLen 放在msg消息中
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error:", err)
			break
		}

		//根据dataLen 再次读取data，放在msg.Data中
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error:", err)
				break
			}
		}
		msg.SetData(data)

		//得到当前conn数据的Request请求数据
		req := Request{
			conn: c,
			msg:  msg,
		}

		if utils.GlobalObject.WorkerPoolSize > 0 {
			//已经开启了工作池机制，将消息发送给worker工作池处理即可
			c.MsgHandler.SendMsgToTaskQueue(&req)

		} else {
			//从路由中找到注册绑定的Conn对应的router调用
			//根据绑定好的MsgId找到对应处理api业务 执行
			go c.MsgHandler.DoMsgHandler(&req)
		}
	}

}

//写消息goroutine，专门发送给客户端消息的模块
func (c *Connection) StartWriter() {
	fmt.Println("[writer goroutine is running]")

	defer fmt.Println(c.RemoteAddr().String(), " [conn writer exit]")

	//不断的阻塞的等待channel的消息，进行写给客户端
	for {
		select {
		case data := <-c.msgChan:
			//有数据要写给客户端
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("send data error:", err)
				return
			}
		case <-c.ExitChan:
			//代表Reader一斤刚退出，此时Writer也要退出
			return
		}
	}

}

//启动连接，让当前的连接准备开始工作
func (c *Connection) Start() {
	fmt.Println("Conn Start()...ConnID=", c.ConnID)

	//启动从当前链接读数据的业务
	go c.StartReader()

	//启动从当前链接写数据的业务
	go c.StartWriter()

}

//停止连接，结束当前连接的工作
func (c *Connection) Stop() {
	fmt.Println("Conn Stop().. ConnID=", c.ConnID)

	//如果当前链接已关闭
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	//关闭socket链接
	c.Conn.Close()

	//告知writer关闭
	c.ExitChan <- true

	//回收资源
	close(c.ExitChan)
	close(c.msgChan)
}

//获取当前连接的绑定socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

//获取当前连接模块的连接id
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

//获取远程客户端的tcp状态 ip port
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()

}

//发送数据，将要发送给客户端的数据，先进行封包，再发送
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("connection closed when send msg")
	}

	//将data进行封包 MsgDataLen|MsgId|Data
	dp := NewDataPack()
	binaryMsg, err := dp.Pack(NewMsgPackage(msgId, data))
	if err != nil {
		fmt.Println("Pack error msg id = ", msgId)
		return errors.New("pack error msg")
	}

	//将数据发送给客户端
	c.msgChan <- binaryMsg
	return nil
}
