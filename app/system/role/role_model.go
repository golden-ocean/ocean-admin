package role

type CreateInput struct {
	Name      string `zh:"角色名称" json:"name" validate:"required,min=2,max=32"`
	Code      string `zh:"角色编码" json:"code" validate:"required,min=2,max=64"`
	Status    string `zh:"状态" json:"status" validate:"required,oneof=Disable Enable Unknown"`
	Sort      int32  `zh:"排序" json:"sort" validate:"omitempty,number,gt=0"`
	Remark    string `zh:"备注" json:"remark" validate:"omitempty,max=128"`
	CreatedBy string `zh:"创建人员" json:"created_by" validate:"omitempty"`
}

type UpdateInput struct {
	ID        string `zh:"唯一标识符" json:"id" validate:"required"`
	Name      string `zh:"角色名称" json:"name" validate:"omitempty,min=2,max=32"`
	Code      string `zh:"角色编码" json:"code" validate:"omitempty,min=2,max=64"`
	Status    string `zh:"状态" json:"status" validate:"omitempty,oneof=Disable Enable Unknown"`
	Sort      int32  `zh:"排序" json:"sort" validate:"omitempty,number,gt=0"`
	Remark    string `zh:"备注" json:"remark" validate:"omitempty,max=128"`
	UpdatedBy string `zh:"更新人员" json:"updated_by" validate:"omitempty"`
}

type DeleteInput struct {
	ID string `zh:"唯一标识符" json:"id" validate:"required"`
}

type WhereParams struct {
	Name     string `zh:"角色名称" query:"name" json:"name" validate:"omitempty,max=32"`
	Code     string `zh:"角色编码" query:"code" json:"code" validate:"omitempty,max=64"`
	Status   string `zh:"状态" query:"status" json:"status" validate:"omitempty,oneof=Disable Enable Unknown"`
	Remark   string `zh:"备注" query:"remark" json:"remark" validate:"omitempty,max=128"`
	PageSize uint64 `zh:"分页数量" query:"pageSize" json:"pageSize" validate:"omitempty,number,gt=0,max=50"`
	Current  uint64 `zh:"页数" query:"current" json:"current" validate:"omitempty,number,gt=0"`
}

type RoleMenuInput struct {
	RoleID  string                  `zh:"角色ID" params:"role_id" json:"role_id" validate:"required"`
	MenuIDs []string                `zh:"菜单ID列表" params:"menu_ids" json:"menu_ids" validate:"omitempty"`
	Rules   []*CasbinPermissionRule `zh:"规则" json:"rules" validate:"omitempty"`
}

type CasbinPermissionRule struct {
	RoleID string `zh:"角色ID" params:"role_id" json:"role_id" validate:"omitempty"`
	MenuID string `zh:"菜单ID" params:"menu_id" json:"menu_id" validate:"omitempty"`
	Path   string `zh:"菜单路径" params:"path" json:"path" validate:"omitempty,max=128"`
	Method string `zh:"请求方法" params:"method" json:"method" validate:"omitempty,max=16"`
}

type Role struct {
	ID        string `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Code      string `db:"code" json:"code"`
	Status    string `db:"status" json:"status"`
	Sort      int32  `db:"sort" json:"sort"`
	Remark    string `db:"remark" json:"remark,omitempty"`
	CreatedAt int64  `db:"created_at" json:"-"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
	CreatedBy string `db:"created_by" json:"-"`
	UpdatedBy string `db:"updated_by" json:"updated_by"`
}

type RoleOutput struct {
	ID     string `json:"id"`
	Name   string `json:"name,omitempty"`
	Code   string `json:"code,omitempty"`
	Status string `json:"status,omitempty"`
	Sort   int32  `json:"sort,omitempty"`
	Remark string `json:"remark,omitempty"`
	// Menus  []*menu.MenuOutput `json:"menus,omitempty"`
}

// func (output *RoleOutput) Edges(e ent.RoleEdges) {
// 	// m_ids := lo.Map(e.Menus, func(item *ent.Menu, _ int) string {
// 	// 	return item.ID
// 	// })
// 	output.Menus = e.Menus
// }
