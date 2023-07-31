package v1

import "github.com/gogf/gf/v2/frame/g"

type UserListReq struct {
	g.Meta `path:"/list" tags:"User" method:"get" summary:"You first hello api"`
	Page   int `p:"page"`
	Rows   int `p:"rows"`
	IsTmp  int `p:"is_tmp"`
}

type UserListRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type UserDetailsReq struct {
	g.Meta `path:"/details" tags:"User" method:"get" summary:"人员详情信息"`
	Uuid   int `json:"uuid"`
}

type UserAddReq struct {
	g.Meta   `path:"/add" tags:"User" method:"get" summary:"添加"`
	Username string `json:"username"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Unit     string `json:"unit"`
}

type UserDetailsRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

type UserEditReq struct {
	g.Meta   `path:"/edit" tags:"User" method:"get" summary:"编辑"`
	Uuid     int    `json:"uuid"`
	Username string `json:"username"`
}

type UserDeleteReq struct {
	g.Meta `path:"/delete" tags:"User" method:"get" summary:"删除"`
	Uuid   int `json:"uuid"`
}

type UserRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
