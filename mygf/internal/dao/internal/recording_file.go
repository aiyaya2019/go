// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RecordingFileDao is the data access object for table recording_file.
type RecordingFileDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns RecordingFileColumns // columns contains all the column names of Table for convenient usage.
}

// RecordingFileColumns defines and stores column names for table recording_file.
type RecordingFileColumns struct {
	Id          string //
	MeetingUuid string // 会议uuid，meeting表的uuid
	TransId     string // 语音转写那边的id
	Type        string // 0=>单个话筒，1=》所有话筒，2全部的压缩包
	FileName    string // 文件名称
	FileSize    string // 文件大小
	FilePath    string // 文件路径
	FileTime    string // 转写时长
	StartTime   string // 转写开始时间
	EndTime     string // 转写结束时间
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// recordingFileColumns holds the columns for table recording_file.
var recordingFileColumns = RecordingFileColumns{
	Id:          "id",
	MeetingUuid: "meeting_uuid",
	TransId:     "trans_id",
	Type:        "type",
	FileName:    "file_name",
	FileSize:    "file_size",
	FilePath:    "file_path",
	FileTime:    "file_time",
	StartTime:   "start_time",
	EndTime:     "end_time",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewRecordingFileDao creates and returns a new DAO object for table data access.
func NewRecordingFileDao() *RecordingFileDao {
	return &RecordingFileDao{
		group:   "default",
		table:   "recording_file",
		columns: recordingFileColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RecordingFileDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RecordingFileDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RecordingFileDao) Columns() RecordingFileColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RecordingFileDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RecordingFileDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RecordingFileDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
