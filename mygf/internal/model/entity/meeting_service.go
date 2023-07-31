// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingService is the golang structure for table meeting_service.
type MeetingService struct {
	Id                  uint        `json:"id"                  ` // 逻辑ID
	MeetingUuid         uint        `json:"meetingUuid"         ` // 会议uuid，meeting表的uuid
	UserUuid            uint        `json:"userUuid"            ` // 用户uuid，user表的uuid
	ServerId            uint        `json:"serverId"            ` // 服务ID，server表的id
	Content             string      `json:"content"             ` // 服务内容
	Time                *gtime.Time `json:"time"                ` // 时间
	IsConfirm           uint        `json:"isConfirm"           ` // 是否确认：0=>未确认，1=> 已确认
	Mark                string      `json:"mark"                ` // 备注
	Status              uint        `json:"status"              ` // 0:未处理，1:稍后处理，2:已处理，3:已过期,4:已取消
	LaterHandleUserUuid uint        `json:"laterHandleUserUuid" ` // 稍后处理者uuid，user表的uuid
	HandleUserUuid      uint        `json:"handleUserUuid"      ` // 处理者uuid，user表的uuid
	LaterHandleTime     *gtime.Time `json:"laterHandleTime"     ` // 稍后处理时间
	HandleTime          *gtime.Time `json:"handleTime"          ` // 处理时间
	DeviceType          int         `json:"deviceType"          ` // 设备类型：1无纸化；2电子桌牌
	CreatedAt           *gtime.Time `json:"createdAt"           ` // 创建时间
	UpdatedAt           *gtime.Time `json:"updatedAt"           ` // 更新时间
}
