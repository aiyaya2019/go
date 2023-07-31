// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingUserEnlist is the golang structure of table meeting_user_enlist for DAO operations like Where/Data.
type MeetingUserEnlist struct {
	g.Meta      `orm:"table:meeting_user_enlist, do:true"`
	Id          interface{} //
	MeetingUuid interface{} //
	UserUuid    interface{} //
	Sort        interface{} //
	Username    interface{} //
	Unit        interface{} //
	Dept        interface{} //
	Position    interface{} //
	Mobile      interface{} //
	IsFirst     interface{} //
	Mark        interface{} //
	CreateTime  *gtime.Time //
	Status      interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
