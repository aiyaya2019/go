// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DeviceMsg is the golang structure for table device_msg.
type DeviceMsg struct {
	Id        uint        `json:"id"        ` //
	DeviceId  int         `json:"deviceId"  ` // 设备id，device表主键id
	Msg       string      `json:"msg"       ` // 消息内容
	Time      int         `json:"time"      ` // 消息播报时间/秒
	Type      int         `json:"type"      ` // 设备类型：1信息屏
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
}
