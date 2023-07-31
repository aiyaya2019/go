// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Device is the golang structure of table device for DAO operations like Where/Data.
type Device struct {
	g.Meta    `orm:"table:device, do:true"`
	Id        interface{} //
	DeviceIp  interface{} // 设备ip
	DeviceMac interface{} // 设备mac
	RoomId    interface{} // 会议室id
	Status    interface{} // 设备状态:0空闲；1使用中；2设备离线
	IsShowing interface{} // 当前是否显示图文：0否；1是
	IsPush    interface{} // 是否允许推送：0否；1是
	Type      interface{} // 设备类型：1信息屏
	Deleted   interface{} // 删除：0未删除；1删除
	Ip        interface{} // 服务器ip
	Port      interface{} // 服务器端口
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
