// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// VideoFormat is the golang structure for table video_format.
type VideoFormat struct {
	Id             uint        `json:"id"             ` // 逻辑ID
	SystemFileUuid uint        `json:"systemFileUuid" ` // 视频文件uuid，system_file表的uuid
	Progress       string      `json:"progress"       ` // 进度
	Status         uint        `json:"status"         ` // 转码状态，0=》未开始，1=》转换中，2=》成功，3=》失败，4 =》待确认
	Path           string      `json:"path"           ` // 文件路径
	Type           uint        `json:"type"           ` // 0=》 转MP4，aac，1=》全部 copy，2=》视频copy，音频转aac
	Force          int         `json:"force"          ` // 强制转码0:否1:是
	Remain         *gtime.Time `json:"remain"         ` // 预计剩余时长
	CreatedAt      *gtime.Time `json:"createdAt"      ` // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      ` // 更新时间
}
