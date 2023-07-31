// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserFile is the golang structure of table user_file for DAO operations like Where/Data.
type UserFile struct {
	g.Meta         `orm:"table:user_file, do:true"`
	MeetingUuid    interface{} // 会议uuid，meeting表的uuid
	UserUuid       interface{} // 用户uuid，user表的uuid
	SystemFileUuid interface{} // 文件uuid，system_file表的uuid
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
