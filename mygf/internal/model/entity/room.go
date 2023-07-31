// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Room is the golang structure for table room.
type Room struct {
	Id             uint        `json:"id"             ` // 逻辑id,会议室ID
	Name           string      `json:"name"           ` // 会议室名称
	ServerId       uint        `json:"serverId"       ` // 服务器ID，server表的id
	Roomlayout     string      `json:"roomlayout"     ` // 会议室布局json(单个)
	Status         uint        `json:"status"         ` // 状态:0=>未使用，1=>使用中
	Mark           string      `json:"mark"           ` // 备注
	Roombackground string      `json:"roombackground" ` // 会议室布局背景
	BgWidth        int         `json:"bgWidth"        ` // 背景图宽
	BgHeight       int         `json:"bgHeight"       ` // 背景图高
	TerminalConfig string      `json:"terminalConfig" ` // 终端配置json格式
	Transfer       int         `json:"transfer"       ` // 0不启用语音转写，1启用
	Electronic     int         `json:"electronic"     ` // 0不启用电子桌牌，1启用
	CreatedAt      *gtime.Time `json:"createdAt"      ` // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      ` // 更新时间
	TransferRoomId int         `json:"transferRoomId" ` //
}
