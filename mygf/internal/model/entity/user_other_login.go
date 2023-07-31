// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserOtherLogin is the golang structure for table user_other_login.
type UserOtherLogin struct {
	Id        uint        `json:"id"        ` // 逻辑ID
	UserUuid  uint        `json:"userUuid"  ` // 用户uuid，user表的uuid
	Name      string      `json:"name"      ` // 名称
	Pass      string      `json:"pass"      ` // 特征码
	Type      uint        `json:"type"      ` // 1:指纹
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
}
