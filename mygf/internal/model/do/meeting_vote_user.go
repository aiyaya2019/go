// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingVoteUser is the golang structure of table meeting_vote_user for DAO operations like Where/Data.
type MeetingVoteUser struct {
	g.Meta          `orm:"table:meeting_vote_user, do:true"`
	MeetingUuid     interface{} // 会议uuid，meeting表的uuid
	UserUuid        interface{} // 用户uuid，user表的uuid
	MeetingVoteUuid interface{} // 投票uuid，meeting_vote表的uuid
	SystemFileUuid  interface{} // 文件uuid，system_file表的uuid
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 更新时间
}
