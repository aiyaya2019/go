// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingDatumUser is the golang structure of table meeting_datum_user for DAO operations like Where/Data.
type MeetingDatumUser struct {
	g.Meta           `orm:"table:meeting_datum_user, do:true"`
	MeetingUuid      interface{} // 会议uuid，meeting表的uuid
	UserUuid         interface{} // 用户uuid，user表的uuid
	MeetingDatumUuid interface{} // 议题uuid，meeting_datum表的uuid
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
}
