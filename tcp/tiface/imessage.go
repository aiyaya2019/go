package tiface

//将请求的消息封装到一个Message中，定义抽象的接口
type IMessage interface {
	//获取消息的id
	GetMsgId() uint32
	//获取消息的长度
	GetMsgLen() uint32
	//获取消息内容
	GetMsg() []byte

	//获取消息的id
	SetMsgId(uint32)
	//获取消息的长度
	SetMsgLen(uint32)
	//获取消息内容
	SetMsg([]byte)
}
