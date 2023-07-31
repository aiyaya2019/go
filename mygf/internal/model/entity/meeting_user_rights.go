// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingUserRights is the golang structure for table meeting_user_rights.
type MeetingUserRights struct {
	Id          int         `json:"id"          ` // 逻辑id
	MeetingUuid uint        `json:"meetingUuid" ` // 会议uuid，meeting表的uuid
	UserUuid    uint        `json:"userUuid"    ` // 用户uuid，user表的uuid
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 更新时间
}
