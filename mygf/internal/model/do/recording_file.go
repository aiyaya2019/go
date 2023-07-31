// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RecordingFile is the golang structure of table recording_file for DAO operations like Where/Data.
type RecordingFile struct {
	g.Meta      `orm:"table:recording_file, do:true"`
	Id          interface{} //
	MeetingUuid interface{} // 会议uuid，meeting表的uuid
	TransId     interface{} // 语音转写那边的id
	Type        interface{} // 0=>单个话筒，1=》所有话筒，2全部的压缩包
	FileName    interface{} // 文件名称
	FileSize    interface{} // 文件大小
	FilePath    interface{} // 文件路径
	FileTime    interface{} // 转写时长
	StartTime   *gtime.Time // 转写开始时间
	EndTime     *gtime.Time // 转写结束时间
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
