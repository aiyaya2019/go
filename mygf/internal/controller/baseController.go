package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	BController = &bController{}
)

type bController struct{}

type ReturnData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (b *bController) ReturnJson(ctx context.Context, returnData ReturnData) {
	g.RequestFromCtx(ctx).Response.WriteJson(returnData)
}
