// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Organization is the golang structure for table organization.
type Organization struct {
	Id           uint        `json:"id"           ` // 自增ID
	Uuid         uint        `json:"uuid"         ` // 组织架构分布式唯一ID
	Puuid        uint        `json:"puuid"        ` // 父uuid
	PlatformUuid uint        `json:"platformUuid" ` // 平台uuid
	Level        uint        `json:"level"        ` // 层级(1：单位，2：部门，3：职位)
	Name         string      `json:"name"         ` // 名称
	Sort         uint        `json:"sort"         ` // 排序（数值越小，排序越前）
	IsTmp        int         `json:"isTmp"        ` // 是否为临时：0=》否，1=》是
	Status       uint        `json:"status"       ` // 状态：1正常；2禁用
	Deleted      int         `json:"deleted"      ` // 0不删除；1删除
	CreatedAt    *gtime.Time `json:"createdAt"    ` // 创建时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    ` // 更新时间
}
