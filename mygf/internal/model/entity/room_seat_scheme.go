// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RoomSeatScheme is the golang structure for table room_seat_scheme.
type RoomSeatScheme struct {
	Id          int         `json:"id"          ` // id
	MeetingUuid uint        `json:"meetingUuid" ` // 会议uuid，meeting表的uuid
	Title       string      `json:"title"       ` // 方案名称
	Type        int         `json:"type"        ` // 0:默认方案,1:普通方案
	Status      int         `json:"status"      ` // 0:未使用,1:使用中
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 更新时间
}
