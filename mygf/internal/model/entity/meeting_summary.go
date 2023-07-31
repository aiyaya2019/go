// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingSummary is the golang structure for table meeting_summary.
type MeetingSummary struct {
	Id                       uint        `json:"id"                       ` // 逻辑id
	MeetingUuid              uint        `json:"meetingUuid"              ` // 会议uuid，meeting表的uuid
	SystemFileUuid           uint        `json:"systemFileUuid"           ` // 纪要文件uuid，system_file表的uuid
	UserUuid                 uint        `json:"userUuid"                 ` // 纪要上传者用户uuid, user表的uuid
	Status                   uint        `json:"status"                   ` // 状态 0草稿 1会签中 2已会签
	IsUpdate                 int         `json:"isUpdate"                 ` // 是否已更新:0否；1是
	FileFormatSystemFileUuid uint        `json:"fileFormatSystemFileUuid" ` // 结束会签原始文件；file_format表的system_file_uuid
	MeetingDatumUuid         uint        `json:"meetingDatumUuid"         ` // 关联议题uuid;meeting_datum表的uuid
	CreatedAt                *gtime.Time `json:"createdAt"                ` // 创建时间
	UpdatedAt                *gtime.Time `json:"updatedAt"                ` // 更新时间
	SyncInfo                 string      `json:"syncInfo"                 ` //
}
