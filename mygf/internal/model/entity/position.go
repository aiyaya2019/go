// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Position is the golang structure for table position.
type Position struct {
	Id             uint        `json:"id"             ` // 自增ID
	PositionTypeId int         `json:"positionTypeId" ` // 职位组id；position_type表的id
	Name           string      `json:"name"           ` // 职位名称
	Sort           int         `json:"sort"           ` // 排序
	Deleted        int         `json:"deleted"        ` // 软删除：0未删除；1已删除
	CreatedAt      *gtime.Time `json:"createdAt"      ` // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      ` // 更新时间
}
