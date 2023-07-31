package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

func main() {
	//fmt.Println("Hello, World!")
	s := g.Server()
	fmt.Println(s)

	var aa int = 1
	fmt.Println(aa)

	s.BindHandler("/welcome", func(r *ghttp.Request) {
		glog.Info("哈喽")
		//glog.Panic("qwe")
		glog.Error("你异常啦！")
		r.Response.Write("哈喽世界！")
	})

	s.BindHandler("/panic", func(r *ghttp.Request) {
		glog.Panic("123")
	})

	// post请求
	s.BindHandler("/hello", func(r *ghttp.Request) {
		r.Response.Writeln("Hello World!")
	})
	s.SetPort(8002)
	s.Run()
}
