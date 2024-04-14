package dictionary

const (
	CreatedSuccess = "添加字典成功！"
	CreatedFail    = "添加字典失败！"
	UpdatedSuccess = "修改字典成功！"
	UpdatedFail    = "修改字典失败！"
	DeletedSuccess = "删除字典成功！"
	DeletedFail    = "删除字典失败！"

	ErrorNotExist      = "字典不存在！"
	ErrorNameRepeat    = "字典名称重复，请重新输入！"
	ErrorCodeRepeat    = "字典编码重复，请重新输入！"
	ErrorExistChildren = "此字典下有选项，请删除其所有选项后再操作！"
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

	Table      = "system_dictionaries"
	ItemsTable = "system_dictionary_items"
)

var SelectFields = []string{FieldID, FieldName, FieldCode, FieldStatus, FieldSort, FieldRemark}
