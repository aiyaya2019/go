// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SystemFormat is the golang structure for table system_format.
type SystemFormat struct {
	Id           int         `json:"id"           ` // 主键
	Type         int         `json:"type"         ` // 文件类型1:议程2:议题;4:临时资料;20:会议纪要;99:u盘上传
	UploadSize   string      `json:"uploadSize"   ` // 上传文件大小M选项json
	Size         int         `json:"size"         ` // 文件大小 <=值
	SystemFormat string      `json:"systemFormat" ` // 系统允许配置json
	DefineFormat string      `json:"defineFormat" ` // 自定义配置。英文,分割
	Name         string      `json:"name"         ` // 名称
	CreatedAt    *gtime.Time `json:"createdAt"    ` // 创建时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    ` // 更新时间
}
