// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mygf/internal/dao/internal"
)

// internalSystemBackupsDao is internal type for wrapping internal DAO implements.
type internalSystemBackupsDao = *internal.SystemBackupsDao

// systemBackupsDao is the data access object for table system_backups.
// You can define custom methods on it to extend its functionality as you wish.
type systemBackupsDao struct {
	internalSystemBackupsDao
}

var (
	// SystemBackups is globally public accessible object for table system_backups operations.
	SystemBackups = systemBackupsDao{
		internal.NewSystemBackupsDao(),
	}
)

// Fill with you ideas below.
