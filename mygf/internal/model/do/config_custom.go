// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ConfigCustom is the golang structure of table config_custom for DAO operations like Where/Data.
type ConfigCustom struct {
	g.Meta         `orm:"table:config_custom, do:true"`
	Id             interface{} //
	SysName        interface{} // 名称
	Logo           interface{} // logo图标
	Background     interface{} // 背景图
	DefaultSetting interface{} // 0默认配置
	Type           interface{} // 类型:1=>cms,2=>c#,3=>qt,4=>android
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
