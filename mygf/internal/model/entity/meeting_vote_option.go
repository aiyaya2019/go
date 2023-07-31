// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingVoteOption is the golang structure for table meeting_vote_option.
type MeetingVoteOption struct {
	Id              uint        `json:"id"              ` // 逻辑ID
	Uuid            uint        `json:"uuid"            ` // 分布式uuid
	PlatformUuid    uint64      `json:"platformUuid"    ` // 平台uuid
	MeetingVoteUuid uint64      `json:"meetingVoteUuid" ` // 投票uuid，meeting_vote表uuid
	Sort            int         `json:"sort"            ` // 权重
	VoteOptionName  string      `json:"voteOptionName"  ` // 投票选项名称
	CreatedAt       *gtime.Time `json:"createdAt"       ` // 创建时间
	UpdatedAt       *gtime.Time `json:"updatedAt"       ` // 更新时间
}
