package controller

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	"mygf/api/v1"
)

var (
	Hello  = cHello{}
	Hello2 = cHello2{}
)

type cHello struct{}
type cHello2 struct{}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	//g.RequestFromCtx(ctx).Response.Writeln("Hello World!")

	g.RequestFromCtx(ctx).Response.WriteJson(g.Map{
		"id":   1,
		"name": "姓名",
	})
	return
}

func (c *cHello2) Hello2(ctx context.Context, req *v1.Hello2Req) (res *v1.Hello2Res, err error) {
	//g.RequestFromCtx(ctx).Response.Writeln("Hello World!")

	g.RequestFromCtx(ctx).Response.WriteJson(g.Map{
		"id":   1,
		"name": "姓名2",
	})
	return
}
