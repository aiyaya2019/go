// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta               `orm:"table:user, do:true"`
	Id                   interface{} // 逻辑ID
	Uuid                 interface{} // 用户分布式唯一ID
	PlatformUuid         interface{} // 平台uuid
	Account              interface{} // 账号:唯一
	Username             interface{} // 名称
	Password             interface{} // 密码
	Unit                 interface{} // 单位
	Dept                 interface{} // 部门
	Position             interface{} // 职务
	Salutatory           interface{} // 欢迎词
	IsValid              interface{} // 是否有效:0=>无效，1=>有效
	IsTmp                interface{} // 是否为临时用户：0=》否，1=》是,2=>测试用户,3=>级联平台用户
	UserType             interface{} // 用户类型：0管理员，1秘书, 2普通用户, 3审计员（不能变换角色）, 10超级管理员
	Mark                 interface{} // 备注
	Theme                interface{} // 主题
	Headimg              interface{} // 头像
	Mobile               interface{} // 手机号
	Email                interface{} // 邮箱
	FaceId               interface{} // 人脸id
	Sort                 interface{} // 排序,数值越大排名越前
	SummitMeetingManager interface{} // 管理高级会议0=>不能
	UserLogType          interface{} // 1后勤人员(用于登录会议服务APP）
	Deleted              interface{} // 软删除：0未删除1已删除
	ChangePassword       interface{} // 管理员修改密码：0未修改1修改
	CreatedAt            *gtime.Time // 创建时间
	UpdatedAt            *gtime.Time // 更新时间
	LastPasswdUpdateTime *gtime.Time // 最后一次密码修改时间
	Status               interface{} // 状态 0：正常 1：锁定
}
