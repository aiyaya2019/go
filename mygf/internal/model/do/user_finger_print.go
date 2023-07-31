// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserFingerPrint is the golang structure of table user_finger_print for DAO operations like Where/Data.
type UserFingerPrint struct {
	g.Meta      `orm:"table:user_finger_print, do:true"`
	Id          interface{} // 逻辑id
	UserUuid    interface{} // 用户uuid，user表的uuid
	Finger      interface{} // 手指
	Fingerprint interface{} // 指纹信息
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
