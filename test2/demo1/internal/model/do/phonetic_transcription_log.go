// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
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
	Type      interface{} //
	Status    interface{} //
	Detailed  interface{} //
	Creator   interface{} //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}