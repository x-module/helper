/**
 * Created by Goland.
 * @file   gen.go
 * @author 李锦 <lijin@cavemanstudio.net>
 * @date   2023/7/23 01:10
 * @desc   gen.go
 */

package utils

import (
	"github.com/x-module/helper/strutil"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// GetGenerator 获取生成器
func GetGenerator(db *gorm.DB, outPath string) *gen.Generator {
	// 生成实例
	generator := gen.NewGenerator(gen.Config{
		// 相对执行`go run`时的路径, 会自动创建目录
		OutPath: outPath,
		// // 代码输出文件名，默认: gen.go
		OutFile: "gen.go",
		// 生成的 model 包名
		ModelPkgPath: "model",
		// 是否为生成的查询类代码生成单元测试
		WithUnitTest: false,

		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,

		// 表字段可为 null 值时, 对应结体字段使用指针类型
		FieldNullable: true, // generate pointer when field is nullable

		// 表字段默认值与模型结构体字段零值不一致的字段, 在插入数据时需要赋值该字段值为零值的, 结构体字段须是指针类型才能成功, 即`FieldCoverable:true`配置下生成的结构体字段.
		// 因为在插入时遇到字段为零值的会被GORM赋予默认值. 如字段`age`表默认值为10, 即使你显式设置为0最后也会被GORM设为10提交.
		// 如果该字段没有上面提到的插入时赋零值的特殊需要, 则字段为非指针类型使用起来会比较方便.
		FieldCoverable: false, // generate pointer when field has default value, to fix problem zero value cannot be assign: https://gorm.io/docs/create.html#Default-Values

		// 模型结构体字段的数字类型的符号表示是否与表字段的一致, `false`指示都用有符号类型
		FieldSignable: false, // detect integer field's unsigned type, adjust generated data type
		// 生成 gorm 标签的字段索引属性
		FieldWithIndexTag: true, // generate with gorm index tag
		// 生成 gorm 标签的字段类型属性
		FieldWithTypeTag: true, // generate with gorm column type tag
	})
	// 设置目标 db
	generator.UseDB(db)
	// 统一数字类型为int64,兼容protobuf
	var dataMap = map[string]func(gorm.ColumnType) (dataType string){
		// int mapping
		"decimal": func(columnType gorm.ColumnType) (dataType string) {
			// if n, ok := columnType.Nullable(); ok && n {
			//	return "*decimal.Decimal"
			// }
			return "decimal.Decimal"
		}, // int mapping
		"tinyint": func(columnType gorm.ColumnType) (dataType string) {
			// if n, ok := columnType.Nullable(); ok && n {
			//	return "int64"
			// }
			return "int64"
		},
		"smallint": func(columnType gorm.ColumnType) (dataType string) {
			// if n, ok := columnType.Nullable(); ok && n {
			//	return "*int64"
			// }
			return "int64"
		},
		"mediumint": func(columnType gorm.ColumnType) (dataType string) {
			// if n, ok := columnType.Nullable(); ok && n {
			//	return "*int64"
			// }
			return "int64"
		},
		"bigint": func(columnType gorm.ColumnType) (dataType string) {
			// if n, ok := columnType.Nullable(); ok && n {
			//	return "*int64"
			// }
			return "int64"
		},
		"int": func(columnType gorm.ColumnType) (dataType string) {
			// if n, ok := columnType.Nullable(); ok && n {
			//	return "*int64"
			// }
			return "int64"
		},
		// bool mapping
		// "tinyint": func(columnType gorm.ColumnType) (dataType string) {
		//	ct, _ := columnType.ColumnType()
		//	if strings.HasPrefix(ct, "tinyint(1)") {
		//		return "bool"
		//	}
		//	return "byte"
		// },
	}
	// 要先于`ApplyBasic`执行
	generator.WithDataTypeMap(dataMap)
	generator.WithJSONTagNameStrategy(func(c string) string {
		return strutil.CamelCase(c)
	})
	return generator
}

func GetGeneratorOptions() []gen.ModelOpt {
	// 自定义模型结体字段的标签
	// 将特定字段名的 json 标签加上`string`属性,即 MarshalJSON 时该字段由数字类型转成字符串类型
	// jsonField := gen.FieldJSONTagWithNS(func(columnName string) (tagContent string) {
	//	toStringField := `uid, `
	//	if strings.Contains(toStringField, columnName) {
	//		return columnName + ",string"
	//	}
	//	return columnName
	// })
	// 将非默认字段名的字段定义为自动时间戳和软删除字段;
	// 自动时间戳默认字段名为:`updated_at`、`created_at, 表字段数据类型为: INT 或 DATETIME
	// 软删除默认字段名为:`deleted_at`, 表字段数据类型为: DATETIME
	// autoUpdateTimeField := gen.FieldGORMTag("id", func(tag field.GormTag) field.GormTag {
	//	tag["comment"] = []string{"啦啦啦===啦啦啦"}
	//	return tag
	// })

	// base := gen.FieldNew("BaseModel", "", field.Tag{"json": "-"})

	// autoUpdateTimeField := gen.FieldGORMTag("update_time", "column:update_time;type:int unsigned;autoUpdateTime")
	// autoCreateTimeField := gen.FieldGORMTag("create_time", "column:create_time;type:int unsigned;autoCreateTime")
	// softDeleteField := gen.FieldType("delete_time", "gorm.DeletedAt")

	// 模型自定义选项组
	return []gen.ModelOpt{
		// jsonField,
		// autoCreateTimeField,
		// base,
		// autoUpdateTimeField,
		// softDeleteField,
	}
}
