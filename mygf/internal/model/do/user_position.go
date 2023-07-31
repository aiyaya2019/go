// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserPosition is the golang structure of table user_position for DAO operations like Where/Data.
type UserPosition struct {
	g.Meta     `orm:"table:user_position, do:true"`
	Id         interface{} // 自增ID
	UserUuid   interface{} // 人员uuid；user表的uuid
	PositionId interface{} // 职位id；position表的id
	Default    interface{} // 是否默认职位：0非默认；1默认
	Deleted    interface{} // 软删除：0未删除；1已删除
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
}
