// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingUserRights is the golang structure of table meeting_user_rights for DAO operations like Where/Data.
type MeetingUserRights struct {
	g.Meta      `orm:"table:meeting_user_rights, do:true"`
	Id          interface{} // 逻辑id
	MeetingUuid interface{} // 会议uuid，meeting表的uuid
	UserUuid    interface{} // 用户uuid，user表的uuid
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
