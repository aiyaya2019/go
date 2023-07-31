// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"demo1/internal/dao/internal"
)

// internalRoomLayoutDao is internal type for wrapping internal DAO implements.
type internalRoomLayoutDao = *internal.RoomLayoutDao

// roomLayoutDao is the data access object for table room_layout.
// You can define custom methods on it to extend its functionality as you wish.
type roomLayoutDao struct {
	internalRoomLayoutDao
}

var (
	// RoomLayout is globally public accessible object for table room_layout operations.
	RoomLayout = roomLayoutDao{
		internal.NewRoomLayoutDao(),
	}
)

// Fill with you ideas below.
