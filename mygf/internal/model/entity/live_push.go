// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// LivePush is the golang structure for table live_push.
type LivePush struct {
	Id             uint        `json:"id"             ` // 逻辑ID
	SystemFileUuid uint        `json:"systemFileUuid" ` // 直播文件uuid，system_file表的uuid
	Pid            int         `json:"pid"            ` // 进程ID
	Progress       string      `json:"progress"       ` // 进度
	Status         uint        `json:"status"         ` // 推流状态，0=》未开始，1 =》开始，2 =》暂停， 3 =》停止， 4=》 错误，5=》异常，6=》等待中
	CreatedAt      *gtime.Time `json:"createdAt"      ` // 创建时间
	UpdatedAt      *gtime.Time `json:"updatedAt"      ` // 更新时间
}
