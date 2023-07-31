// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingDatumUserDao is the data access object for table meeting_datum_user.
type MeetingDatumUserDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns MeetingDatumUserColumns // columns contains all the column names of Table for convenient usage.
}

// MeetingDatumUserColumns defines and stores column names for table meeting_datum_user.
type MeetingDatumUserColumns struct {
	MeetingUuid      string // 会议uuid，meeting表的uuid
	UserUuid         string // 用户uuid，user表的uuid
	MeetingDatumUuid string // 议题uuid，meeting_datum表的uuid
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
}

// meetingDatumUserColumns holds the columns for table meeting_datum_user.
var meetingDatumUserColumns = MeetingDatumUserColumns{
	MeetingUuid:      "meeting_uuid",
	UserUuid:         "user_uuid",
	MeetingDatumUuid: "meeting_datum_uuid",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

// NewMeetingDatumUserDao creates and returns a new DAO object for table data access.
func NewMeetingDatumUserDao() *MeetingDatumUserDao {
	return &MeetingDatumUserDao{
		group:   "default",
		table:   "meeting_datum_user",
		columns: meetingDatumUserColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MeetingDatumUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MeetingDatumUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MeetingDatumUserDao) Columns() MeetingDatumUserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MeetingDatumUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MeetingDatumUserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MeetingDatumUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
