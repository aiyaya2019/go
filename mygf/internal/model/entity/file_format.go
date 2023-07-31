// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FileFormat is the golang structure for table file_format.
type FileFormat struct {
	Id             uint        `json:"id"             ` // 逻辑ID
	SystemFileUuid uint        `json:"systemFileUuid" ` // 文件uuid，system_file表的uuid
	Progress       string      `json:"progress"       ` // 进度
	Path           string      `json:"path"           ` // 文件路径
	Status         uint        `json:"status"         ` // 转码状态，0未开始，1转换中，2成功，3失败，4待确认
	Type           uint        `json:"type"           ` // 类型，0=》无，1=》doc to pdf
	CreatedAt      *gtime.Time `json:"createdAt"      ` // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      ` // 更新时间
}
