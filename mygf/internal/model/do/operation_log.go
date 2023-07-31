// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OperationLog is the golang structure of table operation_log for DAO operations like Where/Data.
type OperationLog struct {
	g.Meta        `orm:"table:operation_log, do:true"`
	Id            interface{} // 逻辑ID
	UserUuid      interface{} // 操作用户uuid，user表的uuid
	Module        interface{} // 模块
	Operation     interface{} // 操作
	OperationType interface{} // 操作类型
	Detail        interface{} // 操作详细信息
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
}
