// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingEnrollMessageDao is the data access object for table meeting_enroll_message.
type MeetingEnrollMessageDao struct {
	table   string                      // table is the underlying table name of the DAO.
	group   string                      // group is the database configuration group name of current DAO.
	columns MeetingEnrollMessageColumns // columns contains all the column names of Table for convenient usage.
}

// MeetingEnrollMessageColumns defines and stores column names for table meeting_enroll_message.
type MeetingEnrollMessageColumns struct {
	Id          string // 逻辑id
	MeetingUuid string // 会议uuid，meeting表的uuid
	Contact     string // 报名联系
	Link        string // 报名链接
	TimeType    string // 0:开始到结束时间内报名;1:会议状态开始前可以报名
	StartTime   string // 开始时间
	EndTime     string // 结束时间
	QrcodePath  string // 二维码保存路径
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// meetingEnrollMessageColumns holds the columns for table meeting_enroll_message.
var meetingEnrollMessageColumns = MeetingEnrollMessageColumns{
	Id:          "id",
	MeetingUuid: "meeting_uuid",
	Contact:     "contact",
	Link:        "link",
	TimeType:    "time_type",
	StartTime:   "start_time",
	EndTime:     "end_time",
	QrcodePath:  "qrcode_path",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewMeetingEnrollMessageDao creates and returns a new DAO object for table data access.
func NewMeetingEnrollMessageDao() *MeetingEnrollMessageDao {
	return &MeetingEnrollMessageDao{
		group:   "default",
		table:   "meeting_enroll_message",
		columns: meetingEnrollMessageColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MeetingEnrollMessageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MeetingEnrollMessageDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MeetingEnrollMessageDao) Columns() MeetingEnrollMessageColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MeetingEnrollMessageDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MeetingEnrollMessageDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MeetingEnrollMessageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
