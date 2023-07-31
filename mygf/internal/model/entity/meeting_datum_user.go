// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingDatumUser is the golang structure for table meeting_datum_user.
type MeetingDatumUser struct {
	MeetingUuid      uint        `json:"meetingUuid"      ` // 会议uuid，meeting表的uuid
	UserUuid         uint        `json:"userUuid"         ` // 用户uuid，user表的uuid
	MeetingDatumUuid uint        `json:"meetingDatumUuid" ` // 议题uuid，meeting_datum表的uuid
	CreatedAt        *gtime.Time `json:"createdAt"        ` // 创建时间
	UpdatedAt        *gtime.Time `json:"updatedAt"        ` // 更新时间
}
