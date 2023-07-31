// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RecordingFile is the golang structure for table recording_file.
type RecordingFile struct {
	Id          uint        `json:"id"          ` //
	MeetingUuid uint        `json:"meetingUuid" ` // 会议uuid，meeting表的uuid
	TransId     int         `json:"transId"     ` // 语音转写那边的id
	Type        int         `json:"type"        ` // 0=>单个话筒，1=》所有话筒，2全部的压缩包
	FileName    string      `json:"fileName"    ` // 文件名称
	FileSize    string      `json:"fileSize"    ` // 文件大小
	FilePath    string      `json:"filePath"    ` // 文件路径
	FileTime    string      `json:"fileTime"    ` // 转写时长
	StartTime   *gtime.Time `json:"startTime"   ` // 转写开始时间
	EndTime     *gtime.Time `json:"endTime"     ` // 转写结束时间
	CreatedAt   *gtime.Time `json:"createdAt"   ` // 创建时间
	UpdatedAt   *gtime.Time `json:"updatedAt"   ` // 更新时间
}
