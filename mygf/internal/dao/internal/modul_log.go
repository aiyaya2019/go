// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ModulLogDao is the data access object for table modul_log.
type ModulLogDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns ModulLogColumns // columns contains all the column names of Table for convenient usage.
}

// ModulLogColumns defines and stores column names for table modul_log.
type ModulLogColumns struct {
	Id        string // 主键id
	ModulId   string // server的id/terminal的id
	Type      string // 类型标志:0=>组，1=>中心后台服务器(apache)，2=>文件服务器，3=>流媒体服务器(nginx),4=>通讯服务器(workerman)，5=>数据服务器(mysql),6=>数据服务器2(redis),7=>ps,9=>转码，10=>大屏，14=>升降器服，15=>人脸识别服务，11=>客户端
	UserUuid  string // 操作用户uuid，user表的uuid
	Url       string // 下载路径
	Path      string // 存放路径
	Size      string // 大小
	Status    string // 转码状态，0=》未开始，1=》转换中，2=》成功，3=》失败
	Time      string // 时间
	StartTime string // 开始时间
	EndTime   string // 结束时间
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// modulLogColumns holds the columns for table modul_log.
var modulLogColumns = ModulLogColumns{
	Id:        "id",
	ModulId:   "modul_id",
	Type:      "type",
	UserUuid:  "user_uuid",
	Url:       "url",
	Path:      "path",
	Size:      "size",
	Status:    "status",
	Time:      "time",
	StartTime: "start_time",
	EndTime:   "end_time",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewModulLogDao creates and returns a new DAO object for table data access.
func NewModulLogDao() *ModulLogDao {
	return &ModulLogDao{
		group:   "default",
		table:   "modul_log",
		columns: modulLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ModulLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ModulLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ModulLogDao) Columns() ModulLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ModulLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ModulLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ModulLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
