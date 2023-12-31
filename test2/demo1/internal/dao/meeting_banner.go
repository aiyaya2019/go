// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"demo1/internal/dao/internal"
)

// internalMeetingBannerDao is internal type for wrapping internal DAO implements.
type internalMeetingBannerDao = *internal.MeetingBannerDao

// meetingBannerDao is the data access object for table meeting_banner.
// You can define custom methods on it to extend its functionality as you wish.
type meetingBannerDao struct {
	internalMeetingBannerDao
}

var (
	// MeetingBanner is globally public accessible object for table meeting_banner operations.
	MeetingBanner = meetingBannerDao{
		internal.NewMeetingBannerDao(),
	}
)

// Fill with you ideas below.
