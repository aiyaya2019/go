// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mygf/internal/dao/internal"
)

// internalMeetingServiceDao is internal type for wrapping internal DAO implements.
type internalMeetingServiceDao = *internal.MeetingServiceDao

// meetingServiceDao is the data access object for table meeting_service.
// You can define custom methods on it to extend its functionality as you wish.
type meetingServiceDao struct {
	internalMeetingServiceDao
}

var (
	// MeetingService is globally public accessible object for table meeting_service operations.
	MeetingService = meetingServiceDao{
		internal.NewMeetingServiceDao(),
	}
)

// Fill with you ideas below.