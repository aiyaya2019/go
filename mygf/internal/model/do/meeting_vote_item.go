// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingVoteItem is the golang structure of table meeting_vote_item for DAO operations like Where/Data.
type MeetingVoteItem struct {
	g.Meta               `orm:"table:meeting_vote_item, do:true"`
	Id                   interface{} //
	Name                 interface{} // 名称
	IsMultiple           interface{} // 是否多选:0=>否，1=>是
	IsTruename           interface{} // 是否实名:0=>否，1=>是
	IsLimit              interface{} // 是否限时:0不限时,1限时
	IsLimitVal           interface{} // 限时时间
	IsRate               interface{} // 是否需要通过率
	IsRateVal            interface{} // 通过率百分比
	MoreThan             interface{} // 0大于1大于等于
	ScreenMode           interface{} // 投屏方式0文字,1柱状图,2饼状图
	ItemJson             interface{} // 投票选项
	IsRemark             interface{} // 是否需要备注(0为否,1为是)
	VotingAuthentication interface{} // 认证方式:0不使用 1签名验证 2人脸识别 3指纹验证
	IsNumberMic          interface{} // 1代表数字会议话筒投票类型
	CreatedAt            *gtime.Time // 创建时间
	UpdatedAt            *gtime.Time // 更新时间
}
