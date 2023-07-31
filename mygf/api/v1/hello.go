package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HelloReq struct {
	g.Meta `path:"/hello" tags:"Hello" method:"get" summary:"You first hello api"`
}

type HelloRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type Hello2Req struct {
	g.Meta `path:"/hello2" tags:"Hello" method:"get" summary:"You sec hello api"`
}

type Hello2Res struct {
	g.Meta `mime:"text/html" example:"string"`
}
