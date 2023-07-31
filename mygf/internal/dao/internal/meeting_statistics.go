// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingStatisticsDao is the data access object for table meeting_statistics.
type MeetingStatisticsDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns MeetingStatisticsColumns // columns contains all the column names of Table for convenient usage.
}

// MeetingStatisticsColumns defines and stores column names for table meeting_statistics.
type MeetingStatisticsColumns struct {
	Id               string // 主键
	MeetingUuid      string // 关联会议uuid
	MeetingUserCount string // 参会人数
	MeetingDuration  string // 会议时长
	SavePaperCount   string // 节约纸张数
	MeetingTime      string // 会议时间
	StartTime        string // 会议开始时间
	EndTime          string // 会议结束时间
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
}

// meetingStatisticsColumns holds the columns for table meeting_statistics.
var meetingStatisticsColumns = MeetingStatisticsColumns{
	Id:               "id",
	MeetingUuid:      "meeting_uuid",
	MeetingUserCount: "meeting_user_count",
	MeetingDuration:  "meeting_duration",
	SavePaperCount:   "save_paper_count",
	MeetingTime:      "meeting_time",
	StartTime:        "start_time",
	EndTime:          "end_time",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

// NewMeetingStatisticsDao creates and returns a new DAO object for table data access.
func NewMeetingStatisticsDao() *MeetingStatisticsDao {
	return &MeetingStatisticsDao{
		group:   "default",
		table:   "meeting_statistics",
		columns: meetingStatisticsColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MeetingStatisticsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MeetingStatisticsDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MeetingStatisticsDao) Columns() MeetingStatisticsColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MeetingStatisticsDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MeetingStatisticsDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MeetingStatisticsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
