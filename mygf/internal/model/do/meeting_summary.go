// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingSummary is the golang structure of table meeting_summary for DAO operations like Where/Data.
type MeetingSummary struct {
	g.Meta                   `orm:"table:meeting_summary, do:true"`
	Id                       interface{} // 逻辑id
	MeetingUuid              interface{} // 会议uuid，meeting表的uuid
	SystemFileUuid           interface{} // 纪要文件uuid，system_file表的uuid
	UserUuid                 interface{} // 纪要上传者用户uuid, user表的uuid
	Status                   interface{} // 状态 0草稿 1会签中 2已会签
	IsUpdate                 interface{} // 是否已更新:0否；1是
	FileFormatSystemFileUuid interface{} // 结束会签原始文件；file_format表的system_file_uuid
	MeetingDatumUuid         interface{} // 关联议题uuid;meeting_datum表的uuid
	CreatedAt                *gtime.Time // 创建时间
	UpdatedAt                *gtime.Time // 更新时间
	SyncInfo                 interface{} //
}
