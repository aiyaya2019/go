// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemFormatDao is the data access object for table system_format.
type SystemFormatDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns SystemFormatColumns // columns contains all the column names of Table for convenient usage.
}

// SystemFormatColumns defines and stores column names for table system_format.
type SystemFormatColumns struct {
	Id           string //
	Type         string //
	UploadSize   string //
	Size         string //
	SystemFormat string //
	DefineFormat string //
	Name         string //
	CreatedAt    string //
	UpdatedAt    string //
}

// systemFormatColumns holds the columns for table system_format.
var systemFormatColumns = SystemFormatColumns{
	Id:           "id",
	Type:         "type",
	UploadSize:   "upload_size",
	Size:         "size",
	SystemFormat: "system_format",
	DefineFormat: "define_format",
	Name:         "name",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewSystemFormatDao creates and returns a new DAO object for table data access.
func NewSystemFormatDao() *SystemFormatDao {
	return &SystemFormatDao{
		group:   "default",
		table:   "system_format",
		columns: systemFormatColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemFormatDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemFormatDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemFormatDao) Columns() SystemFormatColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemFormatDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemFormatDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemFormatDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}