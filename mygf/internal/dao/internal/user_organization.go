// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserOrganizationDao is the data access object for table user_organization.
type UserOrganizationDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns UserOrganizationColumns // columns contains all the column names of Table for convenient usage.
}

// UserOrganizationColumns defines and stores column names for table user_organization.
type UserOrganizationColumns struct {
	Id                string //
	UserUuid          string // 用户uuid，user表uuid
	OrganizationUuid  string // 组织架构uuid，organization表uuid
	OrganizationLevel string // 层级
	Default           string // 是否默认：0非默认；1默认
	Binding           string // 0虚拟绑定的上级(统计人数用)；1真实绑定所在组织
	Deleted           string // 软删除：0未删除1已删除
}

// userOrganizationColumns holds the columns for table user_organization.
var userOrganizationColumns = UserOrganizationColumns{
	Id:                "id",
	UserUuid:          "user_uuid",
	OrganizationUuid:  "organization_uuid",
	OrganizationLevel: "organization_level",
	Default:           "default",
	Binding:           "binding",
	Deleted:           "deleted",
}

// NewUserOrganizationDao creates and returns a new DAO object for table data access.
func NewUserOrganizationDao() *UserOrganizationDao {
	return &UserOrganizationDao{
		group:   "default",
		table:   "user_organization",
		columns: userOrganizationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserOrganizationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserOrganizationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserOrganizationDao) Columns() UserOrganizationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserOrganizationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserOrganizationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserOrganizationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
