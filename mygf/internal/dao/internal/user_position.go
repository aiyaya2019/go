// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserPositionDao is the data access object for table user_position.
type UserPositionDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns UserPositionColumns // columns contains all the column names of Table for convenient usage.
}

// UserPositionColumns defines and stores column names for table user_position.
type UserPositionColumns struct {
	Id         string // 自增ID
	UserUuid   string // 人员uuid；user表的uuid
	PositionId string // 职位id；position表的id
	Default    string // 是否默认职位：0非默认；1默认
	Deleted    string // 软删除：0未删除；1已删除
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

// userPositionColumns holds the columns for table user_position.
var userPositionColumns = UserPositionColumns{
	Id:         "id",
	UserUuid:   "user_uuid",
	PositionId: "position_id",
	Default:    "default",
	Deleted:    "deleted",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewUserPositionDao creates and returns a new DAO object for table data access.
func NewUserPositionDao() *UserPositionDao {
	return &UserPositionDao{
		group:   "default",
		table:   "user_position",
		columns: userPositionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserPositionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserPositionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserPositionDao) Columns() UserPositionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserPositionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserPositionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserPositionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
