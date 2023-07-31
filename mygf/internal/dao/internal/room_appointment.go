// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoomAppointmentDao is the data access object for table room_appointment.
type RoomAppointmentDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns RoomAppointmentColumns // columns contains all the column names of Table for convenient usage.
}

// RoomAppointmentColumns defines and stores column names for table room_appointment.
type RoomAppointmentColumns struct {
	Id          string // 逻辑ID
	MeetingUuid string // 会议uuid，meeting表的uuid
	UserUuid    string // 预约用户的uuid，user表的uuid
	RoomId      string // 会议室ID,room表id
	Name        string // 预约人员的名字
	StartTime   string // 预约的开始时间
	EndTime     string // 预约的结束时间
	Status      string // 状态，1正常，0停用
	Note        string // 备注
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// roomAppointmentColumns holds the columns for table room_appointment.
var roomAppointmentColumns = RoomAppointmentColumns{
	Id:          "id",
	MeetingUuid: "meeting_uuid",
	UserUuid:    "user_uuid",
	RoomId:      "room_id",
	Name:        "name",
	StartTime:   "start_time",
	EndTime:     "end_time",
	Status:      "status",
	Note:        "note",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewRoomAppointmentDao creates and returns a new DAO object for table data access.
func NewRoomAppointmentDao() *RoomAppointmentDao {
	return &RoomAppointmentDao{
		group:   "default",
		table:   "room_appointment",
		columns: roomAppointmentColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoomAppointmentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoomAppointmentDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoomAppointmentDao) Columns() RoomAppointmentColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoomAppointmentDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoomAppointmentDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoomAppointmentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
