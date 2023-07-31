// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserGroup is the golang structure of table user_group for DAO operations like Where/Data.
type UserGroup struct {
	g.Meta    `orm:"table:user_group, do:true"`
	Id        interface{} // 逻辑ID
	Name      interface{} // 名称
	Mark      interface{} // 备注
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
