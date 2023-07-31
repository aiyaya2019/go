// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserPosition is the golang structure for table user_position.
type UserPosition struct {
	Id         uint        `json:"id"         ` // 自增ID
	UserUuid   uint64      `json:"userUuid"   ` // 人员uuid；user表的uuid
	PositionId int         `json:"positionId" ` // 职位id；position表的id
	Default    int         `json:"default"    ` // 是否默认职位：0非默认；1默认
	Deleted    int         `json:"deleted"    ` // 软删除：0未删除；1已删除
	CreatedAt  *gtime.Time `json:"createdAt"  ` // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  ` // 更新时间
}
