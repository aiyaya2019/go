// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingSign is the golang structure for table meeting_sign.
type MeetingSign struct {
	Id             uint        `json:"id"             ` // 逻辑ID
	MeetingUuid    uint        `json:"meetingUuid"    ` // 会议uuid，meeting表的uuid
	SystemFileUuid uint        `json:"systemFileUuid" ` // 文件uuid，system_file表的uuid
	SignStartTime  *gtime.Time `json:"signStartTime"  ` // 发起签到时间
	Title          string      `json:"title"          ` // 标题
	SignType       int         `json:"signType"       ` // 签到方式：1签名签到；2登录即签到 3按钮签到 4拍照签到 5人脸签到 6指纹签到
	Status         int         `json:"status"         ` // 当前签到状态：0未开始，1签到中, 2 已结束
	QrCode         string      `json:"qrCode"         ` // 签到二维码
	CreatedAt      *gtime.Time `json:"createdAt"      ` // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      ` // 更新时间
}
