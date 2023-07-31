// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MicrophoneServerTerminal is the golang structure for table microphone_server_terminal.
type MicrophoneServerTerminal struct {
	Id                 uint        `json:"id"                 ` // 逻辑id
	MicId              int         `json:"micId"              ` // 外部传入
	Type               uint        `json:"type"               ` // 话筒类型 1.主席，0.普通
	Status             int         `json:"status"             ` // 话筒状态 1.打开，0.关闭
	MicrophoneServerId uint        `json:"microphoneServerId" ` // 连接会议主机id，microphone_server表的id
	Mark               string      `json:"mark"               ` //
	Mac                string      `json:"mac"                ` // 安卓话筒主机有带mac地址，在客户端列表关联话筒
	ScreenIp           string      `json:"screenIp"           ` // 安卓话筒背屏ip
	CreatedAt          *gtime.Time `json:"createdAt"          ` // 创建时间
	UpdatedAt          *gtime.Time `json:"updatedAt"          ` // 更新时间
}
