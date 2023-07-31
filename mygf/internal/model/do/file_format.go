// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FileFormat is the golang structure of table file_format for DAO operations like Where/Data.
type FileFormat struct {
	g.Meta         `orm:"table:file_format, do:true"`
	Id             interface{} // 逻辑ID
	SystemFileUuid interface{} // 文件uuid，system_file表的uuid
	Progress       interface{} // 进度
	Path           interface{} // 文件路径
	Status         interface{} // 转码状态，0未开始，1转换中，2成功，3失败，4待确认
	Type           interface{} // 类型，0=》无，1=》doc to pdf
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
