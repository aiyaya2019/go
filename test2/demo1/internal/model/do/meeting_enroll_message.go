// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingEnrollMessage is the golang structure of table meeting_enroll_message for DAO operations like Where/Data.
type MeetingEnrollMessage struct {
	g.Meta      `orm:"table:meeting_enroll_message, do:true"`
	Id          interface{} //
	MeetingUuid interface{} //
	Contact     interface{} //
	Link        interface{} //
	TimeType    interface{} //
	StartTime   interface{} //
	EndTime     interface{} //
	QrcodePath  interface{} //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}