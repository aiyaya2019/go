// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Position is the golang structure of table position for DAO operations like Where/Data.
type Position struct {
	g.Meta         `orm:"table:position, do:true"`
	Id             interface{} // 自增ID
	PositionTypeId interface{} // 职位组id；position_type表的id
	Name           interface{} // 职位名称
	Sort           interface{} // 排序
	Deleted        interface{} // 软删除：0未删除；1已删除
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
