// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingVoteOption is the golang structure of table meeting_vote_option for DAO operations like Where/Data.
type MeetingVoteOption struct {
	g.Meta          `orm:"table:meeting_vote_option, do:true"`
	Id              interface{} // 逻辑ID
	Uuid            interface{} // 分布式uuid
	PlatformUuid    interface{} // 平台uuid
	MeetingVoteUuid interface{} // 投票uuid，meeting_vote表uuid
	Sort            interface{} // 权重
	VoteOptionName  interface{} // 投票选项名称
	CreatedAt       *gtime.Time // 创建时间
	UpdatedAt       *gtime.Time // 更新时间
}
