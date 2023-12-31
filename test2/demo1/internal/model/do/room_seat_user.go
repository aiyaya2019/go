// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RoomSeatUser is the golang structure of table room_seat_user for DAO operations like Where/Data.
type RoomSeatUser struct {
	g.Meta           `orm:"table:room_seat_user, do:true"`
	RoomSeatId       interface{} //
	MeetingUuid      interface{} //
	UserUuid         interface{} //
	RoomSeatSchemeId interface{} //
	CreatedAt        *gtime.Time //
	UpdatedAt        *gtime.Time //
}
