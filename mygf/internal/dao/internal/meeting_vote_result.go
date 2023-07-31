// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingVoteResultDao is the data access object for table meeting_vote_result.
type MeetingVoteResultDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns MeetingVoteResultColumns // columns contains all the column names of Table for convenient usage.
}

// MeetingVoteResultColumns defines and stores column names for table meeting_vote_result.
type MeetingVoteResultColumns struct {
	Id              string // 逻辑ID
	UserUuid        string // 用户uuid，user表的id
	MeetingVoteUuid string // 投票uuid，meeting_vote表uuid
	VoteOptionUuid  string // 选项uuid，vote_option表的uuid
	VoteTime        string // 投票时间
	Mark            string // 备注
	Score           string // 分值
	CreatedAt       string // 创建时间
	UpdatedAt       string // 更新时间
}

// meetingVoteResultColumns holds the columns for table meeting_vote_result.
var meetingVoteResultColumns = MeetingVoteResultColumns{
	Id:              "id",
	UserUuid:        "user_uuid",
	MeetingVoteUuid: "meeting_vote_uuid",
	VoteOptionUuid:  "vote_option_uuid",
	VoteTime:        "vote_time",
	Mark:            "mark",
	Score:           "score",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewMeetingVoteResultDao creates and returns a new DAO object for table data access.
func NewMeetingVoteResultDao() *MeetingVoteResultDao {
	return &MeetingVoteResultDao{
		group:   "default",
		table:   "meeting_vote_result",
		columns: meetingVoteResultColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MeetingVoteResultDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MeetingVoteResultDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MeetingVoteResultDao) Columns() MeetingVoteResultColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MeetingVoteResultDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MeetingVoteResultDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MeetingVoteResultDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
