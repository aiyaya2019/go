// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mygf/internal/dao/internal"
)

// internalConfigCustomDao is internal type for wrapping internal DAO implements.
type internalConfigCustomDao = *internal.ConfigCustomDao

// configCustomDao is the data access object for table config_custom.
// You can define custom methods on it to extend its functionality as you wish.
type configCustomDao struct {
	internalConfigCustomDao
}

var (
	// ConfigCustom is globally public accessible object for table config_custom operations.
	ConfigCustom = configCustomDao{
		internal.NewConfigCustomDao(),
	}
)

// Fill with you ideas below.