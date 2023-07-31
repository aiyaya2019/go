// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PhoneticTranscriptionLogDao is the data access object for table phonetic_transcription_log.
type PhoneticTranscriptionLogDao struct {
	table   string                          // table is the underlying table name of the DAO.
	group   string                          // group is the database configuration group name of current DAO.
	columns PhoneticTranscriptionLogColumns // columns contains all the column names of Table for convenient usage.
}

// PhoneticTranscriptionLogColumns defines and stores column names for table phonetic_transcription_log.
type PhoneticTranscriptionLogColumns struct {
	Id        string //
	Type      string // 1=>手动同步，2=》自动同步
	Status    string // 1=》同步成功，2=》同步失败
	Detailed  string // 同步详情
	Creator   string // 创建者
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// phoneticTranscriptionLogColumns holds the columns for table phonetic_transcription_log.
var phoneticTranscriptionLogColumns = PhoneticTranscriptionLogColumns{
	Id:        "id",
	Type:      "type",
	Status:    "status",
	Detailed:  "detailed",
	Creator:   "creator",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewPhoneticTranscriptionLogDao creates and returns a new DAO object for table data access.
func NewPhoneticTranscriptionLogDao() *PhoneticTranscriptionLogDao {
	return &PhoneticTranscriptionLogDao{
		group:   "default",
		table:   "phonetic_transcription_log",
		columns: phoneticTranscriptionLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PhoneticTranscriptionLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PhoneticTranscriptionLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PhoneticTranscriptionLogDao) Columns() PhoneticTranscriptionLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PhoneticTranscriptionLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PhoneticTranscriptionLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PhoneticTranscriptionLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
