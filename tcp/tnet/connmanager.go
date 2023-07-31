package tnet

import (
	"errors"
	"fmt"
	"sync"
	"tcp/tiface"
)

type ConnManager struct {
	connections map[uint32]tiface.IConnection //管理的连接集合
	connLock    sync.RWMutex                  //保护连接集合的读写锁
}

//创建当前链接的方法
func NewConnManager() *ConnManager {
	return &ConnManager{
		connections: make(map[uint32]tiface.IConnection),
	}
}

//添加连接
func (connMgr *ConnManager) Add(conn tiface.IConnection) {
	//保护共享资源map，加写锁
	connMgr.connLock.Lock()

	//defer 是一个关键字，用于在函数退出之前延迟执行一段代码块。defer 语句可以被用于释放资源、关闭文件、解锁互斥锁等清理操作，或者用于记录日志、计算函数执行时间等其他需要在函数退出前执行的操作
	defer connMgr.connLock.Unlock()

	//将conn加入到ConnManager中
	connMgr.connections[conn.GetConnID()] = conn

	fmt.Println("connId=", conn.GetConnID(), " add to ConnManager successfully: conn num = ", connMgr.ConnTotal())
}

//删除连接
func (connMgr *ConnManager) Remove(conn tiface.IConnection) {
	//保护共享资源map，加写锁
	connMgr.connLock.Lock()

	defer connMgr.connLock.Unlock()

	//删除连接信息
	delete(connMgr.connections, conn.GetConnID())

	fmt.Println("connId=", conn.GetConnID(), " remove from ConnManager successfully: conn num = ", connMgr.ConnTotal())
}

//根据ConnID获取连接
func (connMgr *ConnManager) Get(connId uint32) (tiface.IConnection, error) {
	//保护共享资源map，加读锁
	connMgr.connLock.RLock()

	defer connMgr.connLock.RUnlock()

	if conn, ok := connMgr.connections[connId]; ok {
		return conn, nil
	} else {
		return nil, errors.New("connection not found")
	}
}

//得到当前连接总数
func (connMgr *ConnManager) ConnTotal() int {
	return len(connMgr.connections)
}

//清除并终止所有连接
func (connMgr *ConnManager) ClearConn() {
	//保护共享资源map，加写锁
	connMgr.connLock.Lock()

	defer connMgr.connLock.Unlock()

	//删除conn并停止conn的工作
	for connId, conn := range connMgr.connections {
		//停止
		conn.Stop()

		//删除
		delete(connMgr.connections, connId)
	}

	fmt.Println("clear all connections success. conn num = ", connMgr.ConnTotal())
}
