// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoomDao is the data access object for table room.
type RoomDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns RoomColumns // columns contains all the column names of Table for convenient usage.
}

// RoomColumns defines and stores column names for table room.
type RoomColumns struct {
	Id             string //
	Name           string //
	ServerId       string //
	Roomlayout     string //
	Status         string //
	Mark           string //
	Roombackground string //
	BgWidth        string //
	BgHeight       string //
	TerminalConfig string //
	Transfer       string //
	Electronic     string //
	CreatedAt      string //
	UpdatedAt      string //
}

// roomColumns holds the columns for table room.
var roomColumns = RoomColumns{
	Id:             "id",
	Name:           "name",
	ServerId:       "server_id",
	Roomlayout:     "roomlayout",
	Status:         "status",
	Mark:           "mark",
	Roombackground: "roombackground",
	BgWidth:        "bg_width",
	BgHeight:       "bg_height",
	TerminalConfig: "terminal_config",
	Transfer:       "transfer",
	Electronic:     "electronic",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewRoomDao creates and returns a new DAO object for table data access.
func NewRoomDao() *RoomDao {
	return &RoomDao{
		group:   "default",
		table:   "room",
		columns: roomColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoomDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoomDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoomDao) Columns() RoomColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoomDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoomDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoomDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
