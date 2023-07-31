// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"demo1/internal/dao/internal"
)

// internalConfigDao is internal type for wrapping internal DAO implements.
type internalConfigDao = *internal.ConfigDao

// configDao is the data access object for table config.
// You can define custom methods on it to extend its functionality as you wish.
type configDao struct {
	internalConfigDao
}

var (
	// Config is globally public accessible object for table config operations.
	Config = configDao{
		internal.NewConfigDao(),
	}
)

// Fill with you ideas below.
