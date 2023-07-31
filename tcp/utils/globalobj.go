package utils

type GlobalObj struct {
	IPVersion        string //ip版本
	IP               string //ip
	Port             int    //端口
	Name             string //服务器名称
	Version          string //服务器版本号
	MaxPackageSize   uint32 //数据包的最大值
	WorkerPoolSize   uint32 //当前业务工作worker池的goroutine数量
	MaxWorkerTaskLen uint32 //允许用户最多开辟多少个worker（限定条件）
	MaxConn          int    //当前服务区主机允许的最大连接数
}

// GlobalObject 定义一个全局的对外GlobalObj
var GlobalObject *GlobalObj

func init() {
	GlobalObject = &GlobalObj{
		IPVersion:        "tcp",
		IP:               "127.0.0.1",
		Port:             8888,
		Name:             "tcp worker",
		Version:          "v1.0",
		MaxPackageSize:   4096,
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
		MaxConn:          3,
	}
}
