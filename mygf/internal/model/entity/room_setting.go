// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RoomSetting is the golang structure for table room_setting.
type RoomSetting struct {
	Id        int         `json:"id"        ` // 逻辑id
	RoomId    int         `json:"roomId"    ` // 会议室id,room表的id
	Type      string      `json:"type"      ` // 功能类型
	Content   string      `json:"content"   ` // 设置的json内容
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
}
