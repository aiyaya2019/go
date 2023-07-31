// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PositionType is the golang structure of table position_type for DAO operations like Where/Data.
type PositionType struct {
	g.Meta    `orm:"table:position_type, do:true"`
	Id        interface{} // 自增ID
	Name      interface{} // 职位名组称
	Sort      interface{} // 排序
	Deleted   interface{} // 软删除：0未删除；1已删除
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
