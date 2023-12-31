package tiface

// IDataPack 封包、拆包模块。直接面向tcp链接中的数据流，用于处理tcp沾包问题
type IDataPack interface {
	//获取包头的长度
	GetHeadLen() uint32
	//封包方法
	Pack(msg IMessage) ([]byte, error)
	//拆包方法
	Unpack([]byte) (IMessage, error)
}
