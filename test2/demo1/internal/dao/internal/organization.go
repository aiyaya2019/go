// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrganizationDao is the data access object for table organization.
type OrganizationDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns OrganizationColumns // columns contains all the column names of Table for convenient usage.
}

// OrganizationColumns defines and stores column names for table organization.
type OrganizationColumns struct {
	Id           string //
	Uuid         string //
	Puuid        string //
	PlatformUuid string //
	Level        string //
	Name         string //
	Sort         string //
	IsTmp        string //
	Status       string //
	Deleted      string //
	CreatedAt    string //
	UpdatedAt    string //
}

// organizationColumns holds the columns for table organization.
var organizationColumns = OrganizationColumns{
	Id:           "id",
	Uuid:         "uuid",
	Puuid:        "puuid",
	PlatformUuid: "platform_uuid",
	Level:        "level",
	Name:         "name",
	Sort:         "sort",
	IsTmp:        "is_tmp",
	Status:       "status",
	Deleted:      "deleted",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewOrganizationDao creates and returns a new DAO object for table data access.
func NewOrganizationDao() *OrganizationDao {
	return &OrganizationDao{
		group:   "default",
		table:   "organization",
		columns: organizationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrganizationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrganizationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OrganizationDao) Columns() OrganizationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrganizationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrganizationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrganizationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
