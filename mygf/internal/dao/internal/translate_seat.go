// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TranslateSeatDao is the data access object for table translate_seat.
type TranslateSeatDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns TranslateSeatColumns // columns contains all the column names of Table for convenient usage.
}

// TranslateSeatColumns defines and stores column names for table translate_seat.
type TranslateSeatColumns struct {
	Id     string // 主键id
	SeatNo string // 编号
	Title  string // 标题
	Xway   string // x坐标
	Yway   string // y坐标
}

// translateSeatColumns holds the columns for table translate_seat.
var translateSeatColumns = TranslateSeatColumns{
	Id:     "id",
	SeatNo: "seat_no",
	Title:  "title",
	Xway:   "xway",
	Yway:   "yway",
}

// NewTranslateSeatDao creates and returns a new DAO object for table data access.
func NewTranslateSeatDao() *TranslateSeatDao {
	return &TranslateSeatDao{
		group:   "default",
		table:   "translate_seat",
		columns: translateSeatColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TranslateSeatDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TranslateSeatDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TranslateSeatDao) Columns() TranslateSeatColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TranslateSeatDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TranslateSeatDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TranslateSeatDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
