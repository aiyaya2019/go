package znet

import (
	"fmt"
	"strconv"
	"zinx/ziface"
)

/*
消息处理模块的实现
*/

type MsgHandle struct {
	//存放每个MsgId所对应的处理方法
	Apis map[uint32]ziface.IRouter
}

//初始化/创建MsgHandle方法
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
	}
}

//调度/执行对应的Router消息处理方法
func (mh *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	//1、从request中找到msgId
	handler, ok := mh.Apis[request.GetMsgId()]
	if !ok {
		fmt.Println("api msgIdo=", request.GetMsgId(), " is not found, need register")
	}

	//2、根据MsgId调度对应router业务即可
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)

}

//为消息添加具体的处理逻辑
func (mh *MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	//1、判断当前msg绑定的api处理方法是否已经存在
	if _, ok := mh.Apis[msgId]; ok {
		//id已经注册了
		panic("repeat api, msgId=" + strconv.Itoa(int(msgId)))
	}

	//2、添加msg与api的绑定关系
	mh.Apis[msgId] = router
	fmt.Println("add api msgId=", msgId, " success")

}
