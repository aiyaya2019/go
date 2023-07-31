// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RoomSeatUser is the golang structure of table room_seat_user for DAO operations like Where/Data.
type RoomSeatUser struct {
	g.Meta           `orm:"table:room_seat_user, do:true"`
	RoomSeatId       interface{} // 座位id
	MeetingUuid      interface{} // 会议uuid，meeting表的uuid
	UserUuid         interface{} // 用户id，room_seat表的id
	RoomSeatSchemeId interface{} // 排位方案id，room_seat_scheme表的id
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
}
