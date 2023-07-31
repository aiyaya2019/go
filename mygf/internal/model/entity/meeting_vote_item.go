// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingVoteItem is the golang structure for table meeting_vote_item.
type MeetingVoteItem struct {
	Id                   int         `json:"id"                   ` //
	Name                 string      `json:"name"                 ` // 名称
	IsMultiple           uint        `json:"isMultiple"           ` // 是否多选:0=>否，1=>是
	IsTruename           uint        `json:"isTruename"           ` // 是否实名:0=>否，1=>是
	IsLimit              uint        `json:"isLimit"              ` // 是否限时:0不限时,1限时
	IsLimitVal           string      `json:"isLimitVal"           ` // 限时时间
	IsRate               uint        `json:"isRate"               ` // 是否需要通过率
	IsRateVal            uint        `json:"isRateVal"            ` // 通过率百分比
	MoreThan             uint        `json:"moreThan"             ` // 0大于1大于等于
	ScreenMode           uint        `json:"screenMode"           ` // 投屏方式0文字,1柱状图,2饼状图
	ItemJson             string      `json:"itemJson"             ` // 投票选项
	IsRemark             int         `json:"isRemark"             ` // 是否需要备注(0为否,1为是)
	VotingAuthentication uint        `json:"votingAuthentication" ` // 认证方式:0不使用 1签名验证 2人脸识别 3指纹验证
	IsNumberMic          uint        `json:"isNumberMic"          ` // 1代表数字会议话筒投票类型
	CreatedAt            *gtime.Time `json:"createdAt"            ` // 创建时间
	UpdatedAt            *gtime.Time `json:"updatedAt"            ` // 更新时间
}
