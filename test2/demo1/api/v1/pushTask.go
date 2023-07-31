package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type PushTaskReq struct {
	g.Meta      `path:"v1/pushTask" tags:"pushTask" method:"get" summary:"task push"`
	Id          int    `json:"id"`1
	Inst        int    `json:"inst"`
	ServerId    int    `json:"server_id"`
	MeetingUuid int    `json:"meeting_uuid"`
	Content     string `json:"content"`
	Mark        string `json:"mark"`
}

type PushTaskRes struct {
	g.Meta      `mime:"text/html" example:"string"`
	Id          int    `json:"id"`
	Inst        int    `json:"inst"`
	ServerId    int    `json:"server_id"`
	MeetingUuid int    `json:"meeting_uuid"`
	Content     string `json:"content"`
	Mark        string `json:"mark"`
}
