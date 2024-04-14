package staff

const (
	CreatedSuccess = "添加员工成功！"
	CreatedFail    = "添加员工失败！"
	UpdatedSuccess = "修改员工成功！"
	UpdatedFail    = "修改员工失败！"
	DeletedSuccess = "删除员工成功！"
	DeletedFail    = "删除员工失败！"

	ErrorNotExist          = "员工不存在！"
	ErrorUsernameRepeat    = "用户名称重复，请重新输入！"
	ErrorEmailRepeat       = "电子邮箱重复，请重新输入！"
	ErrorMobileRepeat      = "移动电话重复，请重新输入！"
	ErrorDeleteAdminRecord = "不能删除超级管理员！"
)

const (
	FieldID             = "id"
	FieldUsername       = "username"
	FieldPassword       = "password"
	FieldName           = "name"
	FieldEmail          = "email"
	FieldMobile         = "mobile"
	FieldAvatar         = "avatar"
	FieldGender         = "gender"
	FieldOrganizationID = "organization_id"
	FieldWorkStatus     = "work_status"

	FieldStatus    = "status"
	FieldSort      = "sort"
	FieldRemark    = "remark"
	FieldCreatedAt = "created_at"
	FieldUpdatedAt = "updated_at"
	FieldCreatedBy = "created_by"
	FieldUpdatedBy = "updated_by"

	Table              = "system_staffs"
	StaffPositionTable = "system_staffs_positions"
	StaffRoleTable     = "system_staffs_roles"
)

var SelectFields = []string{
	FieldID, FieldUsername, FieldName, FieldEmail, FieldMobile, FieldAvatar, FieldGender, FieldOrganizationID, FieldWorkStatus, FieldStatus, FieldSort, FieldRemark,
}
