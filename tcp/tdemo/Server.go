package main

import "tcp/tnet"

type PingRouter struct {
	tnet.BaseRouter
}

func main() {
	s := tnet.NewServer()

	s.AddRouter(1, &PingRouter{})
	s.AddRouter(2, &PingRouter{})

	s.Serve()
}
