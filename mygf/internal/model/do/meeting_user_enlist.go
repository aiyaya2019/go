// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingUserEnlist is the golang structure of table meeting_user_enlist for DAO operations like Where/Data.
type MeetingUserEnlist struct {
	g.Meta      `orm:"table:meeting_user_enlist, do:true"`
	Id          interface{} // 逻辑ID
	MeetingUuid interface{} // 会议uuid，meeting表的uuid
	UserUuid    interface{} // 用户uuid，user表的uuid
	Sort        interface{} // 序号
	Username    interface{} // 名称
	Unit        interface{} // 单位
	Dept        interface{} // 部门
	Position    interface{} // 职务
	Mobile      interface{} // 联系电话
	IsFirst     interface{} // 是否新注册人员:0否;1是
	Mark        interface{} // 备注
	CreateTime  *gtime.Time // 创建时间
	Status      interface{} // 审核状态:-1不通过;0待审核;1审核通过
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
