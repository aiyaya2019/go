// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemBackups is the golang structure of table system_backups for DAO operations like Where/Data.
type SystemBackups struct {
	g.Meta     `orm:"table:system_backups, do:true"`
	Id         interface{} // 主键
	Name       interface{} // 备份名字
	CreateTime *gtime.Time // 创建时间
	Time       *gtime.Time // 备份花费时间
	FileSize   interface{} // 文件大小
	Type       interface{} // 1=>系统配置 2=>服务器文件
	Status     interface{} // 1=>备份中 2=>成功 3=>失败
	Path       interface{} // 文件路径
	SqlPath    interface{} // 数据库文件存放路径
	Mark       interface{} // 备注
	CurrentUse interface{} // 0=>默认不使用1=>升级中 2=>成功使用 3=>升级失败
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
}
