// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MicrophoneServerDao is the data access object for table microphone_server.
type MicrophoneServerDao struct {
	table   string                  // table is the underlying table name of the DAO.
	group   string                  // group is the database configuration group name of current DAO.
	columns MicrophoneServerColumns // columns contains all the column names of Table for convenient usage.
}

// MicrophoneServerColumns defines and stores column names for table microphone_server.
type MicrophoneServerColumns struct {
	Id          string // 逻辑id
	RoomId      string // 会议室id,room表的id
	Ip          string // ip
	Port        string // 端口
	MicMax      string // 话筒最大打开数量 1,2,4,8
	MicType     string // 话筒模式 1.FIFO 2.NORMAL 3.VOIC 4.APPLY
	Volume      string // 总音量
	Alt         string // 高音
	Bass        string // 低音
	Sensitivity string // 灵敏度
	ClosingTime string // 关闭时间
	Mark        string // 备注
	HostSn      string // 注册码
	Status      string // 1.正常  0.连接不上
	Register    string // 1.已注册 0.没注册
	DeviceType  string // 1. W100 2. 0200MC
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// microphoneServerColumns holds the columns for table microphone_server.
var microphoneServerColumns = MicrophoneServerColumns{
	Id:          "id",
	RoomId:      "room_id",
	Ip:          "ip",
	Port:        "port",
	MicMax:      "mic_max",
	MicType:     "mic_type",
	Volume:      "volume",
	Alt:         "alt",
	Bass:        "bass",
	Sensitivity: "sensitivity",
	ClosingTime: "closing_time",
	Mark:        "mark",
	HostSn:      "host_sn",
	Status:      "status",
	Register:    "register",
	DeviceType:  "device_type",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewMicrophoneServerDao creates and returns a new DAO object for table data access.
func NewMicrophoneServerDao() *MicrophoneServerDao {
	return &MicrophoneServerDao{
		group:   "default",
		table:   "microphone_server",
		columns: microphoneServerColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MicrophoneServerDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MicrophoneServerDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MicrophoneServerDao) Columns() MicrophoneServerColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MicrophoneServerDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MicrophoneServerDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MicrophoneServerDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
