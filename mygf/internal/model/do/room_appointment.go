// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RoomAppointment is the golang structure of table room_appointment for DAO operations like Where/Data.
type RoomAppointment struct {
	g.Meta      `orm:"table:room_appointment, do:true"`
	Id          interface{} // 逻辑ID
	MeetingUuid interface{} // 会议uuid，meeting表的uuid
	UserUuid    interface{} // 预约用户的uuid，user表的uuid
	RoomId      interface{} // 会议室ID,room表id
	Name        interface{} // 预约人员的名字
	StartTime   *gtime.Time // 预约的开始时间
	EndTime     *gtime.Time // 预约的结束时间
	Status      interface{} // 状态，1正常，0停用
	Note        interface{} // 备注
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
