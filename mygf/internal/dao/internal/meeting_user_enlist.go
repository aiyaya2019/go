// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingUserEnlistDao is the data access object for table meeting_user_enlist.
type MeetingUserEnlistDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns MeetingUserEnlistColumns // columns contains all the column names of Table for convenient usage.
}

// MeetingUserEnlistColumns defines and stores column names for table meeting_user_enlist.
type MeetingUserEnlistColumns struct {
	Id          string // 逻辑ID
	MeetingUuid string // 会议uuid，meeting表的uuid
	UserUuid    string // 用户uuid，user表的uuid
	Sort        string // 序号
	Username    string // 名称
	Unit        string // 单位
	Dept        string // 部门
	Position    string // 职务
	Mobile      string // 联系电话
	IsFirst     string // 是否新注册人员:0否;1是
	Mark        string // 备注
	CreateTime  string // 创建时间
	Status      string // 审核状态:-1不通过;0待审核;1审核通过
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// meetingUserEnlistColumns holds the columns for table meeting_user_enlist.
var meetingUserEnlistColumns = MeetingUserEnlistColumns{
	Id:          "id",
	MeetingUuid: "meeting_uuid",
	UserUuid:    "user_uuid",
	Sort:        "sort",
	Username:    "username",
	Unit:        "unit",
	Dept:        "dept",
	Position:    "position",
	Mobile:      "mobile",
	IsFirst:     "is_first",
	Mark:        "mark",
	CreateTime:  "create_time",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewMeetingUserEnlistDao creates and returns a new DAO object for table data access.
func NewMeetingUserEnlistDao() *MeetingUserEnlistDao {
	return &MeetingUserEnlistDao{
		group:   "default",
		table:   "meeting_user_enlist",
		columns: meetingUserEnlistColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MeetingUserEnlistDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MeetingUserEnlistDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MeetingUserEnlistDao) Columns() MeetingUserEnlistColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MeetingUserEnlistDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MeetingUserEnlistDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MeetingUserEnlistDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
