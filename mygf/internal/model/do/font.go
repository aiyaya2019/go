// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Font is the golang structure of table font for DAO operations like Where/Data.
type Font struct {
	g.Meta     `orm:"table:font, do:true"`
	Id         interface{} // 逻辑ID
	Name       interface{} // 字体名称
	Path       interface{} // 文件路径
	UploadTime *gtime.Time // 上传时间
	Status     interface{} // 状态;1:启用,0:禁用
	IsSysFont  interface{} // 是否系统自带字体;1:是,0:否
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
}
