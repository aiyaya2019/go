// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserFingerPrint is the golang structure for table user_finger_print.
type UserFingerPrint struct {
	Id          uint        `json:"id"          ` // 逻辑id
	UserUuid    uint        `json:"userUuid"    ` // 用户uuid，user表的uuid
	Finger      string      `json:"finger"      ` // 手指
	Fingerprint string      `json:"fingerprint" ` // 指纹信息
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 更新时间
}
