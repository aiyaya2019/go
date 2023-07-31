// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Device is the golang structure for table device.
type Device struct {
	Id        uint        `json:"id"        ` //
	DeviceIp  string      `json:"deviceIp"  ` // 设备ip
	DeviceMac string      `json:"deviceMac" ` // 设备mac
	RoomId    int         `json:"roomId"    ` // 会议室id
	Status    int         `json:"status"    ` // 设备状态:0空闲；1使用中；2设备离线
	IsShowing int         `json:"isShowing" ` // 当前是否显示图文：0否；1是
	IsPush    int         `json:"isPush"    ` // 是否允许推送：0否；1是
	Type      int         `json:"type"      ` // 设备类型：1信息屏
	Deleted   int         `json:"deleted"   ` // 删除：0未删除；1删除
	Ip        string      `json:"ip"        ` // 服务器ip
	Port      int         `json:"port"      ` // 服务器端口
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
}
