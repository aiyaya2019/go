// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoomLayoutDao is the data access object for table room_layout.
type RoomLayoutDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns RoomLayoutColumns // columns contains all the column names of Table for convenient usage.
}

// RoomLayoutColumns defines and stores column names for table room_layout.
type RoomLayoutColumns struct {
	Id        string //
	RoomId    string // 会议室id
	StartTime string // 开始推送时间
	EndTime   string // 结束推送时间
	PushType  string // 推送类型：0单次；1每天；2每周一；3每周二；4每周三；5每周四；6每周五；7每周六；8每周日
	PushState string // 允许推送的设备状态：0空闲；1使用中；2空闲和使用中
	Status    string // 0不推送；1推送
	Layout    string // 空闲页布局
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// roomLayoutColumns holds the columns for table room_layout.
var roomLayoutColumns = RoomLayoutColumns{
	Id:        "id",
	RoomId:    "room_id",
	StartTime: "start_time",
	EndTime:   "end_time",
	PushType:  "push_type",
	PushState: "push_state",
	Status:    "status",
	Layout:    "layout",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewRoomLayoutDao creates and returns a new DAO object for table data access.
func NewRoomLayoutDao() *RoomLayoutDao {
	return &RoomLayoutDao{
		group:   "default",
		table:   "room_layout",
		columns: roomLayoutColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoomLayoutDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoomLayoutDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoomLayoutDao) Columns() RoomLayoutColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoomLayoutDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoomLayoutDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoomLayoutDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
