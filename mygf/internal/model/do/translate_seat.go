// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// TranslateSeat is the golang structure of table translate_seat for DAO operations like Where/Data.
type TranslateSeat struct {
	g.Meta `orm:"table:translate_seat, do:true"`
	Id     interface{} // 主键id
	SeatNo interface{} // 编号
	Title  interface{} // 标题
	Xway   interface{} // x坐标
	Yway   interface{} // y坐标
}
