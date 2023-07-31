// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
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
	MeetingUuid interface{} //
	TransId     interface{} //
	Type        interface{} //
	FileName    interface{} //
	FileSize    interface{} //
	FilePath    interface{} //
	FileTime    interface{} //
	StartTime   *gtime.Time //
	EndTime     *gtime.Time //
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
}
