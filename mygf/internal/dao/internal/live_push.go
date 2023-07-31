// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LivePushDao is the data access object for table live_push.
type LivePushDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns LivePushColumns // columns contains all the column names of Table for convenient usage.
}

// LivePushColumns defines and stores column names for table live_push.
type LivePushColumns struct {
	Id             string // 逻辑ID
	SystemFileUuid string // 直播文件uuid，system_file表的uuid
	Pid            string // 进程ID
	Progress       string // 进度
	Status         string // 推流状态，0=》未开始，1 =》开始，2 =》暂停， 3 =》停止， 4=》 错误，5=》异常，6=》等待中
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// livePushColumns holds the columns for table live_push.
var livePushColumns = LivePushColumns{
	Id:             "id",
	SystemFileUuid: "system_file_uuid",
	Pid:            "pid",
	Progress:       "progress",
	Status:         "status",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewLivePushDao creates and returns a new DAO object for table data access.
func NewLivePushDao() *LivePushDao {
	return &LivePushDao{
		group:   "default",
		table:   "live_push",
		columns: livePushColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *LivePushDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *LivePushDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *LivePushDao) Columns() LivePushColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *LivePushDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *LivePushDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *LivePushDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
