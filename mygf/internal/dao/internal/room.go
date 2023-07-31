// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// RoomDao is the data access object for table room.
type RoomDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns RoomColumns // columns contains all the column names of Table for convenient usage.
}

// RoomColumns defines and stores column names for table room.
type RoomColumns struct {
	Id             string // 逻辑id,会议室ID
	Name           string // 会议室名称
	ServerId       string // 服务器ID，server表的id
	Roomlayout     string // 会议室布局json(单个)
	Status         string // 状态:0=>未使用，1=>使用中
	Mark           string // 备注
	Roombackground string // 会议室布局背景
	BgWidth        string // 背景图宽
	BgHeight       string // 背景图高
	TerminalConfig string // 终端配置json格式
	Transfer       string // 0不启用语音转写，1启用
	Electronic     string // 0不启用电子桌牌，1启用
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
	TransferRoomId string //
}

// roomColumns holds the columns for table room.
var roomColumns = RoomColumns{
	Id:             "id",
	Name:           "name",
	ServerId:       "server_id",
	Roomlayout:     "roomlayout",
	Status:         "status",
	Mark:           "mark",
	Roombackground: "roombackground",
	BgWidth:        "bg_width",
	BgHeight:       "bg_height",
	TerminalConfig: "terminal_config",
	Transfer:       "transfer",
	Electronic:     "electronic",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	TransferRoomId: "transfer_room_id",
}

// NewRoomDao creates and returns a new DAO object for table data access.
func NewRoomDao() *RoomDao {
	return &RoomDao{
		group:   "default",
		table:   "room",
		columns: roomColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *RoomDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *RoomDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *RoomDao) Columns() RoomColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *RoomDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *RoomDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *RoomDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
