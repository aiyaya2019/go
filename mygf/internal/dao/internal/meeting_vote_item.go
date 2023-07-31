// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingVoteItemDao is the data access object for table meeting_vote_item.
type MeetingVoteItemDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns MeetingVoteItemColumns // columns contains all the column names of Table for convenient usage.
}

// MeetingVoteItemColumns defines and stores column names for table meeting_vote_item.
type MeetingVoteItemColumns struct {
	Id                   string //
	Name                 string // 名称
	IsMultiple           string // 是否多选:0=>否，1=>是
	IsTruename           string // 是否实名:0=>否，1=>是
	IsLimit              string // 是否限时:0不限时,1限时
	IsLimitVal           string // 限时时间
	IsRate               string // 是否需要通过率
	IsRateVal            string // 通过率百分比
	MoreThan             string // 0大于1大于等于
	ScreenMode           string // 投屏方式0文字,1柱状图,2饼状图
	ItemJson             string // 投票选项
	IsRemark             string // 是否需要备注(0为否,1为是)
	VotingAuthentication string // 认证方式:0不使用 1签名验证 2人脸识别 3指纹验证
	IsNumberMic          string // 1代表数字会议话筒投票类型
	CreatedAt            string // 创建时间
	UpdatedAt            string // 更新时间
}

// meetingVoteItemColumns holds the columns for table meeting_vote_item.
var meetingVoteItemColumns = MeetingVoteItemColumns{
	Id:                   "id",
	Name:                 "name",
	IsMultiple:           "is_multiple",
	IsTruename:           "is_truename",
	IsLimit:              "is_limit",
	IsLimitVal:           "is_limit_val",
	IsRate:               "is_rate",
	IsRateVal:            "is_rate_val",
	MoreThan:             "more_than",
	ScreenMode:           "screen_mode",
	ItemJson:             "item_json",
	IsRemark:             "is_remark",
	VotingAuthentication: "voting_authentication",
	IsNumberMic:          "is_number_mic",
	CreatedAt:            "created_at",
	UpdatedAt:            "updated_at",
}

// NewMeetingVoteItemDao creates and returns a new DAO object for table data access.
func NewMeetingVoteItemDao() *MeetingVoteItemDao {
	return &MeetingVoteItemDao{
		group:   "default",
		table:   "meeting_vote_item",
		columns: meetingVoteItemColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *MeetingVoteItemDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *MeetingVoteItemDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *MeetingVoteItemDao) Columns() MeetingVoteItemColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *MeetingVoteItemDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *MeetingVoteItemDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *MeetingVoteItemDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
