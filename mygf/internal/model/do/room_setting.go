// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RoomSetting is the golang structure of table room_setting for DAO operations like Where/Data.
type RoomSetting struct {
	g.Meta    `orm:"table:room_setting, do:true"`
	Id        interface{} // 逻辑id
	RoomId    interface{} // 会议室id,room表的id
	Type      interface{} // 功能类型
	Content   interface{} // 设置的json内容
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
