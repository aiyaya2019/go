// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlatformPushTaskDao is the data access object for table platform_push_task.
type PlatformPushTaskDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns PlatformPushTaskColumns // columns contains all the column names of Table for convenient usage.
}

// PlatformPushTaskColumns defines and stores column names for table platform_push_task.
type PlatformPushTaskColumns struct {
	Id           string // 主键
	PlatformUuid string // 平台uuid
	MeetingUuid  string // 会议uuid，meeting表的uuid
	Status       string // 0:未邀请1:进行中2:成功3:失败4:编辑中5:编辑成功6:编辑失败7:删除中8:删除成功9:删除失败10:移除中11:移除成功12:移除失败13:取消中14:取消成功15:取消失败16:已退出,17:已拒绝,18:副平台结束
	Operate      string // 1add 2edit 3del4:移除5:取消
	JoinTime     string // 加入时间
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
}

// platformPushTaskColumns holds the columns for table platform_push_task.
var platformPushTaskColumns = PlatformPushTaskColumns{
	Id:           "id",
	PlatformUuid: "platform_uuid",
	MeetingUuid:  "meeting_uuid",
	Status:       "status",
	Operate:      "operate",
	JoinTime:     "join_time",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewPlatformPushTaskDao creates and returns a new DAO object for table data access.
func NewPlatformPushTaskDao() *PlatformPushTaskDao {
	return &PlatformPushTaskDao{
		group:   "default",
		table:   "platform_push_task",
		columns: platformPushTaskColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PlatformPushTaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PlatformPushTaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PlatformPushTaskDao) Columns() PlatformPushTaskColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PlatformPushTaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PlatformPushTaskDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PlatformPushTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
