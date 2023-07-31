// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OrganizationDao is the data access object for table organization.
type OrganizationDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns OrganizationColumns // columns contains all the column names of Table for convenient usage.
}

// OrganizationColumns defines and stores column names for table organization.
type OrganizationColumns struct {
	Id           string // 自增ID
	Uuid         string // 组织架构分布式唯一ID
	Puuid        string // 父uuid
	PlatformUuid string // 平台uuid
	Level        string // 层级(1：单位，2：部门，3：职位)
	Name         string // 名称
	Sort         string // 排序（数值越小，排序越前）
	IsTmp        string // 是否为临时：0=》否，1=》是
	Status       string // 状态：1正常；2禁用
	Deleted      string // 0不删除；1删除
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
}

// organizationColumns holds the columns for table organization.
var organizationColumns = OrganizationColumns{
	Id:           "id",
	Uuid:         "uuid",
	Puuid:        "puuid",
	PlatformUuid: "platform_uuid",
	Level:        "level",
	Name:         "name",
	Sort:         "sort",
	IsTmp:        "is_tmp",
	Status:       "status",
	Deleted:      "deleted",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewOrganizationDao creates and returns a new DAO object for table data access.
func NewOrganizationDao() *OrganizationDao {
	return &OrganizationDao{
		group:   "default",
		table:   "organization",
		columns: organizationColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *OrganizationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *OrganizationDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *OrganizationDao) Columns() OrganizationColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *OrganizationDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *OrganizationDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *OrganizationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
