// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingVoteDao is the data access object for table meeting_vote.
type MeetingVoteDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns MeetingVoteColumns // columns contains all the column names of Table for convenient usage.
}

// MeetingVoteColumns defines and stores column names for table meeting_vote.
type MeetingVoteColumns struct {
	Id                   string // 逻辑ID
	Uuid                 string // 分布式uuid
	PlatformUuid         string // 平台uuid，platform表的uuid
	MeetingUuid          string // 会议uuid，meeting表的uuid
	MeetingDatumUuid     string // 议题uuid，meeting_datum表的uuid
	SystemFileUuid       string // 关联文件uuid，system_file表的uuid
	VoteTitle            string // 投票标题
	VoteName             string // 投票内容
	StartTime            string // 开始时间
	EndTime              string // 结束时间
	Countdown            string // 倒计时
	VoteType             string // 投票类型:0=>自定义；1=>表决
	IsMultiple           string // 是否多选:0=>否，1=>是
	IsTruename           string // 是否实名:0=>否，1=>是
	Status               string // 状态：0=》未启用，1=》启用，2=》 已结束
	Mark                 string // 备注
	IsSign               string // 是否签到
	IsPassRate           string // 是否要通过率
	PassRate             string // 通过率
	IsVoteTimeLimit      string // 是否限时投票
	VoteTimeLimit        string // 投票限时时长
	IsMark               string // 是否需要备注(1为是，0为否)
	ScreenMode           string // 投屏方式0文字,1柱状图,2饼状图
	VotingAuthentication string // 认证方式:0不使用 1签名验证 2人脸识别 3指纹验证
	VotingForce          string // 是否强制认证，0否，1是
	ScoreScore           string // 评分分值，“-”分割
	ScoreType            string // 类型：0 投票 ； 1 评分
	IsSetting            string // 类评分规则设置：0 正常 ； 1 去最高分和最低分
	RandomSort           string // 随机排序；0否(默认),1是
	ForceVote            string // 强制投票；0否(默认),1是
	AutoEnd              string // 自动结束；0否(默认),1是
	EditResult           string // 修改投票结果；0否(默认),1是
	CountVotes           string // 计票人数；0应投票人数；1实际签到人数
	CreatedAt            string // 创建时间
	UpdatedAt            string // 更新时间
}

// meetingVoteColumns holds the columns for table meeting_vote.
var meetingVoteColumns = MeetingVoteColumns{
	Id:                   "id",
	Uuid:                 "uuid",
	PlatformUuid:         "platform_uuid",
	MeetingUuid:          "meeting_uuid",
	MeetingDatumUuid:     "meeting_datum_uuid",
	SystemFileUuid:       "system_file_uuid",
	VoteTitle:            "vote_title",
	VoteName:             "vote_name",
	StartTime:            "start_time",
	EndTime:              "end_time",
	Countdown:            "countdown",
	VoteType:             "vote_type",
	IsMultiple:           "is_multiple",
	IsTruename:           "is_truename",
	Status:               "status",
	Mark:                 "mark",
	IsSign:               "is_sign",
	IsPassRate:           "is_pass_rate",
	PassRate:             "pass_rate",
	IsVoteTimeLimit:      "is_vote_time_limit",
	VoteTimeLimit:        "vote_time_limit",
	IsMark:               "is_mark",
	ScreenMode:           "screen_mode",
	VotingAuthentication: "voting_authentication",
	VotingForce:          "voting_force",
	ScoreScore:           "score_score",
	ScoreType:            "score_type",
	IsSetting:            "is_setting",
	RandomSort:           "random_sort",
	ForceVote:            "force_vote",
	AutoEnd:              "auto_end",
	EditResult:           "edit_result",
	CountVotes:           "count_votes",
	CreatedAt:            "created_at",
	UpdatedAt:            "updated_at",
}

// NewMeetingVoteDao creates and returns a new DAO object for table data access.
func NewMeetingVoteDao() *MeetingVoteDao {
	return &MeetingVoteDao{
		group:   "default",
		table:   "meeting_vote",
		columns: meetingVoteColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MeetingVoteDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MeetingVoteDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MeetingVoteDao) Columns() MeetingVoteColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MeetingVoteDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MeetingVoteDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MeetingVoteDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
