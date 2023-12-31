// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mygf/internal/dao/internal"
)

// internalMeetingVoteDao is internal type for wrapping internal DAO implements.
type internalMeetingVoteDao = *internal.MeetingVoteDao

// meetingVoteDao is the data access object for table meeting_vote.
// You can define custom methods on it to extend its functionality as you wish.
type meetingVoteDao struct {
	internalMeetingVoteDao
}

var (
	// MeetingVote is globally public accessible object for table meeting_vote operations.
	MeetingVote = meetingVoteDao{
		internal.NewMeetingVoteDao(),
	}
)

// Fill with you ideas below.
