// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"demo1/internal/dao/internal"
)

// internalRoomDao is internal type for wrapping internal DAO implements.
type internalRoomDao = *internal.RoomDao

// roomDao is the data access object for table room.
// You can define custom methods on it to extend its functionality as you wish.
type roomDao struct {
	internalRoomDao
}

var (
	// Room is globally public accessible object for table room operations.
	Room = roomDao{
		internal.NewRoomDao(),
	}
)

// Fill with you ideas below.
