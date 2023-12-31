// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingVoteUser is the golang structure of table meeting_vote_user for DAO operations like Where/Data.
type MeetingVoteUser struct {
	g.Meta          `orm:"table:meeting_vote_user, do:true"`
	MeetingUuid     interface{} //
	UserUuid        interface{} //
	MeetingVoteUuid interface{} //
	SystemFileUuid  interface{} //
	CreatedAt       *gtime.Time //
	UpdatedAt       *gtime.Time //
}
