// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingVoteResult is the golang structure for table meeting_vote_result.
type MeetingVoteResult struct {
	Id              uint        `json:"id"              ` // 逻辑ID
	UserUuid        uint        `json:"userUuid"        ` // 用户uuid，user表的id
	MeetingVoteUuid uint        `json:"meetingVoteUuid" ` // 投票uuid，meeting_vote表uuid
	VoteOptionUuid  uint        `json:"voteOptionUuid"  ` // 选项uuid，vote_option表的uuid
	VoteTime        *gtime.Time `json:"voteTime"        ` // 投票时间
	Mark            string      `json:"mark"            ` // 备注
	Score           int         `json:"score"           ` // 分值
	CreatedAt       *gtime.Time `json:"createdAt"       ` // 创建时间
	UpdatedAt       *gtime.Time `json:"updatedAt"       ` // 更新时间
}
