// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Platform is the golang structure of table platform for DAO operations like Where/Data.
type Platform struct {
	g.Meta    `orm:"table:platform, do:true"`
	Id        interface{} //
	Uuid      interface{} // 平台uuid
	Name      interface{} // 平台名称
	IsLocal   interface{} // 是否为本地：0远端；1本地
	Sysip     interface{} // 平台ip
	Sysport   interface{} // 平台端口
	Sysid     interface{} // 平台自定义id，唯一
	Status    interface{} // 状态:0离线，1=>在线
	Enable    interface{} // 0禁用；1开启
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
