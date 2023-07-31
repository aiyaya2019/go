// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DeviceCustom is the golang structure for table device_custom.
type DeviceCustom struct {
	Id           uint        `json:"id"           ` //
	IsShowDatum  int         `json:"isShowDatum"  ` // 是否显示议题：0不显示；1显示
	SkinColor    int         `json:"skinColor"    ` // 1浅蓝；2深蓝
	BgImg        string      `json:"bgImg"        ` // 背景皮肤
	Logo         string      `json:"logo"         ` // 空闲页logo
	PropagateImg string      `json:"propagateImg" ` // 空闲页宣传图
	Type         int         `json:"type"         ` // 设备类型：1信息屏
	CreatedAt    *gtime.Time `json:"createdAt"    ` // 创建时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    ` // 更新时间
}
