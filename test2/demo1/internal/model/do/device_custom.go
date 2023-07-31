// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DeviceCustom is the golang structure of table device_custom for DAO operations like Where/Data.
type DeviceCustom struct {
	g.Meta       `orm:"table:device_custom, do:true"`
	Id           interface{} //
	IsShowDatum  interface{} //
	SkinColor    interface{} //
	BgImg        interface{} //
	Logo         interface{} //
	PropagateImg interface{} //
	Type         interface{} //
	CreatedAt    *gtime.Time //
	UpdatedAt    *gtime.Time //
}
