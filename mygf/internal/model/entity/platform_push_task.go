// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PlatformPushTask is the golang structure for table platform_push_task.
type PlatformPushTask struct {
	Id           int         `json:"id"           ` // 主键
	PlatformUuid uint64      `json:"platformUuid" ` // 平台uuid
	MeetingUuid  uint        `json:"meetingUuid"  ` // 会议uuid，meeting表的uuid
	Status       int         `json:"status"       ` // 0:未邀请1:进行中2:成功3:失败4:编辑中5:编辑成功6:编辑失败7:删除中8:删除成功9:删除失败10:移除中11:移除成功12:移除失败13:取消中14:取消成功15:取消失败16:已退出,17:已拒绝,18:副平台结束
	Operate      int         `json:"operate"      ` // 1add 2edit 3del4:移除5:取消
	JoinTime     *gtime.Time `json:"joinTime"     ` // 加入时间
	CreatedAt    *gtime.Time `json:"createdAt"    ` // 创建时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    ` // 更新时间
}
