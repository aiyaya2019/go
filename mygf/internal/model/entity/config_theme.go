// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ConfigTheme is the golang structure for table config_theme.
type ConfigTheme struct {
	Id         uint        `json:"id"         ` // 逻辑ID
	SysName    string      `json:"sysName"    ` // 系统名称
	ThemeName  string      `json:"themeName"  ` // 主题名称
	Path       string      `json:"path"       ` // 预览文件路径
	Logo       string      `json:"logo"       ` // LOGO
	Background string      `json:"background" ` // 背景图片
	Type       int         `json:"type"       ` // 类型:1=>cms,2=>c#,3=>qt,4=>ios
	CreatedAt  *gtime.Time `json:"createdAt"  ` // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  ` // 更新时间
}
