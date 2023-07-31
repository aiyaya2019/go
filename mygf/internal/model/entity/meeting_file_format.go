// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingFileFormat is the golang structure for table meeting_file_format.
type MeetingFileFormat struct {
	Id          uint        `json:"id"          ` // 逻辑ID
	MeetingUuid uint        `json:"meetingUuid" ` // 会议uuid，meeting表的uuid
	UserUuid    uint        `json:"userUuid"    ` // 用户uuid，user表的uuid
	Url         string      `json:"url"         ` // url路径
	Path        string      `json:"path"        ` // 文件路径
	Progress    string      `json:"progress"    ` // 进度
	Status      uint        `json:"status"      ` // 转码状态，0=》未开始，1=》转换中，2=》成功，3=》失败
	Time        *gtime.Time `json:"time"        ` // 时间
	Type        uint        `json:"type"        ` // 类型，0=》无，1=》html to image
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 更新时间
}
