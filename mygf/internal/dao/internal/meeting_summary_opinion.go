// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingSummaryOpinionDao is the data access object for table meeting_summary_opinion.
type MeetingSummaryOpinionDao struct {
	table   string                       // table is the underlying table name of the DAO.
	group   string                       // group is the database configuration group name of current DAO.
	columns MeetingSummaryOpinionColumns // columns contains all the column names of Table for convenient usage.
}

// MeetingSummaryOpinionColumns defines and stores column names for table meeting_summary_opinion.
type MeetingSummaryOpinionColumns struct {
	Id               string //
	MeetingUuid      string // 会议uuid，meeting表的uuid
	UserUuid         string // 用户uuid，user表的uuid
	MeetingSummaryId string // 纪要ID，meeting_summary表的id
	SendTime         string // 发送时间
	Content          string // 纪要内容
	Page             string // 在纪要文件中第几页提意见
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
}

// meetingSummaryOpinionColumns holds the columns for table meeting_summary_opinion.
var meetingSummaryOpinionColumns = MeetingSummaryOpinionColumns{
	Id:               "id",
	MeetingUuid:      "meeting_uuid",
	UserUuid:         "user_uuid",
	MeetingSummaryId: "meeting_summary_id",
	SendTime:         "send_time",
	Content:          "content",
	Page:             "page",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

// NewMeetingSummaryOpinionDao creates and returns a new DAO object for table data access.
func NewMeetingSummaryOpinionDao() *MeetingSummaryOpinionDao {
	return &MeetingSummaryOpinionDao{
		group:   "default",
		table:   "meeting_summary_opinion",
		columns: meetingSummaryOpinionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MeetingSummaryOpinionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MeetingSummaryOpinionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MeetingSummaryOpinionDao) Columns() MeetingSummaryOpinionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MeetingSummaryOpinionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MeetingSummaryOpinionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MeetingSummaryOpinionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
