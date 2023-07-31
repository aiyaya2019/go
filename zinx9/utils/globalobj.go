package utils

import (
	"encoding/json"
	"io/ioutil"
	"zinx/ziface"
)

/*
存户一切有关Zinx框架的全局参数，供其他模块使用，一些参数也可以通过zinx.json由用户配置
*/

type GlobalObj struct {
	//server
	TcpServer ziface.IRequest //全局的server对象
	Host      string          //当前服务器主机监听的ip
	TcpPort   int             //当前服务器主机监听的端口号
	Name      string          //当前服务器的名称

	//zinx
	Version          string //当前zinx的版本号
	MaxConn          int    //当前服务区主机允许的最大连接数
	MaxPackageSize   uint32 //当前zinx框架数据包的最大值
	WorkerPoolSize   uint32 //当前业务工作worker池的goroutine数量
	MaxWorkerTaskLen uint32 //允许用户最多开辟多少个worker（限定条件）
}

/*
定义一个全局的对外Globalobj
*/
var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	//将json文件数据解析到struct中
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

func init() {
	//如果配置文件没有加载，默认的值
	GlobalObject = &GlobalObj{
		Name:             "ZinServerApp",
		Version:          "v0.9",
		TcpPort:          8999,
		Host:             "127.0.0.1",
		MaxConn:          3,
		MaxPackageSize:   4096,
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
	}

	//GlobalObject.Reload()
}
