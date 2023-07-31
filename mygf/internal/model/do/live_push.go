// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// LivePush is the golang structure of table live_push for DAO operations like Where/Data.
type LivePush struct {
	g.Meta         `orm:"table:live_push, do:true"`
	Id             interface{} // 逻辑ID
	SystemFileUuid interface{} // 直播文件uuid，system_file表的uuid
	Pid            interface{} // 进程ID
	Progress       interface{} // 进度
	Status         interface{} // 推流状态，0=》未开始，1 =》开始，2 =》暂停， 3 =》停止， 4=》 错误，5=》异常，6=》等待中
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
