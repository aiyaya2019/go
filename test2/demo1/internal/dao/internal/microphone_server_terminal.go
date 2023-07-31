// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MicrophoneServerTerminalDao is the data access object for table microphone_server_terminal.
type MicrophoneServerTerminalDao struct {
	table   string                          // table is the underlying table name of the DAO.
	group   string                          // group is the database configuration group name of current DAO.
	columns MicrophoneServerTerminalColumns // columns contains all the column names of Table for convenient usage.
}

// MicrophoneServerTerminalColumns defines and stores column names for table microphone_server_terminal.
type MicrophoneServerTerminalColumns struct {
	Id                 string //
	MicId              string //
	Type               string //
	Status             string //
	MicrophoneServerId string //
	Mark               string //
	Mac                string //
	ScreenIp           string //
	DataType           string // 1:普通话筒，2:8512
	CreatedAt          string //
	UpdatedAt          string //
}

// microphoneServerTerminalColumns holds the columns for table microphone_server_terminal.
var microphoneServerTerminalColumns = MicrophoneServerTerminalColumns{
	Id:                 "id",
	MicId:              "mic_id",
	Type:               "type",
	Status:             "status",
	MicrophoneServerId: "microphone_server_id",
	Mark:               "mark",
	Mac:                "mac",
	ScreenIp:           "screen_ip",
	DataType:           "data_type",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
}

// NewMicrophoneServerTerminalDao creates and returns a new DAO object for table data access.
func NewMicrophoneServerTerminalDao() *MicrophoneServerTerminalDao {
	return &MicrophoneServerTerminalDao{
		group:   "default",
		table:   "microphone_server_terminal",
		columns: microphoneServerTerminalColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MicrophoneServerTerminalDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MicrophoneServerTerminalDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MicrophoneServerTerminalDao) Columns() MicrophoneServerTerminalColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MicrophoneServerTerminalDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MicrophoneServerTerminalDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MicrophoneServerTerminalDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
