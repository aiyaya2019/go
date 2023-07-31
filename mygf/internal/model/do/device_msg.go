// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DeviceMsg is the golang structure of table device_msg for DAO operations like Where/Data.
type DeviceMsg struct {
	g.Meta    `orm:"table:device_msg, do:true"`
	Id        interface{} //
	DeviceId  interface{} // 设备id，device表主键id
	Msg       interface{} // 消息内容
	Time      interface{} // 消息播报时间/秒
	Type      interface{} // 设备类型：1信息屏
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
