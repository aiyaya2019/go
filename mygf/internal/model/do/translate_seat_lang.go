// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TranslateSeatLang is the golang structure of table translate_seat_lang for DAO operations like Where/Data.
type TranslateSeatLang struct {
	g.Meta          `orm:"table:translate_seat_lang, do:true"`
	Id              interface{} // 主键id
	MId             interface{} // 会议id
	TranslateSeatId interface{} // 同传座位translate_seat的id
	TranslateLangId interface{} // 语言id
	Sort            interface{} // 排序id
}
