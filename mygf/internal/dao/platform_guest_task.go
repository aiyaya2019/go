// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mygf/internal/dao/internal"
)

// internalPlatformGuestTaskDao is internal type for wrapping internal DAO implements.
type internalPlatformGuestTaskDao = *internal.PlatformGuestTaskDao

// platformGuestTaskDao is the data access object for table platform_guest_task.
// You can define custom methods on it to extend its functionality as you wish.
type platformGuestTaskDao struct {
	internalPlatformGuestTaskDao
}

var (
	// PlatformGuestTask is globally public accessible object for table platform_guest_task operations.
	PlatformGuestTask = platformGuestTaskDao{
		internal.NewPlatformGuestTaskDao(),
	}
)

// Fill with you ideas below.
