// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingVoteResult is the golang structure of table meeting_vote_result for DAO operations like Where/Data.
type MeetingVoteResult struct {
	g.Meta          `orm:"table:meeting_vote_result, do:true"`
	Id              interface{} //
	UserUuid        interface{} //
	MeetingVoteUuid interface{} //
	VoteOptionUuid  interface{} //
	VoteTime        *gtime.Time //
	Mark            interface{} //
	Score           interface{} //
	CreatedAt       *gtime.Time //
	UpdatedAt       *gtime.Time //
}
