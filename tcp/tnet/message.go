package tnet

type Message struct {
	Id     uint32 //消息id
	MsgLen uint32 //消息长度
	Msg    []byte //消息内容
}

//创建一个message消息包
func NewMsgPackage(id uint32, data []byte) *Message {
	return &Message{
		Id:     id,
		MsgLen: uint32(len(data)),
		Msg:    data,
	}
}

//获取消息的id
func (m *Message) GetMsgId() uint32 {
	return m.Id
}

//获取消息的长度
func (m *Message) GetMsgLen() uint32 {
	return m.MsgLen
}

//获取消息内容
func (m *Message) GetMsg() []byte {
	return m.Msg
}

//获取消息的id
func (m *Message) SetMsgId(id uint32) {
	m.Id = id
}

//获取消息的长度
func (m *Message) SetMsgLen(len uint32) {
	m.Id = len
}

//获取消息内容
func (m *Message) SetMsg(msg []byte) {
	m.Msg = msg
}
