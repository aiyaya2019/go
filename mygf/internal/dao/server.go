// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"mygf/internal/dao/internal"
)

// internalServerDao is internal type for wrapping internal DAO implements.
type internalServerDao = *internal.ServerDao

// serverDao is the data access object for table server.
// You can define custom methods on it to extend its functionality as you wish.
type serverDao struct {
	internalServerDao
}

var (
	// Server is globally public accessible object for table server operations.
	Server = serverDao{
		internal.NewServerDao(),
	}
)

// Fill with you ideas below.
