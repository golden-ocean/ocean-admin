package menu

const (
	CreatedSuccess = "添加菜单成功！"
	CreatedFail    = "添加菜单失败！"
	UpdatedSuccess = "修改菜单成功！"
	UpdatedFail    = "修改菜单失败！"
	DeletedSuccess = "删除菜单成功！"
	DeletedFail    = "删除菜单失败！"

	ErrorNotExist                = "菜单或权限不存在！"
	ErrorPidNotExist             = "父级菜单或权限不存在！"
	ErrorNameRepeat              = "菜单名称重复，请重新输入！"
	ErrorExistChildren           = "菜单下存在子菜单，不能删除！"
	ErrorPidCantEqSelfAndChildId = "父节点不能为自己和其子节点，请重新选择父节点！"
)

const (
	FieldID        = "id"
	FieldName      = "name"
	FieldParentID  = "parent_id"
	FieldIcon      = "icon"
	FieldPath      = "path"
	FieldType      = "type"
	FieldMethod    = "method"
	FieldComponent = "component"
	FieldVisible   = "visible"

	FieldStatus    = "status"
	FieldSort      = "sort"
	FieldRemark    = "remark"
	FieldCreatedAt = "created_at"
	FieldUpdatedAt = "updated_at"
	FieldCreatedBy = "created_by"
	FieldUpdatedBy = "updated_by"

	Table = "system_menus"
)

var SelectFields = []string{
	FieldID, FieldName, FieldParentID, FieldIcon, FieldPath, FieldType, FieldMethod, FieldComponent, FieldVisible,
	FieldStatus, FieldSort, FieldRemark,
}
