// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserOrganization is the golang structure of table user_organization for DAO operations like Where/Data.
type UserOrganization struct {
	g.Meta            `orm:"table:user_organization, do:true"`
	Id                interface{} //
	UserUuid          interface{} //
	OrganizationUuid  interface{} //
	OrganizationLevel interface{} //
	Default           interface{} //
	Binding           interface{} //
	Deleted           interface{} //
}
