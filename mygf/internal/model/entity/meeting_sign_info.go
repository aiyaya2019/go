// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingSignInfo is the golang structure for table meeting_sign_info.
type MeetingSignInfo struct {
	Id             uint        `json:"id"             ` // 逻辑ID
	MeetingUuid    uint        `json:"meetingUuid"    ` // 会议uuid，meeting表的uuid
	UserUuid       uint        `json:"userUuid"       ` // 用户uuid，user表的uuid
	SystemFileUuid uint        `json:"systemFileUuid" ` // 签到文件uuid，system_file表的uuid
	MeetingSignId  uint        `json:"meetingSignId"  ` // 签到标记ID，meeting_sign表的id
	Status         int         `json:"status"         ` // 签到状态：0未签到；1已签到
	SignTime       *gtime.Time `json:"signTime"       ` // 签到时间
	Assist         int         `json:"assist"         ` // 是否协助签到，1:是
	DeviceType     int         `json:"deviceType"     ` // 设备类型：1无纸化；2电子桌牌
	CreatedAt      *gtime.Time `json:"createdAt"      ` // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      ` // 更新时间
}
