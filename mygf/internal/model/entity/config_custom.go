// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ConfigCustom is the golang structure for table config_custom.
type ConfigCustom struct {
	Id             int         `json:"id"             ` //
	SysName        string      `json:"sysName"        ` // 名称
	Logo           string      `json:"logo"           ` // logo图标
	Background     string      `json:"background"     ` // 背景图
	DefaultSetting int         `json:"defaultSetting" ` // 0默认配置
	Type           int         `json:"type"           ` // 类型:1=>cms,2=>c#,3=>qt,4=>android
	CreatedAt      *gtime.Time `json:"createdAt"      ` // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      ` // 更新时间
}
