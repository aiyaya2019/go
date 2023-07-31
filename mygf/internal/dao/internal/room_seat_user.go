// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoomSeatUserDao is the data access object for table room_seat_user.
type RoomSeatUserDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns RoomSeatUserColumns // columns contains all the column names of Table for convenient usage.
}

// RoomSeatUserColumns defines and stores column names for table room_seat_user.
type RoomSeatUserColumns struct {
	RoomSeatId       string // 座位id
	MeetingUuid      string // 会议uuid，meeting表的uuid
	UserUuid         string // 用户id，room_seat表的id
	RoomSeatSchemeId string // 排位方案id，room_seat_scheme表的id
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
}

// roomSeatUserColumns holds the columns for table room_seat_user.
var roomSeatUserColumns = RoomSeatUserColumns{
	RoomSeatId:       "room_seat_id",
	MeetingUuid:      "meeting_uuid",
	UserUuid:         "user_uuid",
	RoomSeatSchemeId: "room_seat_scheme_id",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

// NewRoomSeatUserDao creates and returns a new DAO object for table data access.
func NewRoomSeatUserDao() *RoomSeatUserDao {
	return &RoomSeatUserDao{
		group:   "default",
		table:   "room_seat_user",
		columns: roomSeatUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoomSeatUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoomSeatUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoomSeatUserDao) Columns() RoomSeatUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoomSeatUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoomSeatUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoomSeatUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
