// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserOtherLoginDao is the data access object for table user_other_login.
type UserOtherLoginDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns UserOtherLoginColumns // columns contains all the column names of Table for convenient usage.
}

// UserOtherLoginColumns defines and stores column names for table user_other_login.
type UserOtherLoginColumns struct {
	Id        string // 逻辑ID
	UserUuid  string // 用户uuid，user表的uuid
	Name      string // 名称
	Pass      string // 特征码
	Type      string // 1:指纹
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// userOtherLoginColumns holds the columns for table user_other_login.
var userOtherLoginColumns = UserOtherLoginColumns{
	Id:        "id",
	UserUuid:  "user_uuid",
	Name:      "name",
	Pass:      "pass",
	Type:      "type",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewUserOtherLoginDao creates and returns a new DAO object for table data access.
func NewUserOtherLoginDao() *UserOtherLoginDao {
	return &UserOtherLoginDao{
		group:   "default",
		table:   "user_other_login",
		columns: userOtherLoginColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserOtherLoginDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserOtherLoginDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserOtherLoginDao) Columns() UserOtherLoginColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserOtherLoginDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserOtherLoginDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserOtherLoginDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
