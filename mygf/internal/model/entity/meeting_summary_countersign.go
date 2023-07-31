// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingSummaryCountersign is the golang structure for table meeting_summary_countersign.
type MeetingSummaryCountersign struct {
	MeetingUuid        uint        `json:"meetingUuid"        ` // 会议uuid，meeting表的uuid
	MeetingSummaryId   uint        `json:"meetingSummaryId"   ` // 会议纪要ID，meeting_summary表的id
	UserSystemFileUuid uint        `json:"userSystemFileUuid" ` // 会签文件uuid，user_file的system_file_uuid
	UserUuid           uint        `json:"userUuid"           ` // 用户uuid，user表的uuid
	Status             uint        `json:"status"             ` // 会签状态
	CreatedAt          *gtime.Time `json:"createdAt"          ` // 创建时间
	UpdatedAt          *gtime.Time `json:"updatedAt"          ` // 更新时间
}
