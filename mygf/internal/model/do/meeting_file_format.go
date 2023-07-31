// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingFileFormat is the golang structure of table meeting_file_format for DAO operations like Where/Data.
type MeetingFileFormat struct {
	g.Meta      `orm:"table:meeting_file_format, do:true"`
	Id          interface{} // 逻辑ID
	MeetingUuid interface{} // 会议uuid，meeting表的uuid
	UserUuid    interface{} // 用户uuid，user表的uuid
	Url         interface{} // url路径
	Path        interface{} // 文件路径
	Progress    interface{} // 进度
	Status      interface{} // 转码状态，0=》未开始，1=》转换中，2=》成功，3=》失败
	Time        *gtime.Time // 时间
	Type        interface{} // 类型，0=》无，1=》html to image
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
