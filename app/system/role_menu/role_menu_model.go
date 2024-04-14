package role_menu

type CreateInput struct {
	RoleID    string   `zh:"角色ID" json:"role_id,omitempty"`
	MenuIDs   []string `zh:"菜单ID列表" json:"menu_ids,omitempty"`
	CreatedBy string   `zh:"创建人员" json:"created_by" validate:"omitempty"`
}

type UpdateInput struct {
	RoleID    string `zh:"角色ID" json:"role_id,omitempty"`
	MenuID    string `zh:"菜单ID" json:"menu_id,omitempty"`
	UpdatedBy string `zh:"更新人员" json:"updated_by" validate:"omitempty"`
}

type DeleteInput struct {
	RoleID  string   `zh:"角色ID" json:"role_id" validate:"required"`
	MenuIDs []string `zh:"菜单ID列表" json:"menu_ids" validate:"required"`
}

type WhereParams struct {
	RoleID  string   `zh:"角色ID" query:"role_id" json:"role_id" validate:"omitempty"`
	RoleIDs []string `zh:"角色ID列表" query:"role_ids" json:"role_ids" validate:"omitempty"`
}

type RoleMenu struct {
	RoleID    string `db:"role_id" json:"role_id"`
	MenuID    string `db:"menu_id" json:"menu_id"`
	CreatedAt int64  `db:"created_at" json:"-"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
	CreatedBy string `db:"created_by" json:"-"`
	UpdatedBy string `db:"updated_by" json:"-"`
}
