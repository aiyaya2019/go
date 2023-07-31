// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RoomSeatScheme is the golang structure of table room_seat_scheme for DAO operations like Where/Data.
type RoomSeatScheme struct {
	g.Meta      `orm:"table:room_seat_scheme, do:true"`
	Id          interface{} // id
	MeetingUuid interface{} // 会议uuid，meeting表的uuid
	Title       interface{} // 方案名称
	Type        interface{} // 0:默认方案,1:普通方案
	Status      interface{} // 0:未使用,1:使用中
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
