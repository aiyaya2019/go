// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PlatformPushTask is the golang structure of table platform_push_task for DAO operations like Where/Data.
type PlatformPushTask struct {
	g.Meta       `orm:"table:platform_push_task, do:true"`
	Id           interface{} // 主键
	PlatformUuid interface{} // 平台uuid
	MeetingUuid  interface{} // 会议uuid，meeting表的uuid
	Status       interface{} // 0:未邀请1:进行中2:成功3:失败4:编辑中5:编辑成功6:编辑失败7:删除中8:删除成功9:删除失败10:移除中11:移除成功12:移除失败13:取消中14:取消成功15:取消失败16:已退出,17:已拒绝,18:副平台结束
	Operate      interface{} // 1add 2edit 3del4:移除5:取消
	JoinTime     *gtime.Time // 加入时间
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
}
