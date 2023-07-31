// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RoomSettingServe is the golang structure for table room_setting_serve.
type RoomSettingServe struct {
	Id        int         `json:"id"        ` // 主键id
	RoomId    int         `json:"roomId"    ` // 会议室id，room表的id
	Name      string      `json:"name"      ` // 会议服务名字
	Src       string      `json:"src"       ` // 图片存储路径
	IsSystem  int         `json:"isSystem"  ` // 是否是系统配置
	Sort      int         `json:"sort"      ` // 排序
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
}
