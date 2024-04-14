package staff

type CreateInput struct {
	Username       string   `zh:"用户名称" json:"username" validate:"required,min=4,max=64,alphanum"`
	Password       string   `zh:"密码" json:"password" validate:"required,min=6,max=64"`
	Name           string   `zh:"用户姓名" json:"name" validate:"required,min=2,max=32"`
	Email          string   `zh:"电子邮件" json:"email" validate:"required,email"`
	Mobile         string   `zh:"移动电话" json:"mobile" validate:"required,len=11"`
	Avatar         string   `zh:"用户头像" json:"avatar" validate:"omitempty"`
	Gender         string   `zh:"用户性别" json:"gender" validate:"omitempty"`
	OrganizationID string   `zh:"归属部门" json:"organization_id" validate:"omitempty"`
	WorkStatus     string   `zh:"在职状态" json:"work_status" validate:"omitempty"`
	Status         string   `zh:"状态" json:"status" validate:"required,oneof=Disable Enable"`
	Sort           int32    `zh:"排序" json:"sort" validate:"omitempty,number,gt=0"`
	Remark         string   `zh:"备注" json:"remark" validate:"omitempty,max=128"`
	PositionIDs    []string `zh:"职位IDs" json:"position_ids" validate:"omitempty"`
	RoleIDs        []string `zh:"角色IDs" json:"role_ids" validate:"omitempty"`
	CreatedBy      string   `zh:"创建人员" json:"created_by" validate:"omitempty"`
}

type UpdateInput struct {
	ID             string   `zh:"唯一标识符" json:"id" validate:"required"`
	Username       string   `zh:"用户名称" json:"username" validate:"omitempty,min=4,max=64,alphanum"`
	Password       string   `zh:"密码" json:"password" validate:"omitempty,min=6,max=64"`
	Name           string   `zh:"用户姓名" json:"name" validate:"omitempty,min=2,max=32"`
	Email          string   `zh:"电子邮件" json:"email" validate:"omitempty,email"`
	Mobile         string   `zh:"移动电话" json:"mobile" validate:"omitempty,len=11"`
	Avatar         string   `zh:"用户头像" json:"avatar" validate:"omitempty"`
	Gender         string   `zh:"用户性别" json:"gender" validate:"omitempty,oneof=Male Female Unknown"`
	OrganizationID string   `zh:"归属部门ID" json:"organization_id" validate:"omitempty"`
	WorkStatus     string   `zh:"在职状态" json:"work_status" validate:"omitempty"`
	Status         string   `zh:"状态" json:"status" validate:"omitempty,oneof=Disable Enable"`
	Sort           int32    `zh:"排序" json:"sort" validate:"omitempty,number,gt=0"`
	Remark         string   `zh:"备注" json:"remark" validate:"omitempty,max=128"`
	PositionIDs    []string `zh:"职位IDs" json:"position_ids" validate:"omitempty"`
	RoleIDs        []string `zh:"角色IDs" json:"role_ids" validate:"omitempty"`
	UpdatedBy      string   `zh:"更新人员" json:"updated_by" validate:"omitempty"`
}

type DeleteInput struct {
	ID          string   `zh:"唯一标识符" json:"id" validate:"required"`
	PositionIDs []string `zh:"职位IDs" json:"position_ids" validate:"required"`
	RoleIDs     []string `zh:"角色IDs" json:"role_ids" validate:"required"`
}

type WhereParams struct {
	ID             string `zh:"唯一标识符" query:"id" json:"id" validate:"omitempty"`
	Username       string `zh:"用户名称" query:"username" json:"username" validate:"omitempty,max=64"`
	Name           string `zh:"真实姓名" query:"name" json:"name" validate:"omitempty,max=16"`
	Email          string `zh:"邮件地址" query:"email" json:"email" validate:"omitempty,max=32"`
	Mobile         string `zh:"电话号码" query:"mobile" json:"mobile" validate:"omitempty,max=11"`
	Gender         string `zh:"性别" query:"gender" json:"gender" validate:"omitempty"`
	OrganizationID string `zh:"部门" query:"organization_id" json:"organization_id" validate:"omitempty"`
	WorkStatus     string `zh:"在职状态" query:"work_status" json:"work_status" validate:"omitempty"`
	Status         string `zh:"状态" query:"status" json:"status" validate:"omitempty,oneof=Disable Enable"`
	Remark         string `zh:"备注" query:"remark" json:"remark" validate:"omitempty,max=128"`
	PageSize       uint64 `zh:"分页数量" query:"pageSize" json:"pageSize" validate:"omitempty,number,gt=0,max=50"`
	Current        uint64 `zh:"页数" query:"current" json:"current" validate:"omitempty,number,gt=0"`
}

type Staff struct {
	ID             string `db:"id" json:"id"`
	Username       string `db:"username" json:"username"`
	Password       string `db:"password" json:"-"`
	Name           string `db:"name" json:"name"`
	Email          string `db:"email" json:"email,omitempty"`
	Mobile         string `db:"mobile" json:"mobile,omitempty"`
	Avatar         string `db:"avatar" json:"avatar"`
	Gender         string `db:"gender" json:"gender"`
	OrganizationID string `db:"organization_id" json:"organization_id"`
	WorkStatus     string `db:"work_status" json:"work_status"`
	Status         string `db:"status" json:"status"`
	Sort           int32  `db:"sort" json:"sort"`
	Remark         string `db:"remark" json:"remark,omitempty"`
	CreatedAt      int64  `db:"created_at" json:"-"`
	UpdatedAt      int64  `db:"updated_at" json:"updated_at"`
	CreatedBy      string `db:"created_by" json:"-"`
	UpdatedBy      string `db:"updated_by" json:"-"`
}

type StaffOutput struct {
	ID             string   `json:"id"`
	Username       string   `json:"username"`
	Name           string   `json:"name"`
	Email          string   `json:"email"`
	Mobile         string   `json:"mobile"`
	Avatar         string   `json:"avatar"`
	Gender         string   `json:"gender"`
	OrganizationID string   `json:"organization_id"`
	WorkStatus     string   `json:"work_status"`
	Status         string   `json:"status"`
	Sort           int32    `json:"sort"`
	Remark         string   `json:"remark"`
	PositionIDs    []string `json:"position_ids"`
	RoleIDs        []string `json:"role_ids"`
}

// func (output *StaffOutput) Edges(e ent.StaffEdges) {
// 	p_ids := lo.Map(e.Positions, func(item *ent.Position, _ int) string {
// 		return item.ID
// 	})
// 	r_ids := lo.Map(e.Roles, func(item *ent.Role, _ int) string {
// 		return item.ID
// 	})
// 	output.PositionIDs = p_ids
// 	output.RoleIDs = r_ids
// }
