// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"demo1/internal/dao/internal"
)

// internalMeetingSummaryCountersignDao is internal type for wrapping internal DAO implements.
type internalMeetingSummaryCountersignDao = *internal.MeetingSummaryCountersignDao

// meetingSummaryCountersignDao is the data access object for table meeting_summary_countersign.
// You can define custom methods on it to extend its functionality as you wish.
type meetingSummaryCountersignDao struct {
	internalMeetingSummaryCountersignDao
}

var (
	// MeetingSummaryCountersign is globally public accessible object for table meeting_summary_countersign operations.
	MeetingSummaryCountersign = meetingSummaryCountersignDao{
		internal.NewMeetingSummaryCountersignDao(),
	}
)

// Fill with you ideas below.