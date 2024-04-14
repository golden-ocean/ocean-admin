package staff_role

type CreateInput struct {
	StaffID   string   `zh:"用户ID" json:"staff_id,omitempty"`
	RoleIDs   []string `zh:"角色ID列表" json:"role_ids,omitempty"`
	CreatedBy string   `zh:"创建人员" json:"created_by" validate:"omitempty"`
}

type UpdateInput struct {
	StaffID   string `zh:"用户ID" json:"staff_id,omitempty"`
	RoleID    string `zh:"角色ID" json:"role_id,omitempty"`
	UpdatedBy string `zh:"更新人员" json:"updated_by" validate:"omitempty"`
}

type DeleteInput struct {
	StaffID string   `zh:"用户ID" json:"staff_id" validate:"required"`
	RoleIDs []string `zh:"角色ID列表" json:"role_ids" validate:"required"`
}

type WhereParams struct {
	StaffID  string   `zh:"用户ID" query:"staff_id" json:"staff_id" validate:"omitempty"`
	StaffIDs []string `zh:"用户ID列表" query:"staff_ids" json:"staff_ids" validate:"omitempty"`
}

type StaffRole struct {
	StaffID   string `db:"staff_id" json:"staff_id"`
	RoleID    string `db:"role_id" json:"role_id"`
	CreatedAt int64  `db:"created_at" json:"-"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
	CreatedBy string `db:"created_by" json:"-"`
	UpdatedBy string `db:"updated_by" json:"-"`
}
