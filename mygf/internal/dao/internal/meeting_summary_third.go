// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingSummaryThirdDao is the data access object for table meeting_summary_third.
type MeetingSummaryThirdDao struct {
	table   string                     // table is the underlying table name of the DAO.
	group   string                     // group is the database configuration group name of current DAO.
	columns MeetingSummaryThirdColumns // columns contains all the column names of Table for convenient usage.
}

// MeetingSummaryThirdColumns defines and stores column names for table meeting_summary_third.
type MeetingSummaryThirdColumns struct {
	MeetingSummaryId string // 会议纪要id，meeting_summary的id
	PhoneticTranId   string // 语音转写那边的id，作对应关系
	Creator          string // 创建者
	Modifier         string // 修改者
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
}

// meetingSummaryThirdColumns holds the columns for table meeting_summary_third.
var meetingSummaryThirdColumns = MeetingSummaryThirdColumns{
	MeetingSummaryId: "meeting_summary_id",
	PhoneticTranId:   "phonetic_tran_id",
	Creator:          "creator",
	Modifier:         "modifier",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

// NewMeetingSummaryThirdDao creates and returns a new DAO object for table data access.
func NewMeetingSummaryThirdDao() *MeetingSummaryThirdDao {
	return &MeetingSummaryThirdDao{
		group:   "default",
		table:   "meeting_summary_third",
		columns: meetingSummaryThirdColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MeetingSummaryThirdDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MeetingSummaryThirdDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MeetingSummaryThirdDao) Columns() MeetingSummaryThirdColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MeetingSummaryThirdDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MeetingSummaryThirdDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MeetingSummaryThirdDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
