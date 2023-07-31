// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mygf/internal/dao/internal"
)

// internalMeetingRecordContentDao is internal type for wrapping internal DAO implements.
type internalMeetingRecordContentDao = *internal.MeetingRecordContentDao

// meetingRecordContentDao is the data access object for table meeting_record_content.
// You can define custom methods on it to extend its functionality as you wish.
type meetingRecordContentDao struct {
	internalMeetingRecordContentDao
}

var (
	// MeetingRecordContent is globally public accessible object for table meeting_record_content operations.
	MeetingRecordContent = meetingRecordContentDao{
		internal.NewMeetingRecordContentDao(),
	}
)

// Fill with you ideas below.
