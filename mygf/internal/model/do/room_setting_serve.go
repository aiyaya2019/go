// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RoomSettingServe is the golang structure of table room_setting_serve for DAO operations like Where/Data.
type RoomSettingServe struct {
	g.Meta    `orm:"table:room_setting_serve, do:true"`
	Id        interface{} // 主键id
	RoomId    interface{} // 会议室id，room表的id
	Name      interface{} // 会议服务名字
	Src       interface{} // 图片存储路径
	IsSystem  interface{} // 是否是系统配置
	Sort      interface{} // 排序
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
