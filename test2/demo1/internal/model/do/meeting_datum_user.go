// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingDatumUser is the golang structure of table meeting_datum_user for DAO operations like Where/Data.
type MeetingDatumUser struct {
	g.Meta           `orm:"table:meeting_datum_user, do:true"`
	MeetingUuid      interface{} //
	UserUuid         interface{} //
	MeetingDatumUuid interface{} //
	CreatedAt        *gtime.Time //
	UpdatedAt        *gtime.Time //
}
