// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RoomSeatUser is the golang structure for table room_seat_user.
type RoomSeatUser struct {
	RoomSeatId       int         `json:"roomSeatId"       ` //
	MeetingUuid      uint        `json:"meetingUuid"      ` //
	UserUuid         uint        `json:"userUuid"         ` //
	RoomSeatSchemeId int         `json:"roomSeatSchemeId" ` //
	CreatedAt        *gtime.Time `json:"createdAt"        ` //
	UpdatedAt        *gtime.Time `json:"updatedAt"        ` //
}
