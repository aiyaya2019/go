// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingPcMenuDao is the data access object for table meeting_pc_menu.
type MeetingPcMenuDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns MeetingPcMenuColumns // columns contains all the column names of Table for convenient usage.
}

// MeetingPcMenuColumns defines and stores column names for table meeting_pc_menu.
type MeetingPcMenuColumns struct {
	Id        string // 逻辑id
	Pid       string // 父级id
	Name      string // 菜单名称
	Type      string // PC端类型判断
	Status    string // 状态
	Sort      string // 排序
	Have      string // 菜单模式：1卡片模式；2简洁模式；3导航模式；4经典模式
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// meetingPcMenuColumns holds the columns for table meeting_pc_menu.
var meetingPcMenuColumns = MeetingPcMenuColumns{
	Id:        "id",
	Pid:       "pid",
	Name:      "name",
	Type:      "type",
	Status:    "status",
	Sort:      "sort",
	Have:      "have",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewMeetingPcMenuDao creates and returns a new DAO object for table data access.
func NewMeetingPcMenuDao() *MeetingPcMenuDao {
	return &MeetingPcMenuDao{
		group:   "default",
		table:   "meeting_pc_menu",
		columns: meetingPcMenuColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MeetingPcMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MeetingPcMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MeetingPcMenuDao) Columns() MeetingPcMenuColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MeetingPcMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MeetingPcMenuDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MeetingPcMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
