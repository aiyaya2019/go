// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RoomSeatUser is the golang structure for table room_seat_user.
type RoomSeatUser struct {
	RoomSeatId       int         `json:"roomSeatId"       ` // 座位id
	MeetingUuid      uint        `json:"meetingUuid"      ` // 会议uuid，meeting表的uuid
	UserUuid         uint        `json:"userUuid"         ` // 用户id，room_seat表的id
	RoomSeatSchemeId int         `json:"roomSeatSchemeId" ` // 排位方案id，room_seat_scheme表的id
	CreatedAt        *gtime.Time `json:"createdAt"        ` // 创建时间
	UpdatedAt        *gtime.Time `json:"updatedAt"        ` // 更新时间
}
