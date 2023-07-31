// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingPcMenu is the golang structure for table meeting_pc_menu.
type MeetingPcMenu struct {
	Id        uint        `json:"id"        ` // 逻辑id
	Pid       uint        `json:"pid"       ` // 父级id
	Name      string      `json:"name"      ` // 菜单名称
	Type      string      `json:"type"      ` // PC端类型判断
	Status    uint        `json:"status"    ` // 状态
	Sort      uint        `json:"sort"      ` // 排序
	Have      int         `json:"have"      ` // 菜单模式：1卡片模式；2简洁模式；3导航模式；4经典模式
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
}
