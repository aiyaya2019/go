// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemFormat is the golang structure of table system_format for DAO operations like Where/Data.
type SystemFormat struct {
	g.Meta       `orm:"table:system_format, do:true"`
	Id           interface{} // 主键
	Type         interface{} // 文件类型1:议程2:议题;4:临时资料;20:会议纪要;99:u盘上传
	UploadSize   interface{} // 上传文件大小M选项json
	Size         interface{} // 文件大小 <=值
	SystemFormat interface{} // 系统允许配置json
	DefineFormat interface{} // 自定义配置。英文,分割
	Name         interface{} // 名称
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
}
