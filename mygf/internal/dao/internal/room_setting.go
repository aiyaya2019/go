// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoomSettingDao is the data access object for table room_setting.
type RoomSettingDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns RoomSettingColumns // columns contains all the column names of Table for convenient usage.
}

// RoomSettingColumns defines and stores column names for table room_setting.
type RoomSettingColumns struct {
	Id        string // 逻辑id
	RoomId    string // 会议室id,room表的id
	Type      string // 功能类型
	Content   string // 设置的json内容
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// roomSettingColumns holds the columns for table room_setting.
var roomSettingColumns = RoomSettingColumns{
	Id:        "id",
	RoomId:    "room_id",
	Type:      "type",
	Content:   "content",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewRoomSettingDao creates and returns a new DAO object for table data access.
func NewRoomSettingDao() *RoomSettingDao {
	return &RoomSettingDao{
		group:   "default",
		table:   "room_setting",
		columns: roomSettingColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoomSettingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoomSettingDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoomSettingDao) Columns() RoomSettingColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoomSettingDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoomSettingDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoomSettingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}