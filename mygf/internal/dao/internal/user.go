// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDao is the data access object for table user.
type UserDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns UserColumns // columns contains all the column names of Table for convenient usage.
}

// UserColumns defines and stores column names for table user.
type UserColumns struct {
	Id                   string // 逻辑ID
	Uuid                 string // 用户分布式唯一ID
	PlatformUuid         string // 平台uuid
	Account              string // 账号:唯一
	Username             string // 名称
	Password             string // 密码
	Unit                 string // 单位
	Dept                 string // 部门
	Position             string // 职务
	Salutatory           string // 欢迎词
	IsValid              string // 是否有效:0=>无效，1=>有效
	IsTmp                string // 是否为临时用户：0=》否，1=》是,2=>测试用户,3=>级联平台用户
	UserType             string // 用户类型：0管理员，1秘书, 2普通用户, 3审计员（不能变换角色）, 10超级管理员
	Mark                 string // 备注
	Theme                string // 主题
	Headimg              string // 头像
	Mobile               string // 手机号
	Email                string // 邮箱
	FaceId               string // 人脸id
	Sort                 string // 排序,数值越大排名越前
	SummitMeetingManager string // 管理高级会议0=>不能
	UserLogType          string // 1后勤人员(用于登录会议服务APP）
	Deleted              string // 软删除：0未删除1已删除
	ChangePassword       string // 管理员修改密码：0未修改1修改
	CreatedAt            string // 创建时间
	UpdatedAt            string // 更新时间
	LastPasswdUpdateTime string // 最后一次密码修改时间
	Status               string // 状态 0：正常 1：锁定
}

// userColumns holds the columns for table user.
var userColumns = UserColumns{
	Id:                   "id",
	Uuid:                 "uuid",
	PlatformUuid:         "platform_uuid",
	Account:              "account",
	Username:             "username",
	Password:             "password",
	Unit:                 "unit",
	Dept:                 "dept",
	Position:             "position",
	Salutatory:           "salutatory",
	IsValid:              "is_valid",
	IsTmp:                "is_tmp",
	UserType:             "user_type",
	Mark:                 "mark",
	Theme:                "theme",
	Headimg:              "headimg",
	Mobile:               "mobile",
	Email:                "email",
	FaceId:               "face_id",
	Sort:                 "sort",
	SummitMeetingManager: "summit_meeting_manager",
	UserLogType:          "user_log_type",
	Deleted:              "deleted",
	ChangePassword:       "change_password",
	CreatedAt:            "created_at",
	UpdatedAt:            "updated_at",
	LastPasswdUpdateTime: "last_passwd_update_time",
	Status:               "status",
}

// NewUserDao creates and returns a new DAO object for table data access.
func NewUserDao() *UserDao {
	return &UserDao{
		group:   "default",
		table:   "user",
		columns: userColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserDao) Columns() UserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
