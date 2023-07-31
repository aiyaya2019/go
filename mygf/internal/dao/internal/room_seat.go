// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoomSeatDao is the data access object for table room_seat.
type RoomSeatDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns RoomSeatColumns // columns contains all the column names of Table for convenient usage.
}

// RoomSeatColumns defines and stores column names for table room_seat.
type RoomSeatColumns struct {
	Id               string // 逻辑id
	RoomId           string // 会议室id，room表的id
	UserUuid         string // 用户uuid，user表的uuid
	RoomMicrophoneId string // 麦克风id，room_microphone表的id
	TerminalId       string // 终端id，terminal表的id
	LifterTerminalId string // 升降器id，terminal表的id
	ElectronicId     string // 电子桌牌id
	SeatNo           string // 座位序号
	Title            string // 座位描述
	Status           string // 状态0:停用,1启用
	Xway             string // X轴
	Yway             string // Y轴
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
}

// roomSeatColumns holds the columns for table room_seat.
var roomSeatColumns = RoomSeatColumns{
	Id:               "id",
	RoomId:           "room_id",
	UserUuid:         "user_uuid",
	RoomMicrophoneId: "room_microphone_id",
	TerminalId:       "terminal_id",
	LifterTerminalId: "lifter_terminal_id",
	ElectronicId:     "electronic_id",
	SeatNo:           "seat_no",
	Title:            "title",
	Status:           "status",
	Xway:             "xway",
	Yway:             "yway",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

// NewRoomSeatDao creates and returns a new DAO object for table data access.
func NewRoomSeatDao() *RoomSeatDao {
	return &RoomSeatDao{
		group:   "default",
		table:   "room_seat",
		columns: roomSeatColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoomSeatDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoomSeatDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoomSeatDao) Columns() RoomSeatColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoomSeatDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoomSeatDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoomSeatDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
