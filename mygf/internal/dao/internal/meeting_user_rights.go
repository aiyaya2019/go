// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingUserRightsDao is the data access object for table meeting_user_rights.
type MeetingUserRightsDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns MeetingUserRightsColumns // columns contains all the column names of Table for convenient usage.
}

// MeetingUserRightsColumns defines and stores column names for table meeting_user_rights.
type MeetingUserRightsColumns struct {
	Id          string // 逻辑id
	MeetingUuid string // 会议uuid，meeting表的uuid
	UserUuid    string // 用户uuid，user表的uuid
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// meetingUserRightsColumns holds the columns for table meeting_user_rights.
var meetingUserRightsColumns = MeetingUserRightsColumns{
	Id:          "id",
	MeetingUuid: "meeting_uuid",
	UserUuid:    "user_uuid",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewMeetingUserRightsDao creates and returns a new DAO object for table data access.
func NewMeetingUserRightsDao() *MeetingUserRightsDao {
	return &MeetingUserRightsDao{
		group:   "default",
		table:   "meeting_user_rights",
		columns: meetingUserRightsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MeetingUserRightsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MeetingUserRightsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MeetingUserRightsDao) Columns() MeetingUserRightsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MeetingUserRightsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MeetingUserRightsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MeetingUserRightsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
