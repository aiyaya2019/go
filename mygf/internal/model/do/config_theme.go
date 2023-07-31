// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ConfigTheme is the golang structure of table config_theme for DAO operations like Where/Data.
type ConfigTheme struct {
	g.Meta     `orm:"table:config_theme, do:true"`
	Id         interface{} // 逻辑ID
	SysName    interface{} // 系统名称
	ThemeName  interface{} // 主题名称
	Path       interface{} // 预览文件路径
	Logo       interface{} // LOGO
	Background interface{} // 背景图片
	Type       interface{} // 类型:1=>cms,2=>c#,3=>qt,4=>ios
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
}
