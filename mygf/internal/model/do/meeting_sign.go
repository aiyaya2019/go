// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MeetingSign is the golang structure of table meeting_sign for DAO operations like Where/Data.
type MeetingSign struct {
	g.Meta         `orm:"table:meeting_sign, do:true"`
	Id             interface{} // 逻辑ID
	MeetingUuid    interface{} // 会议uuid，meeting表的uuid
	SystemFileUuid interface{} // 文件uuid，system_file表的uuid
	SignStartTime  *gtime.Time // 发起签到时间
	Title          interface{} // 标题
	SignType       interface{} // 签到方式：1签名签到；2登录即签到 3按钮签到 4拍照签到 5人脸签到 6指纹签到
	Status         interface{} // 当前签到状态：0未开始，1签到中, 2 已结束
	QrCode         interface{} // 签到二维码
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
