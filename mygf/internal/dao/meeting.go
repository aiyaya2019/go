// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mygf/internal/dao/internal"
)

// internalMeetingDao is internal type for wrapping internal DAO implements.
type internalMeetingDao = *internal.MeetingDao

// meetingDao is the data access object for table meeting.
// You can define custom methods on it to extend its functionality as you wish.
type meetingDao struct {
	internalMeetingDao
}

var (
	// Meeting is globally public accessible object for table meeting operations.
	Meeting = meetingDao{
		internal.NewMeetingDao(),
	}
)

// Fill with you ideas below.
