// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TerminalRepairDao is the data access object for table terminal_repair.
type TerminalRepairDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns TerminalRepairColumns // columns contains all the column names of Table for convenient usage.
}

// TerminalRepairColumns defines and stores column names for table terminal_repair.
type TerminalRepairColumns struct {
	Id          string // 主键
	TerminalIds string // 终端ids，terminal表的id
	FileName    string // 文件名称
	Command     string // 命令/客户端存放路径
	Path        string // 上传文件路径
	Type        string // 类型1->命令,2->文件
	Status      string // 1=>进行中,2=>成功,3=>失败
	CreateTime  string // 创建时间
	Mark        string // 备注
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// terminalRepairColumns holds the columns for table terminal_repair.
var terminalRepairColumns = TerminalRepairColumns{
	Id:          "id",
	TerminalIds: "terminal_ids",
	FileName:    "file_name",
	Command:     "command",
	Path:        "path",
	Type:        "type",
	Status:      "status",
	CreateTime:  "create_time",
	Mark:        "mark",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewTerminalRepairDao creates and returns a new DAO object for table data access.
func NewTerminalRepairDao() *TerminalRepairDao {
	return &TerminalRepairDao{
		group:   "default",
		table:   "terminal_repair",
		columns: terminalRepairColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TerminalRepairDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TerminalRepairDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TerminalRepairDao) Columns() TerminalRepairColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TerminalRepairDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TerminalRepairDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TerminalRepairDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
