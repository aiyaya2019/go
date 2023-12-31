// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"demo1/internal/dao/internal"
)

// internalSystemFileDao is internal type for wrapping internal DAO implements.
type internalSystemFileDao = *internal.SystemFileDao

// systemFileDao is the data access object for table system_file.
// You can define custom methods on it to extend its functionality as you wish.
type systemFileDao struct {
	internalSystemFileDao
}

var (
	// SystemFile is globally public accessible object for table system_file operations.
	SystemFile = systemFileDao{
		internal.NewSystemFileDao(),
	}
)

// Fill with you ideas below.
