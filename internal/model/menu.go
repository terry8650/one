package model

type MenuInfoRes struct {
	Id     uint   `orm:"id,primary"  json:"id"`     //
	Pid    uint   `orm:"pid"         json:"pid"`    // 父ID
	Name   string `orm:"name"        json:"name"`   // 规则名称
	Title  string `orm:"title"       json:"title"`  // 规则名称
	Icon   string `orm:"icon"        json:"icon"`   // 图标
	Remark string `orm:"remark"      json:"remark"` // 备注
	Jump   uint   `orm:"jump"        json:"jump"`   // jump

}
type SysMenuTreeRes struct {
	*MenuInfoRes
	Children []*SysMenuTreeRes `json:"children"`
}
type Student struct {
}
