// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingSign is the golang structure of table meeting_sign for DAO operations like Where/Data.
type MeetingSign struct {
	g.Meta         `orm:"table:meeting_sign, do:true"`
	Id             interface{} //
	MeetingUuid    interface{} //
	SystemFileUuid interface{} //
	SignStartTime  *gtime.Time //
	Title          interface{} //
	SignType       interface{} //
	Status         interface{} //
	QrCode         interface{} //
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
}
