package gen

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gen"
	"gorm.io/gen/field"
)

func DefaultModelNameFunc(tableName string) (modelName string) {
	strs := strings.Split(strings.TrimPrefix(tableName, "app_"), "_")
	for i, str := range strs {
		strs[i] = cases.Title(language.English).String(str)
	}
	return strings.Join(strs, "")
}

func DefaultModelOpt() []gen.ModelOpt {
	autoCreateTimeField := gen.FieldGORMTag("create_time", func(tag field.GormTag) field.GormTag {
		tag.Set("column", "create_time")
		tag.Set("type", "int unsigned")
		tag.Set("autoCreateTime", "")
		return tag
	})

	autoUpdateTimeField := gen.FieldGORMTag("update_time", func(tag field.GormTag) field.GormTag {
		tag.Set("column", "update_time")
		tag.Set("type", "int unsigned")
		tag.Set("autoUpdateTime", "")
		return tag
	})

	// 使用时间戳作为软删除字段
	softDeleteField := gen.FieldType("delete_time", "soft_delete.DeletedAt")

	// 模型自定义选项组
	fieldOpts := []gen.ModelOpt{
		autoCreateTimeField,
		autoUpdateTimeField,
		softDeleteField,
	}

	return fieldOpts
}
