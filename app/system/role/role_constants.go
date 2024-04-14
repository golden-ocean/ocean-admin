package role

const (
	CreatedSuccess = "添加角色成功！"
	CreatedFail    = "添加角色失败！"
	UpdatedSuccess = "修改角色成功！"
	UpdatedFail    = "修改角色失败！"
	DeletedSuccess = "删除角色成功！"
	DeletedFail    = "删除角色失败！"

	ErrorNotExist    = "角色不存在！"
	ErrorNameRepeat  = "角色名称重复，请重新输入！"
	ErrorCodeRepeat  = "角色编码重复，请重新输入！"
	ErrorExistMenus  = "该角色下有菜单权限，无法删除! "
	ErrorExistStaffs = "该角色下有员工，无法删除! "
)

const (
	FieldID   = "id"
	FieldName = "name"
	FieldCode = "code"

	FieldStatus    = "status"
	FieldSort      = "sort"
	FieldRemark    = "remark"
	FieldCreatedAt = "created_at"
	FieldUpdatedAt = "updated_at"
	FieldCreatedBy = "created_by"
	FieldUpdatedBy = "updated_by"

	Table         = "system_roles"
	RoleMenuTable = "system_roles_menus"
)

var SelectFields = []string{
	FieldID, FieldName, FieldCode, FieldStatus, FieldSort, FieldRemark,
}
