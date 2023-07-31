// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// TerminalDao is the data access object for table terminal.
type TerminalDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns TerminalColumns // columns contains all the column names of Table for convenient usage.
}

// TerminalColumns defines and stores column names for table terminal.
type TerminalColumns struct {
	Id                      string // 逻辑id
	TId                     string // 终端id
	Name                    string // 终端名称
	Ip                      string // 终端IP
	Mac                     string // 终端 mac
	Status                  string // 状态:0=>离线，1=>在线
	RoomId                  string // 会议室ID，room表的id
	Mark                    string // 备注
	Type                    string // 1.pc客户端2.移动客户端4.解码器大屏8.升降器终端,9三代升降器,10二代升降器,关联二代8终端
	Config                  string // 单个终端配置
	ConfigStatus            string // 终端下发配置，0=>默认1=>成功2=>失败
	AssociatedVersionId     string // 关联版本id
	AssociatedVersionName   string // 版本名称
	AssociatedVersionStatus string // 升级状态，1=>升级中 2=>成功 3=>失败
	Operate                 string // 操作状态：0=>无操作,1=>开机中，2=>关机中,3=>重启中，4=>操作成功，5=>操作失败，6=>操作异常
	OsType                  string // 1：windows；2：android; 3: ios;4: kylin； 5：centos;
	GuardConfig             string // 守护程序配置
	GuardUvId               string // 守护程序版本id
	GuardUvName             string // 守护程序版本名称
	GuardUvStatus           string // 守护程序版本状态
	DeviceSn                string // 序列号
	DeviceType              string // 型号
	IpType                  string // IP类型0动态，1静态
	Mask                    string // 子网掩码
	Gateway                 string // 默认网关
	HostIp                  string // 主机ip
	ClientIp                string // 客户端ip
	OptCounts               string // 上升下降次数
	Concurrence             string // 并发用户名
	LifterType              string // 终端升降器类型2代3代
	SecondTid               string // type10 二代升降器关联type8的id
	CreatedAt               string // 创建时间
	UpdatedAt               string // 更新时间
}

// terminalColumns holds the columns for table terminal.
var terminalColumns = TerminalColumns{
	Id:                      "id",
	TId:                     "t_id",
	Name:                    "name",
	Ip:                      "ip",
	Mac:                     "mac",
	Status:                  "status",
	RoomId:                  "room_id",
	Mark:                    "mark",
	Type:                    "type",
	Config:                  "config",
	ConfigStatus:            "config_status",
	AssociatedVersionId:     "associated_version_id",
	AssociatedVersionName:   "associated_version_name",
	AssociatedVersionStatus: "associated_version_status",
	Operate:                 "operate",
	OsType:                  "os_type",
	GuardConfig:             "guard_config",
	GuardUvId:               "guard_uv_id",
	GuardUvName:             "guard_uv_name",
	GuardUvStatus:           "guard_uv_status",
	DeviceSn:                "device_sn",
	DeviceType:              "device_type",
	IpType:                  "ip_type",
	Mask:                    "mask",
	Gateway:                 "gateway",
	HostIp:                  "host_ip",
	ClientIp:                "client_ip",
	OptCounts:               "opt_counts",
	Concurrence:             "concurrence",
	LifterType:              "lifter_type",
	SecondTid:               "second_tid",
	CreatedAt:               "created_at",
	UpdatedAt:               "updated_at",
}

// NewTerminalDao creates and returns a new DAO object for table data access.
func NewTerminalDao() *TerminalDao {
	return &TerminalDao{
		group:   "default",
		table:   "terminal",
		columns: terminalColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *TerminalDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *TerminalDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *TerminalDao) Columns() TerminalColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *TerminalDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *TerminalDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *TerminalDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
