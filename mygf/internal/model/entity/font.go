// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Font is the golang structure for table font.
type Font struct {
	Id         uint        `json:"id"         ` // 逻辑ID
	Name       string      `json:"name"       ` // 字体名称
	Path       string      `json:"path"       ` // 文件路径
	UploadTime *gtime.Time `json:"uploadTime" ` // 上传时间
	Status     int         `json:"status"     ` // 状态;1:启用,0:禁用
	IsSysFont  int         `json:"isSysFont"  ` // 是否系统自带字体;1:是,0:否
	CreatedAt  *gtime.Time `json:"createdAt"  ` // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  ` // 更新时间
}
