// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ServerDao is the data access object for table server.
type ServerDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns ServerColumns // columns contains all the column names of Table for convenient usage.
}

// ServerColumns defines and stores column names for table server.
type ServerColumns struct {
	Id        string // 逻辑id，服务器ID
	ServerId  string // 服务器ID
	Ip        string // IP地址
	Port      string // Port端口
	Status    string // 状态:0离线，1=>在线
	Flag      string // 服务器标志:0=>组，1=>中心后台服务器(apache)，2=>文件服务器，3=>流媒体服务器(nginx),4=>通讯服务器(workerman)，5=>数据服务器(mysql),6=>数据服务器2(redis),7=>ps,9=>转码，10=>大屏，14=>升降器服，15=>人脸识别服务，16=>语音转写，17=>电子桌牌
	Mark      string // 备注
	Pid       string // 父级id
	Operate   string // 操作状态：0=>无操作,1=>开机中，2=>关机中,3=>重启中，4=>操作成功，5=>操作失败
	Auth      string // 1=>开启、重启，2=>开启、重启、关机，3=>所有权限
	Deleted   string // 是否删除：0否；1删除
	Name      string // 名称
	Content   string // json格式，存配置
	IsHost    string // 是否是主服务器
	UvId      string // 关联upgrade_version版本号id
	UvName    string // 守护程序上报的版本号
	UvStatus  string // ps版本升级状态，1=>升级中 2=>成功 3=>失败
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// serverColumns holds the columns for table server.
var serverColumns = ServerColumns{
	Id:        "id",
	ServerId:  "server_id",
	Ip:        "ip",
	Port:      "port",
	Status:    "status",
	Flag:      "flag",
	Mark:      "mark",
	Pid:       "pid",
	Operate:   "operate",
	Auth:      "auth",
	Deleted:   "deleted",
	Name:      "name",
	Content:   "content",
	IsHost:    "is_host",
	UvId:      "uv_id",
	UvName:    "uv_name",
	UvStatus:  "uv_status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewServerDao creates and returns a new DAO object for table data access.
func NewServerDao() *ServerDao {
	return &ServerDao{
		group:   "default",
		table:   "server",
		columns: serverColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ServerDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ServerDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ServerDao) Columns() ServerColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ServerDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ServerDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ServerDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
