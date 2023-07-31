// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TaskQueueDao is the data access object for table task_queue.
type TaskQueueDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns TaskQueueColumns // columns contains all the column names of Table for convenient usage.
}

// TaskQueueColumns defines and stores column names for table task_queue.
type TaskQueueColumns struct {
	Id          string // 逻辑ID
	MeetingUuid string // 会议uuid，meeting表的uuid
	Inst        string // 任务指令，协议定的指令号
	ServerId    string // 服务器ID，info_server表的id
	InsertTime  string // 插入时间
	ExecTime    string // 执行时间
	Sort        string // 权重
	Content     string // json
	Status      string // 状态： 0=》等待，1=》进行中，2=》成功，3=》失败，4=》服务器离线
	Mark        string // 备注
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// taskQueueColumns holds the columns for table task_queue.
var taskQueueColumns = TaskQueueColumns{
	Id:          "id",
	MeetingUuid: "meeting_uuid",
	Inst:        "inst",
	ServerId:    "server_id",
	InsertTime:  "insert_time",
	ExecTime:    "exec_time",
	Sort:        "sort",
	Content:     "content",
	Status:      "status",
	Mark:        "mark",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewTaskQueueDao creates and returns a new DAO object for table data access.
func NewTaskQueueDao() *TaskQueueDao {
	return &TaskQueueDao{
		group:   "default",
		table:   "task_queue",
		columns: taskQueueColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TaskQueueDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TaskQueueDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TaskQueueDao) Columns() TaskQueueColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TaskQueueDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TaskQueueDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TaskQueueDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
