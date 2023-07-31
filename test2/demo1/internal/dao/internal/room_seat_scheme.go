// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoomSeatSchemeDao is the data access object for table room_seat_scheme.
type RoomSeatSchemeDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns RoomSeatSchemeColumns // columns contains all the column names of Table for convenient usage.
}

// RoomSeatSchemeColumns defines and stores column names for table room_seat_scheme.
type RoomSeatSchemeColumns struct {
	Id          string //
	MeetingUuid string //
	Title       string //
	Type        string //
	Status      string //
	CreatedAt   string //
	UpdatedAt   string //
}

// roomSeatSchemeColumns holds the columns for table room_seat_scheme.
var roomSeatSchemeColumns = RoomSeatSchemeColumns{
	Id:          "id",
	MeetingUuid: "meeting_uuid",
	Title:       "title",
	Type:        "type",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewRoomSeatSchemeDao creates and returns a new DAO object for table data access.
func NewRoomSeatSchemeDao() *RoomSeatSchemeDao {
	return &RoomSeatSchemeDao{
		group:   "default",
		table:   "room_seat_scheme",
		columns: roomSeatSchemeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoomSeatSchemeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoomSeatSchemeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoomSeatSchemeDao) Columns() RoomSeatSchemeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoomSeatSchemeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoomSeatSchemeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoomSeatSchemeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
