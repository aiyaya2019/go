// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingUserRights is the golang structure of table meeting_user_rights for DAO operations like Where/Data.
type MeetingUserRights struct {
	g.Meta      `orm:"table:meeting_user_rights, do:true"`
	Id          interface{} //
	MeetingUuid interface{} //
	UserUuid    interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
