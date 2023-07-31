// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TerminalConcurrenceDao is the data access object for table terminal_concurrence.
type TerminalConcurrenceDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns TerminalConcurrenceColumns // columns contains all the column names of Table for convenient usage.
}

// TerminalConcurrenceColumns defines and stores column names for table terminal_concurrence.
type TerminalConcurrenceColumns struct {
	Id          string // 主键
	MeetingUuid string // 会议uuid，meeting表的uuid
	TotalUser   string // 单节点客户端并发数
	ClientIds   string // 勾选的客户端id逗号隔开，type=8
	Usernames   string // 随机生成用户名erupt_test1
	UserMax     string // 最大用户值
	Type        string // 操作1->添加3->删除
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// terminalConcurrenceColumns holds the columns for table terminal_concurrence.
var terminalConcurrenceColumns = TerminalConcurrenceColumns{
	Id:          "id",
	MeetingUuid: "meeting_uuid",
	TotalUser:   "total_user",
	ClientIds:   "client_ids",
	Usernames:   "usernames",
	UserMax:     "user_max",
	Type:        "type",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewTerminalConcurrenceDao creates and returns a new DAO object for table data access.
func NewTerminalConcurrenceDao() *TerminalConcurrenceDao {
	return &TerminalConcurrenceDao{
		group:   "default",
		table:   "terminal_concurrence",
		columns: terminalConcurrenceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TerminalConcurrenceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TerminalConcurrenceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TerminalConcurrenceDao) Columns() TerminalConcurrenceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TerminalConcurrenceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TerminalConcurrenceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TerminalConcurrenceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
