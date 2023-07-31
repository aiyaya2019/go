// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingStatistics is the golang structure for table meeting_statistics.
type MeetingStatistics struct {
	Id               int         `json:"id"               ` // 主键
	MeetingUuid      int64       `json:"meetingUuid"      ` // 关联会议uuid
	MeetingUserCount int         `json:"meetingUserCount" ` // 参会人数
	MeetingDuration  int         `json:"meetingDuration"  ` // 会议时长
	SavePaperCount   int         `json:"savePaperCount"   ` // 节约纸张数
	MeetingTime      *gtime.Time `json:"meetingTime"      ` // 会议时间
	StartTime        *gtime.Time `json:"startTime"        ` // 会议开始时间
	EndTime          *gtime.Time `json:"endTime"          ` // 会议结束时间
	CreatedAt        *gtime.Time `json:"createdAt"        ` // 创建时间
	UpdatedAt        *gtime.Time `json:"updatedAt"        ` // 更新时间
}
