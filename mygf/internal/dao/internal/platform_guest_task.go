// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PlatformGuestTaskDao is the data access object for table platform_guest_task.
type PlatformGuestTaskDao struct {
	table   string                   // table is the underlying table name of the DAO.
	group   string                   // group is the database configuration group name of current DAO.
	columns PlatformGuestTaskColumns // columns contains all the column names of Table for convenient usage.
}

// PlatformGuestTaskColumns defines and stores column names for table platform_guest_task.
type PlatformGuestTaskColumns struct {
	Id               string // 主键
	Name             string // 会议名字
	MeetingStatus    string // 会议状态
	StartTime        string // 开始时间
	EndTime          string // 结束时间
	MeetingSlogan    string // 会议标语
	Nameplate        string // 参会人铭牌
	IsSecrect        string // 会议设置(0为普通会议，1为保密会议，2开放会议)
	Description      string // 会议描述
	ModeratorName    string // 主席名字无关联
	CreateName       string // 创建者名称
	MenuTab          string // 选择的菜单数值
	AuthorityJson    string // 客户端自定义菜单
	HostIp           string // 主会场ip
	HostPlatformUuid string // 主会场关联uuid，platform表的uuid
	HostPsIp         string // 主会场主ps的ip
	HostPsPort       string // 主会场主ps的port
	HostMeetingUuid  string // 主会场关联会议
	Status           string // 状态 1:未加入2:已加入3:加入失败4:主会场结束5:主会场删除6:被移除7:已处理8:开会中9:已退出10:已拒绝11:副平台结束
	Operate          string // 操作1add 2edit 3del
	CreatedAt        string // 创建时间
	UpdatedAt        string // 更新时间
}

// platformGuestTaskColumns holds the columns for table platform_guest_task.
var platformGuestTaskColumns = PlatformGuestTaskColumns{
	Id:               "id",
	Name:             "name",
	MeetingStatus:    "meeting_status",
	StartTime:        "start_time",
	EndTime:          "end_time",
	MeetingSlogan:    "meeting_slogan",
	Nameplate:        "nameplate",
	IsSecrect:        "is_secrect",
	Description:      "description",
	ModeratorName:    "moderator_name",
	CreateName:       "create_name",
	MenuTab:          "menu_tab",
	AuthorityJson:    "authority_json",
	HostIp:           "host_ip",
	HostPlatformUuid: "host_platform_uuid",
	HostPsIp:         "host_ps_ip",
	HostPsPort:       "host_ps_port",
	HostMeetingUuid:  "host_meeting_uuid",
	Status:           "status",
	Operate:          "operate",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
}

// NewPlatformGuestTaskDao creates and returns a new DAO object for table data access.
func NewPlatformGuestTaskDao() *PlatformGuestTaskDao {
	return &PlatformGuestTaskDao{
		group:   "default",
		table:   "platform_guest_task",
		columns: platformGuestTaskColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PlatformGuestTaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PlatformGuestTaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PlatformGuestTaskDao) Columns() PlatformGuestTaskColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PlatformGuestTaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PlatformGuestTaskDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PlatformGuestTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
