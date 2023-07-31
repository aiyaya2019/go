// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PositionTypeDao is the data access object for table position_type.
type PositionTypeDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns PositionTypeColumns // columns contains all the column names of Table for convenient usage.
}

// PositionTypeColumns defines and stores column names for table position_type.
type PositionTypeColumns struct {
	Id        string // 自增ID
	Name      string // 职位名组称
	Sort      string // 排序
	Deleted   string // 软删除：0未删除；1已删除
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// positionTypeColumns holds the columns for table position_type.
var positionTypeColumns = PositionTypeColumns{
	Id:        "id",
	Name:      "name",
	Sort:      "sort",
	Deleted:   "deleted",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewPositionTypeDao creates and returns a new DAO object for table data access.
func NewPositionTypeDao() *PositionTypeDao {
	return &PositionTypeDao{
		group:   "default",
		table:   "position_type",
		columns: positionTypeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PositionTypeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PositionTypeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PositionTypeDao) Columns() PositionTypeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PositionTypeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PositionTypeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PositionTypeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
