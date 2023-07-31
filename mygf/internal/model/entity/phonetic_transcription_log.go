// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PhoneticTranscriptionLog is the golang structure for table phonetic_transcription_log.
type PhoneticTranscriptionLog struct {
	Id        uint        `json:"id"        ` //
	Type      int         `json:"type"      ` // 1=>手动同步，2=》自动同步
	Status    int         `json:"status"    ` // 1=》同步成功，2=》同步失败
	Detailed  string      `json:"detailed"  ` // 同步详情
	Creator   string      `json:"creator"   ` // 创建者
	CreatedAt *gtime.Time `json:"createdAt" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" ` // 更新时间
}
