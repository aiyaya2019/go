// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingRecordContentDao is the data access object for table meeting_record_content.
type MeetingRecordContentDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of current DAO.
	columns MeetingRecordContentColumns // columns contains all the column names of Table for convenient usage.
}

// MeetingRecordContentColumns defines and stores column names for table meeting_record_content.
type MeetingRecordContentColumns struct {
	Id          string //
	MeetingUuid string // 会议uuid，meeting表的uuid
	TerminalId  string // 终端ID，terminal表的id
	HtmlContent string // HTML会议记录内容
	WordContent string // Word会议记录内容
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// meetingRecordContentColumns holds the columns for table meeting_record_content.
var meetingRecordContentColumns = MeetingRecordContentColumns{
	Id:          "id",
	MeetingUuid: "meeting_uuid",
	TerminalId:  "terminal_id",
	HtmlContent: "html_content",
	WordContent: "word_content",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewMeetingRecordContentDao creates and returns a new DAO object for table data access.
func NewMeetingRecordContentDao() *MeetingRecordContentDao {
	return &MeetingRecordContentDao{
		group:   "default",
		table:   "meeting_record_content",
		columns: meetingRecordContentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MeetingRecordContentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MeetingRecordContentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MeetingRecordContentDao) Columns() MeetingRecordContentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MeetingRecordContentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MeetingRecordContentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MeetingRecordContentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
