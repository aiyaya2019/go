// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"demo1/internal/dao/internal"
)

// internalVideoFormatDao is internal type for wrapping internal DAO implements.
type internalVideoFormatDao = *internal.VideoFormatDao

// videoFormatDao is the data access object for table video_format.
// You can define custom methods on it to extend its functionality as you wish.
type videoFormatDao struct {
	internalVideoFormatDao
}

var (
	// VideoFormat is globally public accessible object for table video_format operations.
	VideoFormat = videoFormatDao{
		internal.NewVideoFormatDao(),
	}
)

// Fill with you ideas below.
