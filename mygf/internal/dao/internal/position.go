// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PositionDao is the data access object for table position.
type PositionDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns PositionColumns // columns contains all the column names of Table for convenient usage.
}

// PositionColumns defines and stores column names for table position.
type PositionColumns struct {
	Id             string // 自增ID
	PositionTypeId string // 职位组id；position_type表的id
	Name           string // 职位名称
	Sort           string // 排序
	Deleted        string // 软删除：0未删除；1已删除
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// positionColumns holds the columns for table position.
var positionColumns = PositionColumns{
	Id:             "id",
	PositionTypeId: "position_type_id",
	Name:           "name",
	Sort:           "sort",
	Deleted:        "deleted",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewPositionDao creates and returns a new DAO object for table data access.
func NewPositionDao() *PositionDao {
	return &PositionDao{
		group:   "default",
		table:   "position",
		columns: positionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PositionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PositionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PositionDao) Columns() PositionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PositionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PositionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PositionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
