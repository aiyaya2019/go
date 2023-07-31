// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TranslateSeatLangDao is the data access object for table translate_seat_lang.
type TranslateSeatLangDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns TranslateSeatLangColumns // columns contains all the column names of Table for convenient usage.
}

// TranslateSeatLangColumns defines and stores column names for table translate_seat_lang.
type TranslateSeatLangColumns struct {
	Id              string // 主键id
	MId             string // 会议id
	TranslateSeatId string // 同传座位translate_seat的id
	TranslateLangId string // 语言id
	Sort            string // 排序id
}

// translateSeatLangColumns holds the columns for table translate_seat_lang.
var translateSeatLangColumns = TranslateSeatLangColumns{
	Id:              "id",
	MId:             "m_id",
	TranslateSeatId: "translate_seat_id",
	TranslateLangId: "translate_lang_id",
	Sort:            "sort",
}

// NewTranslateSeatLangDao creates and returns a new DAO object for table data access.
func NewTranslateSeatLangDao() *TranslateSeatLangDao {
	return &TranslateSeatLangDao{
		group:   "default",
		table:   "translate_seat_lang",
		columns: translateSeatLangColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TranslateSeatLangDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TranslateSeatLangDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TranslateSeatLangDao) Columns() TranslateSeatLangColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TranslateSeatLangDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TranslateSeatLangDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TranslateSeatLangDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
