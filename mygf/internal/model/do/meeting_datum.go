// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingDatum is the golang structure of table meeting_datum for DAO operations like Where/Data.
type MeetingDatum struct {
	g.Meta           `orm:"table:meeting_datum, do:true"`
	Id               interface{} // 逻辑ID
	Uuid             interface{} // 议题uuid
	MeetingUuid      interface{} // 会议uuid，meeting表的uuid
	UserUuid         interface{} // 上传用户uuid，user表的uuid
	PlatformUuid     interface{} // 平台uuid
	SystemFileUuid   interface{} // 文件uuid
	Name             interface{} // 议题名称
	Status           interface{} // 状态:0=>未审议，1=>审议中，2=>已审议，3=>准备审议中
	Sort             interface{} // 序号：用于文件排序
	Mark             interface{} // 备注
	Reporter         interface{} // 汇报人
	ReportTime       interface{} // 汇报时间
	IsJoinVote       interface{} // 是否加入投票，0否 1加入
	IsSecret         interface{} // 0不保密1保密
	StartTime        *gtime.Time // 开始时间
	EndTime          *gtime.Time // 结束时间
	PushTime         interface{} // 推送时间
	PushStatus       interface{} // 0:未推送1:推送成功；1推送失败
	StartDatum       interface{} // 0议题开启前显示,1议题开启后显示
	RoomSeatSchemeId interface{} // 排位方案id，room_seat_scheme表的id
	ReporterField    interface{} // 汇报人字段
	UserShow         interface{} // 参会人是否显示（0隐藏，1显示）
	Attendee         interface{} // 参会人按查看权限或自定义（0查看权限，1自定义）
	UserDefined      interface{} // 用户自定义字段
	AttendeeField    interface{} // 参会人字段
	ReporterShow     interface{} // 汇报人是否显示（0隐藏，1显示）
	UnitShow         interface{} // 汇报单位是否显示（0隐藏，1显示）
	UnitField        interface{} // 汇报单位字段
	Unit             interface{} // 汇报单位
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 更新时间
}
