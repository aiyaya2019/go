// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
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
	Id                       string //
	MeetingUuid              string //
	SystemFileUuid           string //
	UserUuid                 string //
	Status                   string //
	IsUpdate                 string //
	FileFormatSystemFileUuid string //
	MeetingDatumUuid         string //
	CreatedAt                string //
	UpdatedAt                string //
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