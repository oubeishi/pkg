package reflectx

import (
	"fmt"
	"reflect"
)

// StringMapString2Struct 字符串map转结构体
func StringMapAnyToStruct(params map[string]any, structure any) (any, error) {
	typeOfStructure := reflect.TypeOf(structure)
	switch typeOfStructure.Kind() {
	case reflect.Ptr:
		typeOfStructure = typeOfStructure.Elem()
		if typeOfStructure.Kind() == reflect.Struct {
			// 使用 .NumField() 获取字段数量
			numFields := typeOfStructure.NumField()
			// 遍历结构体的字段
			for i := 0; i < numFields; i++ {
				field := typeOfStructure.Field(i)
				gormColumn := getOrmColumnName(field)
				if v, ok := params[gormColumn]; ok {
					fieldValue := reflect.ValueOf(structure).Elem().FieldByName(field.Name)
					fieldValue.Set(reflect.ValueOf(v))
				}
			}
		} else {
			panic("structure must be a struct")
		}
		break
	case reflect.Struct:
		if typeOfStructure.Kind() == reflect.Struct {
			// 使用 .NumField() 获取字段数量
			numFields := typeOfStructure.NumField()
			fmt.Printf("Number of fields: %d\n", numFields)
			for i := 0; i < numFields; i++ {
				field := typeOfStructure.Field(i)
				gormColumn := getOrmColumnName(field)
				if v, ok := params[gormColumn]; ok {
					fieldValue := reflect.ValueOf(structure).Elem().FieldByName(field.Name)
					fieldValue.Set(reflect.ValueOf(v))
				}
			}
		} else {
			panic("structure must be a struct")
		}
		break
	default:
		panic("structure must be a struct")
	}

	return structure, nil
}
