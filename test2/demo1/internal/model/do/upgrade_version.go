// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UpgradeVersion is the golang structure of table upgrade_version for DAO operations like Where/Data.
type UpgradeVersion struct {
	g.Meta     `orm:"table:upgrade_version, do:true"`
	Id         interface{} //
	UserUuid   interface{} //
	Name       interface{} //
	Versions   interface{} //
	Url        interface{} //
	Type       interface{} //
	CurrentUse interface{} //
	Mark       interface{} //
	Deleted    interface{} //
	Status     interface{} //
	OsType     interface{} //
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
}
