// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Platform is the golang structure for table platform.
type Platform struct {
	Id        int         `json:"id"        ` //
	Uuid      uint64      `json:"uuid"      ` // 平台uuid
	Name      string      `json:"name"      ` // 平台名称
	IsLocal   int         `json:"isLocal"   ` // 是否为本地：0远端；1本地
	Sysip     string      `json:"sysip"     ` // 平台ip
	Sysport   int         `json:"sysport"   ` // 平台端口
	Sysid     int         `json:"sysid"     ` // 平台自定义id，唯一
	Status    int         `json:"status"    ` // 状态:0离线，1=>在线
	Enable    int         `json:"enable"    ` // 0禁用；1开启
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
}
