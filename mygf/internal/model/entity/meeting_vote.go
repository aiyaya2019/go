// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingVote is the golang structure for table meeting_vote.
type MeetingVote struct {
	Id                   uint        `json:"id"                   ` // 逻辑ID
	Uuid                 uint        `json:"uuid"                 ` // 分布式uuid
	PlatformUuid         uint        `json:"platformUuid"         ` // 平台uuid，platform表的uuid
	MeetingUuid          uint        `json:"meetingUuid"          ` // 会议uuid，meeting表的uuid
	MeetingDatumUuid     uint        `json:"meetingDatumUuid"     ` // 议题uuid，meeting_datum表的uuid
	SystemFileUuid       uint        `json:"systemFileUuid"       ` // 关联文件uuid，system_file表的uuid
	VoteTitle            string      `json:"voteTitle"            ` // 投票标题
	VoteName             string      `json:"voteName"             ` // 投票内容
	StartTime            *gtime.Time `json:"startTime"            ` // 开始时间
	EndTime              *gtime.Time `json:"endTime"              ` // 结束时间
	Countdown            uint        `json:"countdown"            ` // 倒计时
	VoteType             uint        `json:"voteType"             ` // 投票类型:0=>自定义；1=>表决
	IsMultiple           uint        `json:"isMultiple"           ` // 是否多选:0=>否，1=>是
	IsTruename           uint        `json:"isTruename"           ` // 是否实名:0=>否，1=>是
	Status               int         `json:"status"               ` // 状态：0=》未启用，1=》启用，2=》 已结束
	Mark                 string      `json:"mark"                 ` // 备注
	IsSign               uint        `json:"isSign"               ` // 是否签到
	IsPassRate           string      `json:"isPassRate"           ` // 是否要通过率
	PassRate             float64     `json:"passRate"             ` // 通过率
	IsVoteTimeLimit      uint        `json:"isVoteTimeLimit"      ` // 是否限时投票
	VoteTimeLimit        string      `json:"voteTimeLimit"        ` // 投票限时时长
	IsMark               uint        `json:"isMark"               ` // 是否需要备注(1为是，0为否)
	ScreenMode           uint        `json:"screenMode"           ` // 投屏方式0文字,1柱状图,2饼状图
	VotingAuthentication int         `json:"votingAuthentication" ` // 认证方式:0不使用 1签名验证 2人脸识别 3指纹验证
	VotingForce          int         `json:"votingForce"          ` // 是否强制认证，0否，1是
	ScoreScore           string      `json:"scoreScore"           ` // 评分分值，“-”分割
	ScoreType            int         `json:"scoreType"            ` // 类型：0 投票 ； 1 评分
	IsSetting            int         `json:"isSetting"            ` // 类评分规则设置：0 正常 ； 1 去最高分和最低分
	RandomSort           int         `json:"randomSort"           ` // 随机排序；0否(默认),1是
	ForceVote            int         `json:"forceVote"            ` // 强制投票；0否(默认),1是
	AutoEnd              int         `json:"autoEnd"              ` // 自动结束；0否(默认),1是
	EditResult           int         `json:"editResult"           ` // 修改投票结果；0否(默认),1是
	CountVotes           int         `json:"countVotes"           ` // 计票人数；0应投票人数；1实际签到人数
	CreatedAt            *gtime.Time `json:"createdAt"            ` // 创建时间
	UpdatedAt            *gtime.Time `json:"updatedAt"            ` // 更新时间
}
