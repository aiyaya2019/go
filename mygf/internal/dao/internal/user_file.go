// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserFileDao is the data access object for table user_file.
type UserFileDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns UserFileColumns // columns contains all the column names of Table for convenient usage.
}

// UserFileColumns defines and stores column names for table user_file.
type UserFileColumns struct {
	MeetingUuid    string // 会议uuid，meeting表的uuid
	UserUuid       string // 用户uuid，user表的uuid
	SystemFileUuid string // 文件uuid，system_file表的uuid
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// userFileColumns holds the columns for table user_file.
var userFileColumns = UserFileColumns{
	MeetingUuid:    "meeting_uuid",
	UserUuid:       "user_uuid",
	SystemFileUuid: "system_file_uuid",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewUserFileDao creates and returns a new DAO object for table data access.
func NewUserFileDao() *UserFileDao {
	return &UserFileDao{
		group:   "default",
		table:   "user_file",
		columns: userFileColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserFileDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserFileDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserFileDao) Columns() UserFileColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserFileDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserFileDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserFileDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
