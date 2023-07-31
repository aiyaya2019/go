// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RoomSettingServe is the golang structure of table room_setting_serve for DAO operations like Where/Data.
type RoomSettingServe struct {
	g.Meta    `orm:"table:room_setting_serve, do:true"`
	Id        interface{} //
	RoomId    interface{} //
	Name      interface{} //
	Src       interface{} //
	IsSystem  interface{} //
	Sort      interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
