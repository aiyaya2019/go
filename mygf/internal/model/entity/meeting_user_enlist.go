// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingUserEnlist is the golang structure for table meeting_user_enlist.
type MeetingUserEnlist struct {
	Id          uint        `json:"id"          ` // 逻辑ID
	MeetingUuid uint        `json:"meetingUuid" ` // 会议uuid，meeting表的uuid
	UserUuid    uint        `json:"userUuid"    ` // 用户uuid，user表的uuid
	Sort        uint        `json:"sort"        ` // 序号
	Username    string      `json:"username"    ` // 名称
	Unit        string      `json:"unit"        ` // 单位
	Dept        string      `json:"dept"        ` // 部门
	Position    string      `json:"position"    ` // 职务
	Mobile      string      `json:"mobile"      ` // 联系电话
	IsFirst     uint        `json:"isFirst"     ` // 是否新注册人员:0否;1是
	Mark        string      `json:"mark"        ` // 备注
	CreateTime  *gtime.Time `json:"createTime"  ` // 创建时间
	Status      int         `json:"status"      ` // 审核状态:-1不通过;0待审核;1审核通过
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 更新时间
}
