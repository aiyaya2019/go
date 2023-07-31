// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"demo1/internal/dao/internal"
)

// internalUserFingerPrintDao is internal type for wrapping internal DAO implements.
type internalUserFingerPrintDao = *internal.UserFingerPrintDao

// userFingerPrintDao is the data access object for table user_finger_print.
// You can define custom methods on it to extend its functionality as you wish.
type userFingerPrintDao struct {
	internalUserFingerPrintDao
}

var (
	// UserFingerPrint is globally public accessible object for table user_finger_print operations.
	UserFingerPrint = userFingerPrintDao{
		internal.NewUserFingerPrintDao(),
	}
)

// Fill with you ideas below.
