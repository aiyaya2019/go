// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"demo1/internal/dao/internal"
)

// internalRoomSeatDao is internal type for wrapping internal DAO implements.
type internalRoomSeatDao = *internal.RoomSeatDao

// roomSeatDao is the data access object for table room_seat.
// You can define custom methods on it to extend its functionality as you wish.
type roomSeatDao struct {
	internalRoomSeatDao
}

var (
	// RoomSeat is globally public accessible object for table room_seat operations.
	RoomSeat = roomSeatDao{
		internal.NewRoomSeatDao(),
	}
)

// Fill with you ideas below.
