// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id                   uint        `json:"id"                   ` // 逻辑ID
	Uuid                 uint        `json:"uuid"                 ` // 用户分布式唯一ID
	PlatformUuid         uint        `json:"platformUuid"         ` // 平台uuid
	Account              string      `json:"account"              ` // 账号:唯一
	Username             string      `json:"username"             ` // 名称
	Password             string      `json:"password"             ` // 密码
	Unit                 string      `json:"unit"                 ` // 单位
	Dept                 string      `json:"dept"                 ` // 部门
	Position             string      `json:"position"             ` // 职务
	Salutatory           string      `json:"salutatory"           ` // 欢迎词
	IsValid              uint        `json:"isValid"              ` // 是否有效:0=>无效，1=>有效
	IsTmp                int         `json:"isTmp"                ` // 是否为临时用户：0=》否，1=》是,2=>测试用户,3=>级联平台用户
	UserType             uint        `json:"userType"             ` // 用户类型：0管理员，1秘书, 2普通用户, 3审计员（不能变换角色）, 10超级管理员
	Mark                 string      `json:"mark"                 ` // 备注
	Theme                uint        `json:"theme"                ` // 主题
	Headimg              string      `json:"headimg"              ` // 头像
	Mobile               string      `json:"mobile"               ` // 手机号
	Email                string      `json:"email"                ` // 邮箱
	FaceId               int         `json:"faceId"               ` // 人脸id
	Sort                 int         `json:"sort"                 ` // 排序,数值越大排名越前
	SummitMeetingManager int         `json:"summitMeetingManager" ` // 管理高级会议0=>不能
	UserLogType          int         `json:"userLogType"          ` // 1后勤人员(用于登录会议服务APP）
	Deleted              int         `json:"deleted"              ` // 软删除：0未删除1已删除
	ChangePassword       int         `json:"changePassword"       ` // 管理员修改密码：0未修改1修改
	CreatedAt            *gtime.Time `json:"createdAt"            ` // 创建时间
	UpdatedAt            *gtime.Time `json:"updatedAt"            ` // 更新时间
	LastPasswdUpdateTime *gtime.Time `json:"lastPasswdUpdateTime" ` // 最后一次密码修改时间
	Status               int         `json:"status"               ` // 状态 0：正常 1：锁定
}
