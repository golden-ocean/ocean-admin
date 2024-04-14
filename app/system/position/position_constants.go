package position

const (
	CreatedSuccess = "添加职位成功！"
	CreatedFail    = "添加职位失败！"
	UpdatedSuccess = "修改职位成功！"
	UpdatedFail    = "修改职位失败！"
	DeletedSuccess = "删除职位成功！"
	DeletedFail    = "删除职位失败！"

	ErrorNotExist    = "职位不存在！"
	ErrorNameRepeat  = "职位名称重复，请重新输入！"
	ErrorCodeRepeat  = "职位编码重复，请重新输入！"
	ErrorExistStaffs = "该职位下有员工，无法删除！"
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

	Table = "system_positions"
)

var SelectFields = []string{
	FieldID, FieldName, FieldCode, FieldStatus, FieldSort, FieldRemark,
}
