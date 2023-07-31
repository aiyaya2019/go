// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingSummaryCountersignDao is the data access object for table meeting_summary_countersign.
type MeetingSummaryCountersignDao struct {
	table   string                           // table is the underlying table name of the DAO.
	group   string                           // group is the database configuration group name of current DAO.
	columns MeetingSummaryCountersignColumns // columns contains all the column names of Table for convenient usage.
}

// MeetingSummaryCountersignColumns defines and stores column names for table meeting_summary_countersign.
type MeetingSummaryCountersignColumns struct {
	MeetingUuid        string //
	MeetingSummaryId   string //
	UserSystemFileUuid string //
	UserUuid           string //
	Status             string //
	CreatedAt          string //
	UpdatedAt          string //
}

// meetingSummaryCountersignColumns holds the columns for table meeting_summary_countersign.
var meetingSummaryCountersignColumns = MeetingSummaryCountersignColumns{
	MeetingUuid:        "meeting_uuid",
	MeetingSummaryId:   "meeting_summary_id",
	UserSystemFileUuid: "user_system_file_uuid",
	UserUuid:           "user_uuid",
	Status:             "status",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// NewMeetingSummaryCountersignDao creates and returns a new DAO object for table data access.
func NewMeetingSummaryCountersignDao() *MeetingSummaryCountersignDao {
	return &MeetingSummaryCountersignDao{
		group:   "default",
		table:   "meeting_summary_countersign",
		columns: meetingSummaryCountersignColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MeetingSummaryCountersignDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MeetingSummaryCountersignDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MeetingSummaryCountersignDao) Columns() MeetingSummaryCountersignColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MeetingSummaryCountersignDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MeetingSummaryCountersignDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MeetingSummaryCountersignDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}