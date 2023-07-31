// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OperationLogDao is the data access object for table operation_log.
type OperationLogDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns OperationLogColumns // columns contains all the column names of Table for convenient usage.
}

// OperationLogColumns defines and stores column names for table operation_log.
type OperationLogColumns struct {
	Id            string // 逻辑ID
	UserUuid      string // 操作用户uuid，user表的uuid
	Module        string // 模块
	Operation     string // 操作
	OperationType string // 操作类型
	Detail        string // 操作详细信息
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
}

// operationLogColumns holds the columns for table operation_log.
var operationLogColumns = OperationLogColumns{
	Id:            "id",
	UserUuid:      "user_uuid",
	Module:        "module",
	Operation:     "operation",
	OperationType: "operation_type",
	Detail:        "detail",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewOperationLogDao creates and returns a new DAO object for table data access.
func NewOperationLogDao() *OperationLogDao {
	return &OperationLogDao{
		group:   "default",
		table:   "operation_log",
		columns: operationLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OperationLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OperationLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OperationLogDao) Columns() OperationLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OperationLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OperationLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OperationLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
