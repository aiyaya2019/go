// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"demo1/internal/dao/internal"
)

// internalSystemFormatDao is internal type for wrapping internal DAO implements.
type internalSystemFormatDao = *internal.SystemFormatDao

// systemFormatDao is the data access object for table system_format.
// You can define custom methods on it to extend its functionality as you wish.
type systemFormatDao struct {
	internalSystemFormatDao
}

var (
	// SystemFormat is globally public accessible object for table system_format operations.
	SystemFormat = systemFormatDao{
		internal.NewSystemFormatDao(),
	}
)

// Fill with you ideas below.