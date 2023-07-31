// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingSummaryThird is the golang structure for table meeting_summary_third.
type MeetingSummaryThird struct {
	MeetingSummaryId int         `json:"meetingSummaryId" ` // 会议纪要id，meeting_summary的id
	PhoneticTranId   int         `json:"phoneticTranId"   ` // 语音转写那边的id，作对应关系
	Creator          string      `json:"creator"          ` // 创建者
	Modifier         string      `json:"modifier"         ` // 修改者
	CreatedAt        *gtime.Time `json:"createdAt"        ` // 创建时间
	UpdatedAt        *gtime.Time `json:"updatedAt"        ` // 更新时间
}
