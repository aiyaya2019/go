// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HtmlFormat is the golang structure of table html_format for DAO operations like Where/Data.
type HtmlFormat struct {
	g.Meta      `orm:"table:html_format, do:true"`
	Id          interface{} // 逻辑ID
	MeetingUuid interface{} // 会议uuid，meeting表的uuid
	MultiUuid   interface{} // 用户/标语uuid，user表的uuid/meeting_banner的uuid
	Url         interface{} // url路径
	Path        interface{} // 文件路径
	Progress    interface{} // 进度
	Status      interface{} // 转码状态，0=》未开始，1=》转换中，2=》成功，3=》失败
	Type        interface{} // 类型，0=》无，1=》html to image，2=》标语
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
