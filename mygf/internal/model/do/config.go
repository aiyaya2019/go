// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Config is the golang structure of table config for DAO operations like Where/Data.
type Config struct {
	g.Meta    `orm:"table:config, do:true"`
	Id        interface{} // 配置id
	Key       interface{} // 配置键名标识
	Value     interface{} // 配置内容
	Default   interface{} // 默认配置内容
	Mark      interface{} // 备注
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
