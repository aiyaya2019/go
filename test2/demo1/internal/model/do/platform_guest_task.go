// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlatformGuestTask is the golang structure of table platform_guest_task for DAO operations like Where/Data.
type PlatformGuestTask struct {
	g.Meta           `orm:"table:platform_guest_task, do:true"`
	Id               interface{} //
	Name             interface{} //
	MeetingStatus    interface{} //
	StartTime        *gtime.Time //
	EndTime          *gtime.Time //
	MeetingSlogan    interface{} //
	Nameplate        interface{} //
	IsSecrect        interface{} //
	Description      interface{} //
	ModeratorName    interface{} //
	CreateName       interface{} //
	MenuTab          interface{} //
	AuthorityJson    interface{} //
	HostIp           interface{} //
	HostPlatformUuid interface{} //
	HostPsIp         interface{} //
	HostPsPort       interface{} //
	HostMeetingUuid  interface{} //
	Status           interface{} //
	Operate          interface{} //
	CreatedAt        *gtime.Time //
	UpdatedAt        *gtime.Time //
}