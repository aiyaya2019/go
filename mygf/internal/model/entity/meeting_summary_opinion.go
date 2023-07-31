// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingSummaryOpinion is the golang structure for table meeting_summary_opinion.
type MeetingSummaryOpinion struct {
	Id               uint        `json:"id"               ` //
	MeetingUuid      uint        `json:"meetingUuid"      ` // 会议uuid，meeting表的uuid
	UserUuid         uint        `json:"userUuid"         ` // 用户uuid，user表的uuid
	MeetingSummaryId uint        `json:"meetingSummaryId" ` // 纪要ID，meeting_summary表的id
	SendTime         *gtime.Time `json:"sendTime"         ` // 发送时间
	Content          string      `json:"content"          ` // 纪要内容
	Page             int         `json:"page"             ` // 在纪要文件中第几页提意见
	CreatedAt        *gtime.Time `json:"createdAt"        ` // 创建时间
	UpdatedAt        *gtime.Time `json:"updatedAt"        ` // 更新时间
}
