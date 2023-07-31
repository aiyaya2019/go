// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserOrganization is the golang structure of table user_organization for DAO operations like Where/Data.
type UserOrganization struct {
	g.Meta            `orm:"table:user_organization, do:true"`
	Id                interface{} //
	UserUuid          interface{} // 用户uuid，user表uuid
	OrganizationUuid  interface{} // 组织架构uuid，organization表uuid
	OrganizationLevel interface{} // 层级
	Default           interface{} // 是否默认：0非默认；1默认
	Binding           interface{} // 0虚拟绑定的上级(统计人数用)；1真实绑定所在组织
	Deleted           interface{} // 软删除：0未删除1已删除
}
