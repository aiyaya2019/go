// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingVoteUser is the golang structure for table meeting_vote_user.
type MeetingVoteUser struct {
	MeetingUuid     uint        `json:"meetingUuid"     ` // 会议uuid，meeting表的uuid
	UserUuid        uint        `json:"userUuid"        ` // 用户uuid，user表的uuid
	MeetingVoteUuid uint        `json:"meetingVoteUuid" ` // 投票uuid，meeting_vote表的uuid
	SystemFileUuid  uint        `json:"systemFileUuid"  ` // 文件uuid，system_file表的uuid
	CreatedAt       *gtime.Time `json:"createdAt"       ` // 创建时间
	UpdatedAt       *gtime.Time `json:"updatedAt"       ` // 更新时间
}
