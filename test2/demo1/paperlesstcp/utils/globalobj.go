package utils

type GlobalObj struct {
	Name      string //当前服务器的名称
	IpVersion string
	Ip        string //当前服务器主机监听的ip
	Port      int    //当前服务器主机监听的端口号
}

var GlobalObject *GlobalObj

func init() {
	GlobalObject = &GlobalObj{
		Name:      "paperless",
		IpVersion: "ipv4",
		Ip:        "127.0.0.1",
		Port:      8999,
	}
}
