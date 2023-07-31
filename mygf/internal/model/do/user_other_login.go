// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserOtherLogin is the golang structure of table user_other_login for DAO operations like Where/Data.
type UserOtherLogin struct {
	g.Meta    `orm:"table:user_other_login, do:true"`
	Id        interface{} // 逻辑ID
	UserUuid  interface{} // 用户uuid，user表的uuid
	Name      interface{} // 名称
	Pass      interface{} // 特征码
	Type      interface{} // 1:指纹
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
