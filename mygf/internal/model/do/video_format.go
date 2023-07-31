// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// VideoFormat is the golang structure of table video_format for DAO operations like Where/Data.
type VideoFormat struct {
	g.Meta         `orm:"table:video_format, do:true"`
	Id             interface{} // 逻辑ID
	SystemFileUuid interface{} // 视频文件uuid，system_file表的uuid
	Progress       interface{} // 进度
	Status         interface{} // 转码状态，0=》未开始，1=》转换中，2=》成功，3=》失败，4 =》待确认
	Path           interface{} // 文件路径
	Type           interface{} // 0=》 转MP4，aac，1=》全部 copy，2=》视频copy，音频转aac
	Force          interface{} // 强制转码0:否1:是
	Remain         *gtime.Time // 预计剩余时长
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
