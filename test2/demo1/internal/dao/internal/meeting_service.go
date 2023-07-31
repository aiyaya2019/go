// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingServiceDao is the data access object for table meeting_service.
type MeetingServiceDao struct {
	table   string                // table is the underlying table name of the DAO.
	group   string                // group is the database configuration group name of current DAO.
	columns MeetingServiceColumns // columns contains all the column names of Table for convenient usage.
}

// MeetingServiceColumns defines and stores column names for table meeting_service.
type MeetingServiceColumns struct {
	Id                  string //
	MeetingUuid         string //
	UserUuid            string //
	ServerId            string //
	Content             string //
	Time                string //
	IsConfirm           string //
	Mark                string //
	Status              string //
	LaterHandleUserUuid string //
	HandleUserUuid      string //
	LaterHandleTime     string //
	HandleTime          string //
	DeviceType          string //
	CreatedAt           string //
	UpdatedAt           string //
}

// meetingServiceColumns holds the columns for table meeting_service.
var meetingServiceColumns = MeetingServiceColumns{
	Id:                  "id",
	MeetingUuid:         "meeting_uuid",
	UserUuid:            "user_uuid",
	ServerId:            "server_id",
	Content:             "content",
	Time:                "time",
	IsConfirm:           "is_confirm",
	Mark:                "mark",
	Status:              "status",
	LaterHandleUserUuid: "later_handle_user_uuid",
	HandleUserUuid:      "handle_user_uuid",
	LaterHandleTime:     "later_handle_time",
	HandleTime:          "handle_time",
	DeviceType:          "device_type",
	CreatedAt:           "created_at",
	UpdatedAt:           "updated_at",
}

// NewMeetingServiceDao creates and returns a new DAO object for table data access.
func NewMeetingServiceDao() *MeetingServiceDao {
	return &MeetingServiceDao{
		group:   "default",
		table:   "meeting_service",
		columns: meetingServiceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MeetingServiceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MeetingServiceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MeetingServiceDao) Columns() MeetingServiceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MeetingServiceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MeetingServiceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MeetingServiceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
