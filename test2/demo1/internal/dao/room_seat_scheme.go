// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"demo1/internal/dao/internal"
)

// internalRoomSeatSchemeDao is internal type for wrapping internal DAO implements.
type internalRoomSeatSchemeDao = *internal.RoomSeatSchemeDao

// roomSeatSchemeDao is the data access object for table room_seat_scheme.
// You can define custom methods on it to extend its functionality as you wish.
type roomSeatSchemeDao struct {
	internalRoomSeatSchemeDao
}

var (
	// RoomSeatScheme is globally public accessible object for table room_seat_scheme operations.
	RoomSeatScheme = roomSeatSchemeDao{
		internal.NewRoomSeatSchemeDao(),
	}
)

// Fill with you ideas below.
