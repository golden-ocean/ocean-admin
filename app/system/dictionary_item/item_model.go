package dictionary_item

type CreateInput struct {
	DictionaryID string `zh:"字典ID" json:"dictionary_id" validate:"required"`
	Label        string `zh:"选项标签" json:"label" validate:"required,min=2,max=32"`
	Value        string `zh:"选项值" json:"value" validate:"required,min=2,max=64"`
	Color        string `zh:"颜色" json:"color" validate:"required,min=2,max=64"`
	Status       string `zh:"状态" json:"status" validate:"omitempty,oneof=Disable Enable"`
	Sort         int32  `zh:"排序" json:"sort" validate:"omitempty,number,gt=0"`
	Remark       string `zh:"备注" json:"remark" validate:"omitempty,max=128"`
	CreatedBy    string `zh:"创建人员" json:"created_by" validate:"omitempty"`
}

type UpdateInput struct {
	ID           string `zh:"唯一标识符" json:"id" validate:"required"`
	DictionaryID string `zh:"字典ID" query:"dictionary_id" json:"dictionary_id" validate:"required"`
	Label        string `zh:"选项标签" json:"label" validate:"omitempty,min=2,max=32"`
	Value        string `zh:"选项值" json:"value" validate:"omitempty,min=2,max=64"`
	Color        string `zh:"颜色" json:"color" validate:"omitempty,min=2,max=64"`
	Status       string `zh:"状态" json:"status" validate:"omitempty,oneof=Disable Enable"`
	Sort         int32  `zh:"排序" json:"sort" validate:"omitempty,number,gt=0"`
	Remark       string `zh:"备注" json:"remark" validate:"omitempty,max=128"`
	UpdatedBy    string `zh:"更新人员" json:"updated_by" validate:"omitempty"`
}

type DeleteInput struct {
	ID string `zh:"唯一标识符" json:"id" validate:"required"`
}

type WhereParams struct {
	DictionaryID string `zh:"字典ID" query:"dictionary_id" json:"dictionary_id" validate:"required"`
	Label        string `zh:"选项标签" query:"label" json:"label" validate:"omitempty,max=32"`
	Value        string `zh:"选项值" query:"value" json:"value" validate:"omitempty,max=64"`
	Status       string `zh:"状态" query:"status" json:"status" validate:"omitempty,oneof=Disable Enable"`
	Remark       string `zh:"备注" query:"remark" json:"remark" validate:"omitempty,max=128"`
	PageSize     uint64 `zh:"分页数量" query:"pageSize" json:"pageSize" validate:"omitempty,number,gt=0,max=50"`
	Current      uint64 `zh:"页数" query:"current" json:"current" validate:"omitempty,number,gt=0"`
}

type DictionaryItem struct {
	ID           string `db:"id" json:"id"`
	DictionaryID string `db:"dictionary_id" json:"dictionary_id"`
	Label        string `db:"label" json:"label"`
	Value        string `db:"value" json:"value"`
	Color        string `db:"color" json:"color"`
	Status       string `db:"status" json:"status"`
	Sort         int32  `db:"sort" json:"sort"`
	Remark       string `db:"remark" json:"remark,omitempty"`
	CreatedAt    int64  `db:"created_at" json:"-"`
	UpdatedAt    int64  `db:"updated_at" json:"updated_at"`
	CreatedBy    string `db:"created_by" json:"-"`
	UpdatedBy    string `db:"updated_by" json:"-"`
}

type DictionaryItemOutput struct {
	ID     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Color  string `json:"color"`
	Status string `json:"status"`
	Sort   int32  `json:"sort"`
	Remark string `json:"remark"`
}
