// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingVoteOption is the golang structure for table meeting_vote_option.
type MeetingVoteOption struct {
	Id              uint        `json:"id"              ` //
	Uuid            uint        `json:"uuid"            ` //
	PlatformUuid    uint64      `json:"platformUuid"    ` //
	MeetingVoteUuid uint64      `json:"meetingVoteUuid" ` //
	Sort            int         `json:"sort"            ` //
	VoteOptionName  string      `json:"voteOptionName"  ` //
	CreatedAt       *gtime.Time `json:"createdAt"       ` //
	UpdatedAt       *gtime.Time `json:"updatedAt"       ` //
}
