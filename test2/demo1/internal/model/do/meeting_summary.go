// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingSummary is the golang structure of table meeting_summary for DAO operations like Where/Data.
type MeetingSummary struct {
	g.Meta                   `orm:"table:meeting_summary, do:true"`
	Id                       interface{} //
	MeetingUuid              interface{} //
	SystemFileUuid           interface{} //
	UserUuid                 interface{} //
	Status                   interface{} //
	IsUpdate                 interface{} //
	FileFormatSystemFileUuid interface{} //
	MeetingDatumUuid         interface{} //
	CreatedAt                *gtime.Time //
	UpdatedAt                *gtime.Time //
}