// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Register is the golang structure of table register for DAO operations like Where/Data.
type Register struct {
	g.Meta    `orm:"table:register, do:true"`
	Id        interface{} // 逻辑ID
	Content   interface{} // 注册码
	Status    interface{} // 状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
