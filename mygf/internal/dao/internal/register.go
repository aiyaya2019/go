// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RegisterDao is the data access object for table register.
type RegisterDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns RegisterColumns // columns contains all the column names of Table for convenient usage.
}

// RegisterColumns defines and stores column names for table register.
type RegisterColumns struct {
	Id        string // 逻辑ID
	Content   string // 注册码
	Status    string // 状态
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// registerColumns holds the columns for table register.
var registerColumns = RegisterColumns{
	Id:        "id",
	Content:   "content",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewRegisterDao creates and returns a new DAO object for table data access.
func NewRegisterDao() *RegisterDao {
	return &RegisterDao{
		group:   "default",
		table:   "register",
		columns: registerColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RegisterDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RegisterDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RegisterDao) Columns() RegisterColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RegisterDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RegisterDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RegisterDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}