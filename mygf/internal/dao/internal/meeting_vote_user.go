// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingVoteUserDao is the data access object for table meeting_vote_user.
type MeetingVoteUserDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns MeetingVoteUserColumns // columns contains all the column names of Table for convenient usage.
}

// MeetingVoteUserColumns defines and stores column names for table meeting_vote_user.
type MeetingVoteUserColumns struct {
	MeetingUuid     string // 会议uuid，meeting表的uuid
	UserUuid        string // 用户uuid，user表的uuid
	MeetingVoteUuid string // 投票uuid，meeting_vote表的uuid
	SystemFileUuid  string // 文件uuid，system_file表的uuid
	CreatedAt       string // 创建时间
	UpdatedAt       string // 更新时间
}

// meetingVoteUserColumns holds the columns for table meeting_vote_user.
var meetingVoteUserColumns = MeetingVoteUserColumns{
	MeetingUuid:     "meeting_uuid",
	UserUuid:        "user_uuid",
	MeetingVoteUuid: "meeting_vote_uuid",
	SystemFileUuid:  "system_file_uuid",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewMeetingVoteUserDao creates and returns a new DAO object for table data access.
func NewMeetingVoteUserDao() *MeetingVoteUserDao {
	return &MeetingVoteUserDao{
		group:   "default",
		table:   "meeting_vote_user",
		columns: meetingVoteUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MeetingVoteUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MeetingVoteUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MeetingVoteUserDao) Columns() MeetingVoteUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MeetingVoteUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MeetingVoteUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MeetingVoteUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
