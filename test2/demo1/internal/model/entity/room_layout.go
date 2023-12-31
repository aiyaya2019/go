// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RoomLayout is the golang structure for table room_layout.
type RoomLayout struct {
	Id        uint        `json:"id"        ` //
	RoomId    int         `json:"roomId"    ` // 会议室id
	StartTime *gtime.Time `json:"startTime" ` // 开始推送时间
	EndTime   *gtime.Time `json:"endTime"   ` // 结束推送时间
	PushType  string      `json:"pushType"  ` // 推送类型：0单次；1每天；2每周一；3每周二；4每周三；5每周四；6每周五；7每周六；8每周日
	PushState int         `json:"pushState" ` // 允许推送的设备状态：0空闲；1使用中；2空闲和使用中
	Status    int         `json:"status"    ` // 0不推送；1推送
	Layout    string      `json:"layout"    ` // 空闲页布局
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
}
