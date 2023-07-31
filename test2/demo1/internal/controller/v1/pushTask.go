package v1

import (
	"context"
	v1 "demo1/api/v1"
	"fmt"
)

var (
	PushTask = cPushTask{}
)

type cPushTask struct {
}

func (c *cPushTask) PushTask(ctx context.Context, req *v1.PushTaskReq) (res *v1.PushTaskRes, err error) {
	fmt.Println("controller--v1--PushTask")
	res = &v1.PushTaskRes{
		Id:          req.Id,
		Inst:        req.Inst,
		ServerId:    req.ServerId,
		MeetingUuid: req.MeetingUuid,
		Content:     req.Content,
		Mark:        req.Mark,
	}
	return
}
