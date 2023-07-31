// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingSummaryDao is the data access object for table meeting_summary.
type MeetingSummaryDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns MeetingSummaryColumns // columns contains all the column names of Table for convenient usage.
}

// MeetingSummaryColumns defines and stores column names for table meeting_summary.
type MeetingSummaryColumns struct {
	Id                       string // 逻辑id
	MeetingUuid              string // 会议uuid，meeting表的uuid
	SystemFileUuid           string // 纪要文件uuid，system_file表的uuid
	UserUuid                 string // 纪要上传者用户uuid, user表的uuid
	Status                   string // 状态 0草稿 1会签中 2已会签
	IsUpdate                 string // 是否已更新:0否；1是
	FileFormatSystemFileUuid string // 结束会签原始文件；file_format表的system_file_uuid
	MeetingDatumUuid         string // 关联议题uuid;meeting_datum表的uuid
	CreatedAt                string // 创建时间
	UpdatedAt                string // 更新时间
	SyncInfo                 string //
}

// meetingSummaryColumns holds the columns for table meeting_summary.
var meetingSummaryColumns = MeetingSummaryColumns{
	Id:                       "id",
	MeetingUuid:              "meeting_uuid",
	SystemFileUuid:           "system_file_uuid",
	UserUuid:                 "user_uuid",
	Status:                   "status",
	IsUpdate:                 "is_update",
	FileFormatSystemFileUuid: "file_format_system_file_uuid",
	MeetingDatumUuid:         "meeting_datum_uuid",
	CreatedAt:                "created_at",
	UpdatedAt:                "updated_at",
	SyncInfo:                 "sync_info",
}

// NewMeetingSummaryDao creates and returns a new DAO object for table data access.
func NewMeetingSummaryDao() *MeetingSummaryDao {
	return &MeetingSummaryDao{
		group:   "default",
		table:   "meeting_summary",
		columns: meetingSummaryColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MeetingSummaryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MeetingSummaryDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MeetingSummaryDao) Columns() MeetingSummaryColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MeetingSummaryDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MeetingSummaryDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MeetingSummaryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
