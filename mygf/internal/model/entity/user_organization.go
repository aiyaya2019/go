// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UserOrganization is the golang structure for table user_organization.
type UserOrganization struct {
	Id                int  `json:"id"                ` //
	UserUuid          uint `json:"userUuid"          ` // 用户uuid，user表uuid
	OrganizationUuid  uint `json:"organizationUuid"  ` // 组织架构uuid，organization表uuid
	OrganizationLevel int  `json:"organizationLevel" ` // 层级
	Default           int  `json:"default"           ` // 是否默认：0非默认；1默认
	Binding           int  `json:"binding"           ` // 0虚拟绑定的上级(统计人数用)；1真实绑定所在组织
	Deleted           int  `json:"deleted"           ` // 软删除：0未删除1已删除
}
