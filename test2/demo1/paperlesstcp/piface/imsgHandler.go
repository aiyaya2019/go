package piface

type IMsgHandler interface {
	//启动worker工作池
	StartWorkerPool()
}
