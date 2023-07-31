// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ConfigCustomDao is the data access object for table config_custom.
type ConfigCustomDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns ConfigCustomColumns // columns contains all the column names of Table for convenient usage.
}

// ConfigCustomColumns defines and stores column names for table config_custom.
type ConfigCustomColumns struct {
	Id             string //
	SysName        string //
	Logo           string //
	Background     string //
	DefaultSetting string //
	Type           string //
	CreatedAt      string //
	UpdatedAt      string //
}

// configCustomColumns holds the columns for table config_custom.
var configCustomColumns = ConfigCustomColumns{
	Id:             "id",
	SysName:        "sys_name",
	Logo:           "logo",
	Background:     "background",
	DefaultSetting: "default_setting",
	Type:           "type",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewConfigCustomDao creates and returns a new DAO object for table data access.
func NewConfigCustomDao() *ConfigCustomDao {
	return &ConfigCustomDao{
		group:   "default",
		table:   "config_custom",
		columns: configCustomColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ConfigCustomDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ConfigCustomDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ConfigCustomDao) Columns() ConfigCustomColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ConfigCustomDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ConfigCustomDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ConfigCustomDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
