package piface

type IServer interface {
	//运行服务器
	Server()
	//启动服务器
	Start()
	//停止服务器
	Stop()
}
