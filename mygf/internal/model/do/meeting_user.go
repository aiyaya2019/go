// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingUser is the golang structure of table meeting_user for DAO operations like Where/Data.
type MeetingUser struct {
	g.Meta              `orm:"table:meeting_user, do:true"`
	Id                  interface{} // 逻辑ID
	MeetingUuid         interface{} // 会议uuid，meeting表的uuid
	UserUuid            interface{} // 用户uuid，user表的uuid
	TerminalId          interface{} // 终端ID;终端信息表中ID，terminal表的id
	IsBroadcast         interface{} // 是否为广播:0=>不能广播，1=>广播
	Status              interface{} // 在会状态:0=>未在会，1=>在会
	IsSecretary         interface{} // 是否会务：0否，1是。当==0时，is_attendee=1
	Username            interface{} // 用户名称
	Nameplate           interface{} // 个人铭牌/欢迎界面数据
	IsUpdateNameplate   interface{} // 铭牌是否更新过：0否；1是
	IsUpdateWelcomePage interface{} // 欢迎界面是否更新过：0否；1是
	Mark                interface{} // 备注
	FreeLogin           interface{} // 是否免密登录,0否
	IsChairman          interface{} // 是否为主席
	IsAttendee          interface{} // 是否参会：0否；1是。当==0时，秘书和主持必须都=1
	Sort                interface{} // 排序
	Lifter              interface{} // 0不是升降器；1升降器，对应升降器id记录在terminal_id
	Deleted             interface{} // 软删除，0正常，1已删除
	Unit                interface{} // 单位
	Position            interface{} // 职务
	IsLocal             interface{} // 本地用户,1是
	Alias               interface{} // 铭牌显示名称
	AttendType          interface{} // 1列席2出席3请假
	IsAutoQueue         interface{} // 是否自动排位
	IsSign              interface{} // 是否签到：0否；1是
	CreatedAt           *gtime.Time // 创建时间
	UpdatedAt           *gtime.Time // 更新时间
}
