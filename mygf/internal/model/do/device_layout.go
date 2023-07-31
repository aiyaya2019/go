// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DeviceLayout is the golang structure of table device_layout for DAO operations like Where/Data.
type DeviceLayout struct {
	g.Meta    `orm:"table:device_layout, do:true"`
	Id        interface{} //
	DeviceId  interface{} // 设备id，device表主键id
	RoomId    interface{} // 会议室id
	StartTime *gtime.Time // 开始推送时间
	EndTime   *gtime.Time // 结束推送时间
	PushType  interface{} // 推送类型：0单次；1每天；2每周一；3每周二；4每周三；5每周四；6每周五；7每周六；8每周日
	PushState interface{} // 允许推送的设备状态：0空闲；1使用中；2空闲和使用中
	Status    interface{} // 0不推送；1推送
	Layout    interface{} // 空闲页布局
	Version   interface{} // 版本
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
