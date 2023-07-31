// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ConfigThemeDao is the data access object for table config_theme.
type ConfigThemeDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns ConfigThemeColumns // columns contains all the column names of Table for convenient usage.
}

// ConfigThemeColumns defines and stores column names for table config_theme.
type ConfigThemeColumns struct {
	Id         string // 逻辑ID
	SysName    string // 系统名称
	ThemeName  string // 主题名称
	Path       string // 预览文件路径
	Logo       string // LOGO
	Background string // 背景图片
	Type       string // 类型:1=>cms,2=>c#,3=>qt,4=>ios
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

// configThemeColumns holds the columns for table config_theme.
var configThemeColumns = ConfigThemeColumns{
	Id:         "id",
	SysName:    "sys_name",
	ThemeName:  "theme_name",
	Path:       "path",
	Logo:       "logo",
	Background: "background",
	Type:       "type",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewConfigThemeDao creates and returns a new DAO object for table data access.
func NewConfigThemeDao() *ConfigThemeDao {
	return &ConfigThemeDao{
		group:   "default",
		table:   "config_theme",
		columns: configThemeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ConfigThemeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ConfigThemeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ConfigThemeDao) Columns() ConfigThemeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ConfigThemeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ConfigThemeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ConfigThemeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
