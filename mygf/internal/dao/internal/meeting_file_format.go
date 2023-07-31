// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingFileFormatDao is the data access object for table meeting_file_format.
type MeetingFileFormatDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns MeetingFileFormatColumns // columns contains all the column names of Table for convenient usage.
}

// MeetingFileFormatColumns defines and stores column names for table meeting_file_format.
type MeetingFileFormatColumns struct {
	Id          string // 逻辑ID
	MeetingUuid string // 会议uuid，meeting表的uuid
	UserUuid    string // 用户uuid，user表的uuid
	Url         string // url路径
	Path        string // 文件路径
	Progress    string // 进度
	Status      string // 转码状态，0=》未开始，1=》转换中，2=》成功，3=》失败
	Time        string // 时间
	Type        string // 类型，0=》无，1=》html to image
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// meetingFileFormatColumns holds the columns for table meeting_file_format.
var meetingFileFormatColumns = MeetingFileFormatColumns{
	Id:          "id",
	MeetingUuid: "meeting_uuid",
	UserUuid:    "user_uuid",
	Url:         "url",
	Path:        "path",
	Progress:    "progress",
	Status:      "status",
	Time:        "time",
	Type:        "type",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewMeetingFileFormatDao creates and returns a new DAO object for table data access.
func NewMeetingFileFormatDao() *MeetingFileFormatDao {
	return &MeetingFileFormatDao{
		group:   "default",
		table:   "meeting_file_format",
		columns: meetingFileFormatColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MeetingFileFormatDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MeetingFileFormatDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MeetingFileFormatDao) Columns() MeetingFileFormatColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MeetingFileFormatDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MeetingFileFormatDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MeetingFileFormatDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
