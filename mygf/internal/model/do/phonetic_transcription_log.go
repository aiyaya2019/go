// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// PhoneticTranscriptionLog is the golang structure of table phonetic_transcription_log for DAO operations like Where/Data.
type PhoneticTranscriptionLog struct {
	g.Meta    `orm:"table:phonetic_transcription_log, do:true"`
	Id        interface{} //
	Type      interface{} // 1=>手动同步，2=》自动同步
	Status    interface{} // 1=》同步成功，2=》同步失败
	Detailed  interface{} // 同步详情
	Creator   interface{} // 创建者
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
