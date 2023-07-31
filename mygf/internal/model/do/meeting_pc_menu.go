// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingPcMenu is the golang structure of table meeting_pc_menu for DAO operations like Where/Data.
type MeetingPcMenu struct {
	g.Meta    `orm:"table:meeting_pc_menu, do:true"`
	Id        interface{} // 逻辑id
	Pid       interface{} // 父级id
	Name      interface{} // 菜单名称
	Type      interface{} // PC端类型判断
	Status    interface{} // 状态
	Sort      interface{} // 排序
	Have      interface{} // 菜单模式：1卡片模式；2简洁模式；3导航模式；4经典模式
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
