// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DeviceCustomDao is the data access object for table device_custom.
type DeviceCustomDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns DeviceCustomColumns // columns contains all the column names of Table for convenient usage.
}

// DeviceCustomColumns defines and stores column names for table device_custom.
type DeviceCustomColumns struct {
	Id           string //
	IsShowDatum  string // 是否显示议题：0不显示；1显示
	SkinColor    string // 1浅蓝；2深蓝
	BgImg        string // 背景皮肤
	Logo         string // 空闲页logo
	PropagateImg string // 空闲页宣传图
	Type         string // 设备类型：1信息屏
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
}

// deviceCustomColumns holds the columns for table device_custom.
var deviceCustomColumns = DeviceCustomColumns{
	Id:           "id",
	IsShowDatum:  "is_show_datum",
	SkinColor:    "skin_color",
	BgImg:        "bg_img",
	Logo:         "logo",
	PropagateImg: "propagate_img",
	Type:         "type",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewDeviceCustomDao creates and returns a new DAO object for table data access.
func NewDeviceCustomDao() *DeviceCustomDao {
	return &DeviceCustomDao{
		group:   "default",
		table:   "device_custom",
		columns: deviceCustomColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *DeviceCustomDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *DeviceCustomDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *DeviceCustomDao) Columns() DeviceCustomColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *DeviceCustomDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *DeviceCustomDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *DeviceCustomDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
