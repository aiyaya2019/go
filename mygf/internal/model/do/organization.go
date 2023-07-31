// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Organization is the golang structure of table organization for DAO operations like Where/Data.
type Organization struct {
	g.Meta       `orm:"table:organization, do:true"`
	Id           interface{} // 自增ID
	Uuid         interface{} // 组织架构分布式唯一ID
	Puuid        interface{} // 父uuid
	PlatformUuid interface{} // 平台uuid
	Level        interface{} // 层级(1：单位，2：部门，3：职位)
	Name         interface{} // 名称
	Sort         interface{} // 排序（数值越小，排序越前）
	IsTmp        interface{} // 是否为临时：0=》否，1=》是
	Status       interface{} // 状态：1正常；2禁用
	Deleted      interface{} // 0不删除；1删除
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
}
