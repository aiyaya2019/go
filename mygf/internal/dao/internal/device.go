// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DeviceDao is the data access object for table device.
type DeviceDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns DeviceColumns // columns contains all the column names of Table for convenient usage.
}

// DeviceColumns defines and stores column names for table device.
type DeviceColumns struct {
	Id        string //
	DeviceIp  string // 设备ip
	DeviceMac string // 设备mac
	RoomId    string // 会议室id
	Status    string // 设备状态:0空闲；1使用中；2设备离线
	IsShowing string // 当前是否显示图文：0否；1是
	IsPush    string // 是否允许推送：0否；1是
	Type      string // 设备类型：1信息屏
	Deleted   string // 删除：0未删除；1删除
	Ip        string // 服务器ip
	Port      string // 服务器端口
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// deviceColumns holds the columns for table device.
var deviceColumns = DeviceColumns{
	Id:        "id",
	DeviceIp:  "device_ip",
	DeviceMac: "device_mac",
	RoomId:    "room_id",
	Status:    "status",
	IsShowing: "is_showing",
	IsPush:    "is_push",
	Type:      "type",
	Deleted:   "deleted",
	Ip:        "ip",
	Port:      "port",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewDeviceDao creates and returns a new DAO object for table data access.
func NewDeviceDao() *DeviceDao {
	return &DeviceDao{
		group:   "default",
		table:   "device",
		columns: deviceColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DeviceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DeviceDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DeviceDao) Columns() DeviceColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DeviceDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DeviceDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DeviceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
