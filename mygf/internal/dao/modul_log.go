// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mygf/internal/dao/internal"
)

// internalModulLogDao is internal type for wrapping internal DAO implements.
type internalModulLogDao = *internal.ModulLogDao

// modulLogDao is the data access object for table modul_log.
// You can define custom methods on it to extend its functionality as you wish.
type modulLogDao struct {
	internalModulLogDao
}

var (
	// ModulLog is globally public accessible object for table modul_log operations.
	ModulLog = modulLogDao{
		internal.NewModulLogDao(),
	}
)

// Fill with you ideas below.