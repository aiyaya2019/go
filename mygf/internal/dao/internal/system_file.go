// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SystemFileDao is the data access object for table system_file.
type SystemFileDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns SystemFileColumns // columns contains all the column names of Table for convenient usage.
}

// SystemFileColumns defines and stores column names for table system_file.
type SystemFileColumns struct {
	Id            string // 逻辑ID
	ParentUuid    string // 父级uuid
	Uuid          string // 文件分布式id
	PlatformUuid  string // 平台uuid
	MeetingUuid   string // 会议uuid，meeting表的uuid
	UserUuid      string // 用户uuid，user表的uuid
	Filename      string // 文件名
	Filesize      string // 文件大小
	Filepath      string // 文件路径
	TerminalId    string // (t_id改为：terminal_id)终端ID：上传者终端ID 会控客户上传填 0，info_terminal（改名为terminal）表的id
	UploadTime    string // 上传时间
	Sort          string // 权重
	FileUse       string // 文件用途:1=>议程文件，2=>议题文件，3=>纪要文件,4=>临时文件,5=>批注文件,6=>签到文件,7=>电子白板,8=>直播，9=>点播,10=>决定事项,20=>会议纪要,21=>纪要会签图片,22=>投票,23=>评分,24=>签到
	Mark          string //
	IsDirectory   string // 是否目录（0否1是）
	AllUser       string // 是否选所有用户
	StartDatum    string // 0议题开启前显示,1议题开启后显示
	AllowDownload string // 允许下载0不允许1允许
	IsSecret      string // 是否保密0:不保密,1保密
	Status        string // 0失效
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
}

// systemFileColumns holds the columns for table system_file.
var systemFileColumns = SystemFileColumns{
	Id:            "id",
	ParentUuid:    "parent_uuid",
	Uuid:          "uuid",
	PlatformUuid:  "platform_uuid",
	MeetingUuid:   "meeting_uuid",
	UserUuid:      "user_uuid",
	Filename:      "filename",
	Filesize:      "filesize",
	Filepath:      "filepath",
	TerminalId:    "terminal_id",
	UploadTime:    "upload_time",
	Sort:          "sort",
	FileUse:       "file_use",
	Mark:          "mark",
	IsDirectory:   "is_directory",
	AllUser:       "all_user",
	StartDatum:    "start_datum",
	AllowDownload: "allow_download",
	IsSecret:      "is_secret",
	Status:        "status",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewSystemFileDao creates and returns a new DAO object for table data access.
func NewSystemFileDao() *SystemFileDao {
	return &SystemFileDao{
		group:   "default",
		table:   "system_file",
		columns: systemFileColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SystemFileDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SystemFileDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SystemFileDao) Columns() SystemFileColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SystemFileDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SystemFileDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SystemFileDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
