// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingRecordContent is the golang structure for table meeting_record_content.
type MeetingRecordContent struct {
	Id          uint        `json:"id"          ` //
	MeetingUuid uint        `json:"meetingUuid" ` // 会议uuid，meeting表的uuid
	TerminalId  int         `json:"terminalId"  ` // 终端ID，terminal表的id
	HtmlContent string      `json:"htmlContent" ` // HTML会议记录内容
	WordContent string      `json:"wordContent" ` // Word会议记录内容
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 更新时间
}
