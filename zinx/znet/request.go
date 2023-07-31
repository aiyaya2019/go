package znet

import "zinx/ziface"

type Request struct {
	//已经和客户端简历好的链接
	conn ziface.IConnectioin

	//客户端请求的数据
	data []byte
}

//得到当前链接
func (r *Request) GetConnection() ziface.IConnectioin {
	return r.conn
}

//得到请求的消息数据
func (r *Request) GetData() []byte {
	return r.data
}
