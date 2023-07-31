// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TranslateLangDao is the data access object for table translate_lang.
type TranslateLangDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns TranslateLangColumns // columns contains all the column names of Table for convenient usage.
}

// TranslateLangColumns defines and stores column names for table translate_lang.
type TranslateLangColumns struct {
	Id        string //
	ChannelNo string // 通道id
}

// translateLangColumns holds the columns for table translate_lang.
var translateLangColumns = TranslateLangColumns{
	Id:        "id",
	ChannelNo: "channel_no",
}

// NewTranslateLangDao creates and returns a new DAO object for table data access.
func NewTranslateLangDao() *TranslateLangDao {
	return &TranslateLangDao{
		group:   "default",
		table:   "translate_lang",
		columns: translateLangColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TranslateLangDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TranslateLangDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TranslateLangDao) Columns() TranslateLangColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TranslateLangDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TranslateLangDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TranslateLangDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
