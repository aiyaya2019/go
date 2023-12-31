// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
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
	Id          string //
	MeetingUuid string //
	Inst        string //
	ServerId    string //
	InsertTime  string //
	ExecTime    string //
	Sort        string //
	Content     string //
	Status      string //
	Mark        string //
	CreatedAt   string //
	UpdatedAt   string //
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
