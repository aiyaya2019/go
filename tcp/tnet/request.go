package tnet

import (
	"tcp/tiface"
)

type Request struct {
	conn tiface.IConnection //已经和客户建立好的连接
	msg  tiface.IMessage    //客户端请求的数据
}

//获取请求连接信息
func (r *Request) GetConnection() tiface.IConnection {
	return r.conn
}

//获取请求消息的数据
func (r *Request) GetMsg() []byte {
	return r.msg.GetMsg()
}

//请求的消息id
func (r *Request) GetMsgId() uint32 {
	return r.msg.GetMsgId()
}
