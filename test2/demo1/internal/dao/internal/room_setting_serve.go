// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoomSettingServeDao is the data access object for table room_setting_serve.
type RoomSettingServeDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns RoomSettingServeColumns // columns contains all the column names of Table for convenient usage.
}

// RoomSettingServeColumns defines and stores column names for table room_setting_serve.
type RoomSettingServeColumns struct {
	Id        string //
	RoomId    string //
	Name      string //
	Src       string //
	IsSystem  string //
	Sort      string //
	CreatedAt string //
	UpdatedAt string //
}

// roomSettingServeColumns holds the columns for table room_setting_serve.
var roomSettingServeColumns = RoomSettingServeColumns{
	Id:        "id",
	RoomId:    "room_id",
	Name:      "name",
	Src:       "src",
	IsSystem:  "is_system",
	Sort:      "sort",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewRoomSettingServeDao creates and returns a new DAO object for table data access.
func NewRoomSettingServeDao() *RoomSettingServeDao {
	return &RoomSettingServeDao{
		group:   "default",
		table:   "room_setting_serve",
		columns: roomSettingServeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoomSettingServeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoomSettingServeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoomSettingServeDao) Columns() RoomSettingServeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoomSettingServeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoomSettingServeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoomSettingServeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
