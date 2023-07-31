// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemBackups is the golang structure for table system_backups.
type SystemBackups struct {
	Id         int         `json:"id"         ` // 主键
	Name       string      `json:"name"       ` // 备份名字
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	Time       *gtime.Time `json:"time"       ` // 备份花费时间
	FileSize   float64     `json:"fileSize"   ` // 文件大小
	Type       int         `json:"type"       ` // 1=>系统配置 2=>服务器文件
	Status     int         `json:"status"     ` // 1=>备份中 2=>成功 3=>失败
	Path       string      `json:"path"       ` // 文件路径
	SqlPath    string      `json:"sqlPath"    ` // 数据库文件存放路径
	Mark       string      `json:"mark"       ` // 备注
	CurrentUse int         `json:"currentUse" ` // 0=>默认不使用1=>升级中 2=>成功使用 3=>升级失败
	CreatedAt  *gtime.Time `json:"createdAt"  ` // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  ` // 更新时间
}
