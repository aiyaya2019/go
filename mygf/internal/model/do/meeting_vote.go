// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingVote is the golang structure of table meeting_vote for DAO operations like Where/Data.
type MeetingVote struct {
	g.Meta               `orm:"table:meeting_vote, do:true"`
	Id                   interface{} // 逻辑ID
	Uuid                 interface{} // 分布式uuid
	PlatformUuid         interface{} // 平台uuid，platform表的uuid
	MeetingUuid          interface{} // 会议uuid，meeting表的uuid
	MeetingDatumUuid     interface{} // 议题uuid，meeting_datum表的uuid
	SystemFileUuid       interface{} // 关联文件uuid，system_file表的uuid
	VoteTitle            interface{} // 投票标题
	VoteName             interface{} // 投票内容
	StartTime            *gtime.Time // 开始时间
	EndTime              *gtime.Time // 结束时间
	Countdown            interface{} // 倒计时
	VoteType             interface{} // 投票类型:0=>自定义；1=>表决
	IsMultiple           interface{} // 是否多选:0=>否，1=>是
	IsTruename           interface{} // 是否实名:0=>否，1=>是
	Status               interface{} // 状态：0=》未启用，1=》启用，2=》 已结束
	Mark                 interface{} // 备注
	IsSign               interface{} // 是否签到
	IsPassRate           interface{} // 是否要通过率
	PassRate             interface{} // 通过率
	IsVoteTimeLimit      interface{} // 是否限时投票
	VoteTimeLimit        interface{} // 投票限时时长
	IsMark               interface{} // 是否需要备注(1为是，0为否)
	ScreenMode           interface{} // 投屏方式0文字,1柱状图,2饼状图
	VotingAuthentication interface{} // 认证方式:0不使用 1签名验证 2人脸识别 3指纹验证
	VotingForce          interface{} // 是否强制认证，0否，1是
	ScoreScore           interface{} // 评分分值，“-”分割
	ScoreType            interface{} // 类型：0 投票 ； 1 评分
	IsSetting            interface{} // 类评分规则设置：0 正常 ； 1 去最高分和最低分
	RandomSort           interface{} // 随机排序；0否(默认),1是
	ForceVote            interface{} // 强制投票；0否(默认),1是
	AutoEnd              interface{} // 自动结束；0否(默认),1是
	EditResult           interface{} // 修改投票结果；0否(默认),1是
	CountVotes           interface{} // 计票人数；0应投票人数；1实际签到人数
	CreatedAt            *gtime.Time // 创建时间
	UpdatedAt            *gtime.Time // 更新时间
}
