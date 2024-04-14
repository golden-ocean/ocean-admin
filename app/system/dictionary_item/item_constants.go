package dictionary_item

const (
	CreatedSuccess = "添加选项成功！"
	CreatedFail    = "添加选项失败！"
	UpdatedSuccess = "修改选项成功！"
	UpdatedFail    = "修改选项失败！"
	DeletedSuccess = "删除选项成功！"
	DeletedFail    = "删除选项失败！"

	ErrorNotExist    = "选项不存在！"
	ErrorLabelRepeat = "选项标签重复，请重新输入！"
	ErrorValueRepeat = "选项值重复，请重新输入！"
)

const (
	FieldID           = "id"
	FieldDictionaryID = "dictionary_id"
	FieldLabel        = "label"
	FieldValue        = "value"
	FieldColor        = "color"

	FieldStatus    = "status"
	FieldSort      = "sort"
	FieldRemark    = "remark"
	FieldCreatedAt = "created_at"
	FieldUpdatedAt = "updated_at"
	FieldCreatedBy = "created_by"
	FieldUpdatedBy = "updated_by"

	Table           = "system_dictionary_items"
	DictionaryTable = "system_dictionaries"
)

var SelectFields = []string{FieldID, FieldLabel, FieldValue, FieldColor, FieldStatus, FieldSort, FieldRemark}
