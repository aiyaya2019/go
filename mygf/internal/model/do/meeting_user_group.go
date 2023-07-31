// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// MeetingUserGroup is the golang structure of table meeting_user_group for DAO operations like Where/Data.
type MeetingUserGroup struct {
	g.Meta   `orm:"table:meeting_user_group, do:true"`
	UserUuid interface{} // 用户uuid，user表uuid
	GroupId  interface{} // 分组id，user_group表的id
}
