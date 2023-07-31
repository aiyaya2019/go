// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// HtmlFormat is the golang structure for table html_format.
type HtmlFormat struct {
	Id          uint        `json:"id"          ` // 逻辑ID
	MeetingUuid uint        `json:"meetingUuid" ` // 会议uuid，meeting表的uuid
	MultiUuid   uint        `json:"multiUuid"   ` // 用户/标语uuid，user表的uuid/meeting_banner的uuid
	Url         string      `json:"url"         ` // url路径
	Path        string      `json:"path"        ` // 文件路径
	Progress    string      `json:"progress"    ` // 进度
	Status      uint        `json:"status"      ` // 转码状态，0=》未开始，1=》转换中，2=》成功，3=》失败
	Type        uint        `json:"type"        ` // 类型，0=》无，1=》html to image，2=》标语
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 更新时间
}
