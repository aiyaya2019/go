package service

import "fmt"

type IClient interface {
	Cmd601() (err error)
}

var localClient IClient

func Client() IClient {
	fmt.Println("service-client.go")
	return localClient
}
