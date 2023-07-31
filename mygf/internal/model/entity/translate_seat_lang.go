// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// TranslateSeatLang is the golang structure for table translate_seat_lang.
type TranslateSeatLang struct {
	Id              int `json:"id"              ` // 主键id
	MId             int `json:"mId"             ` // 会议id
	TranslateSeatId int `json:"translateSeatId" ` // 同传座位translate_seat的id
	TranslateLangId int `json:"translateLangId" ` // 语言id
	Sort            int `json:"sort"            ` // 排序id
}
