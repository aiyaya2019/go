package client

import "fmt"

type sClient struct {
}

func (cli *sClient) Cmd601() (err error) {
	fmt.Println("client--cmd601")
	return
}
