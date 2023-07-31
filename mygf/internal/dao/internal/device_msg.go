// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DeviceMsgDao is the data access object for table device_msg.
type DeviceMsgDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns DeviceMsgColumns // columns contains all the column names of Table for convenient usage.
}

// DeviceMsgColumns defines and stores column names for table device_msg.
type DeviceMsgColumns struct {
	Id        string //
	DeviceId  string // 设备id，device表主键id
	Msg       string // 消息内容
	Time      string // 消息播报时间/秒
	Type      string // 设备类型：1信息屏
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// deviceMsgColumns holds the columns for table device_msg.
var deviceMsgColumns = DeviceMsgColumns{
	Id:        "id",
	DeviceId:  "device_id",
	Msg:       "msg",
	Time:      "time",
	Type:      "type",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewDeviceMsgDao creates and returns a new DAO object for table data access.
func NewDeviceMsgDao() *DeviceMsgDao {
	return &DeviceMsgDao{
		group:   "default",
		table:   "device_msg",
		columns: deviceMsgColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DeviceMsgDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DeviceMsgDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DeviceMsgDao) Columns() DeviceMsgColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DeviceMsgDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DeviceMsgDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DeviceMsgDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
