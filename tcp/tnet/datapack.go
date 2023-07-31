package tnet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"strconv"
	"tcp/tiface"
	"tcp/utils"
)

//封包、拆包的具体模块
type DataPack struct {
}

//拆包封包实例的一个初始化方法
func NewDataPack() *DataPack {
	return &DataPack{}
}

//获取包头的长度
func (dp *DataPack) GetHeadLen() uint32 {
	//DataLen uint32 (4字节) + id uint32 (4字节)
	return 8
}

//封包方法。datalen|msgId|data
func (dp *DataPack) Pack(msg tiface.IMessage) ([]byte, error) {
	//创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	//将msgLen写进databuff中
	//binary.Write函数用于将数据以二进制形式写入到一个实现了io.Writer接口的目标中，如文件、网络连接或内存缓冲区
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}

	//将MsgId写进databuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	//将Msg写进databuff中
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsg()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

//拆包方法(将包的head信息读出来，之后再根据head信息里的data的长度再进行一次读)
func (dp *DataPack) Unpack(binaryData []byte) (tiface.IMessage, error) {
	//bytes.NewReader 函数用于创建一个实现了 io.Reader 接口的读取器，该读取器从给定的字节切片中读取数据
	dataBuff := bytes.NewReader(binaryData)

	//只解压head信息，得到MsgLen和MsgId
	msg := &Message{}

	//读MsgLen
	//binary.Read 函数用于从实现了 io.Reader 接口的源中读取二进制数据，并将其解码为指定的数据类型
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.MsgLen); err != nil {
		return nil, err
	}

	//读MsgId
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	//判断MsgLen是否已经超出我们允许的最大包长度
	if utils.GlobalObject.MaxPackageSize > 0 && msg.MsgLen > utils.GlobalObject.MaxPackageSize {
		return nil, errors.New("too large msg data, len:" + strconv.Itoa(int(msg.MsgLen)))
	}
	return msg, nil
}
