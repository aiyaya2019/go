// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mygf/internal/dao/internal"
)

// internalMeetingDatumUserDao is internal type for wrapping internal DAO implements.
type internalMeetingDatumUserDao = *internal.MeetingDatumUserDao

// meetingDatumUserDao is the data access object for table meeting_datum_user.
// You can define custom methods on it to extend its functionality as you wish.
type meetingDatumUserDao struct {
	internalMeetingDatumUserDao
}

var (
	// MeetingDatumUser is globally public accessible object for table meeting_datum_user operations.
	MeetingDatumUser = meetingDatumUserDao{
		internal.NewMeetingDatumUserDao(),
	}
)

// Fill with you ideas below.
