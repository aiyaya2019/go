// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DeviceCustom is the golang structure of table device_custom for DAO operations like Where/Data.
type DeviceCustom struct {
	g.Meta       `orm:"table:device_custom, do:true"`
	Id           interface{} //
	IsShowDatum  interface{} // 是否显示议题：0不显示；1显示
	SkinColor    interface{} // 1浅蓝；2深蓝
	BgImg        interface{} // 背景皮肤
	Logo         interface{} // 空闲页logo
	PropagateImg interface{} // 空闲页宣传图
	Type         interface{} // 设备类型：1信息屏
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
}
