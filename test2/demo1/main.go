package main

import (
	_ "demo1/internal/packed"
	"demo1/paperlesstcp/pnet"
)

func main() {
	s := pnet.NewServer()
	s.Server()

	////启动HTTP服务
	//cmd.Main.Run(gctx.New())
	//
	//fmt.Println("1-----------------")
	////service.Client()

}
