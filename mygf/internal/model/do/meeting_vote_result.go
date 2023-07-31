// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingVoteResult is the golang structure of table meeting_vote_result for DAO operations like Where/Data.
type MeetingVoteResult struct {
	g.Meta          `orm:"table:meeting_vote_result, do:true"`
	Id              interface{} // 逻辑ID
	UserUuid        interface{} // 用户uuid，user表的id
	MeetingVoteUuid interface{} // 投票uuid，meeting_vote表uuid
	VoteOptionUuid  interface{} // 选项uuid，vote_option表的uuid
	VoteTime        *gtime.Time // 投票时间
	Mark            interface{} // 备注
	Score           interface{} // 分值
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 更新时间
}
