package menu

type CreateInput struct {
	Name      string `zh:"菜单名称" json:"name" validate:"required,min=2,max=32"`
	ParentID  string `zh:"父级菜单" json:"parent_id" validate:"omitempty"`
	Icon      string `zh:"菜单图标" json:"icon"`
	Path      string `zh:"路径" json:"path" validate:"required"`
	Type      string `zh:"类型" json:"type" validate:"required,oneof=Catalog Menu Button"`
	Method    string `zh:"方法" json:"method"`
	Component string `zh:"组件" json:"component"`
	Visible   string `zh:"是否隐藏" json:"visible" validate:"omitempty"`
	Status    string `zh:"状态" json:"status" validate:"omitempty,oneof=Disable Enable"`
	Sort      uint32 `zh:"排序" json:"sort" validate:"number,gt=0"`
	Remark    string `zh:"备注" json:"remark" validate:"omitempty,max=128"`
	CreatedBy string `zh:"创建人员" json:"created_by" validate:"omitempty"`
}

type UpdateInput struct {
	ID        string `zh:"唯一标识符" json:"id" validate:"required"`
	Name      string `zh:"菜单名称" json:"name" validate:"required,min=2,max=32"`
	ParentID  string `zh:"父级菜单" json:"parent_id" validate:"omitempty"`
	Icon      string `zh:"菜单图标" json:"icon"`
	Path      string `zh:"路径" json:"path"`
	Type      string `zh:"类型" json:"type" validate:"required,oneof=Catalog Menu Button"`
	Method    string `zh:"方法" json:"method"`
	Component string `zh:"组件" json:"component"`
	Visible   string `zh:"是否隐藏" json:"visible" validate:"omitempty"`
	Status    string `zh:"状态" json:"status" validate:"omitempty,oneof=Disable Enable"`
	Sort      uint32 `zh:"排序" json:"sort" validate:"number,gt=0"`
	Remark    string `zh:"备注" json:"remark" validate:"omitempty,max=128"`
	UpdatedBy string `zh:"更新人员" json:"updated_by" validate:"omitempty"`
}

type DeleteInput struct {
	ID string `zh:"唯一标识符" json:"id" validate:"required"`
}

type WhereParams struct {
	Name     string `zh:"菜单名称" query:"name" json:"name" validate:"omitempty,max=32"`
	ParentID string `zh:"父级菜单" query:"parent_id" json:"parent_id" validate:"omitempty"`
	Status   string `zh:"状态" query:"status" json:"status" validate:"omitempty,oneof=Disable Enable"`
	Visible  string `zh:"显示" query:"visible" json:"visible" validate:"omitempty"`
	Remark   string `zh:"备注" query:"remark" json:"remark" validate:"omitempty,max=128"`
	PageSize uint64 `zh:"分页数量" query:"pageSize" json:"pageSize" validate:"omitempty,number,gt=0,max=50"`
	Current  uint64 `zh:"页数" query:"current" json:"current" validate:"omitempty,number,gt=0"`
}

type Menu struct {
	ID        string `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	ParentID  string `db:"parent_id" json:"parent_id"`
	Icon      string `db:"icon" json:"icon"`
	Path      string `db:"path" json:"path"`
	Type      string `zh:"类型" json:"type"`
	Method    string `zh:"方法" json:"method"`
	Component string `db:"component" json:"component"`
	Visible   string `db:"visible" json:"visible"`
	Status    string `db:"status" json:"status"`
	Sort      uint32 `db:"sort" json:"sort"`
	Remark    string `db:"remark" json:"remark,omitempty"`
	CreatedAt int64  `db:"created_at" json:"-"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
	CreatedBy string `db:"created_by" json:"-"`
	UpdatedBy string `db:"updated_by" json:"-"`
}

type MenuOutput struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	ParentID  string        `json:"parent_id,omitempty"`
	Icon      string        `json:"icon,omitempty"`
	Path      string        `json:"path,omitempty"`
	Type      string        `json:"type"`
	Method    string        `json:"method,omitempty"`
	Component string        `json:"component,omitempty"`
	Visible   string        `json:"visible"`
	Status    string        `json:"status"`
	Sort      uint32        `json:"sort"`
	Remark    string        `json:"remark"`
	Children  []*MenuOutput `json:"children,omitempty"`
}
