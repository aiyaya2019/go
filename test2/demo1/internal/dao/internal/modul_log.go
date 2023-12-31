// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ModulLogDao is the data access object for table modul_log.
type ModulLogDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns ModulLogColumns // columns contains all the column names of Table for convenient usage.
}

// ModulLogColumns defines and stores column names for table modul_log.
type ModulLogColumns struct {
	Id        string //
	ModulId   string //
	Type      string //
	UserUuid  string //
	Url       string //
	Path      string //
	Size      string //
	Status    string //
	Time      string //
	StartTime string //
	EndTime   string //
	CreatedAt string //
	UpdatedAt string //
}

// modulLogColumns holds the columns for table modul_log.
var modulLogColumns = ModulLogColumns{
	Id:        "id",
	ModulId:   "modul_id",
	Type:      "type",
	UserUuid:  "user_uuid",
	Url:       "url",
	Path:      "path",
	Size:      "size",
	Status:    "status",
	Time:      "time",
	StartTime: "start_time",
	EndTime:   "end_time",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewModulLogDao creates and returns a new DAO object for table data access.
func NewModulLogDao() *ModulLogDao {
	return &ModulLogDao{
		group:   "default",
		table:   "modul_log",
		columns: modulLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ModulLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ModulLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ModulLogDao) Columns() ModulLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ModulLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ModulLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ModulLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
