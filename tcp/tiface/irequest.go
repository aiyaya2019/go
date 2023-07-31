package tiface

//IRequest接口：实际上是把客户端请求的链接信息 和 请求的数据 包装到一个Request中
type IRequest interface {
	GetConnection() IConnection //获取请求连接信息
	GetMsg() []byte             //获取请求消息的数据
	GetMsgId() uint32           //请求的消息id
}
