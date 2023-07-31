package znet

import (
	"fmt"
	"strconv"
	"zinx/utils"
	"zinx/ziface"
)

/*
消息处理模块的实现
*/

type MsgHandle struct {
	//存放每个MsgId所对应的处理方法
	Apis map[uint32]ziface.IRouter

	//负责worker取任务的消息队列
	TaskQueue []chan ziface.IRequest
	//业务工作worker池的worker数量
	WorkerPoolSize uint32
}

//初始化/创建MsgHandle方法
func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis:           make(map[uint32]ziface.IRouter),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize, //从全局配置中获取
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen),
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

//启动一个worker工作池(开启工作池的动作只能发生一次，一个zinx框架只能有一个worker工作池)
func (mh *MsgHandle) StartWorkerPool() {
	//根据WorkerPoolSize分别开启worker，每个worker用一个go来承载
	for i := 0; i < int(mh.WorkerPoolSize); i++ {
		//一个worker被启动
		//1、当前的worker对应的channel消息队列， 开辟空间。第0个worker就用第0个channel...
		mh.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerTaskLen)

		//2、启动当前的worker，阻塞等待消息聪channel传递进来
		go mh.StartOneWorker(i, mh.TaskQueue[i])
	}
}

//启动一个worker工作流程
func (mh *MsgHandle) StartOneWorker(workerId int, taskQueue chan ziface.IRequest) {
	fmt.Println("workerId = ", workerId, " is started...")

	//不断的阻塞等待对应消息队列的消息
	for {
		select {
		//如果有消息过来，出列的就是一个客户端的request，执行当前request所绑定的业务
		case request := <-taskQueue:
			mh.DoMsgHandler(request)
		}
	}
}

//将消息交给TaskQueue，由worker进行处理
func (mh *MsgHandle) SendMsgToTaskQueue(request ziface.IRequest) {
	//1、将消息平均分配给不同的worker
	//根据客户端简历的ConnId来进行分配（有requestId可用requestId）
	workerId := request.GetConnection().GetConnID() % mh.WorkerPoolSize
	fmt.Println("add connId=", request.GetConnection().GetConnID(), " request MsgId=", request.GetMsgId(), " to workerId=", workerId)

	//2、将消息发送给对应的worker的TaskQueue即可
	mh.TaskQueue[workerId] <- request

}
