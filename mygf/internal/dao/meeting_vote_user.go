// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mygf/internal/dao/internal"
)

// internalMeetingVoteUserDao is internal type for wrapping internal DAO implements.
type internalMeetingVoteUserDao = *internal.MeetingVoteUserDao

// meetingVoteUserDao is the data access object for table meeting_vote_user.
// You can define custom methods on it to extend its functionality as you wish.
type meetingVoteUserDao struct {
	internalMeetingVoteUserDao
}

var (
	// MeetingVoteUser is globally public accessible object for table meeting_vote_user operations.
	MeetingVoteUser = meetingVoteUserDao{
		internal.NewMeetingVoteUserDao(),
	}
)

// Fill with you ideas below.
