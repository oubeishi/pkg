package reflectx

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// snakeCase converts a string to snake case.
func snakeCase(s string) string {
	var result strings.Builder

	for i, char := range s {
		if i > 0 && 'A' <= char && char <= 'Z' {
			result.WriteRune('_')
		}
		result.WriteRune(char)
	}

	return strings.ToLower(result.String())
}

func setStringDataToField(str string, field reflect.Value) error {
	typeOfStruct := field.Type()
	// 转换 map 中的值到结构体字段
	switch typeOfStruct.Kind() {
	case reflect.Int:
		intValue, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		field.SetInt(int64(intValue))
	case reflect.Uint8:
		uint8Value, err := strconv.ParseUint(str, 10, 8)
		if err != nil {
			return err
		}
		field.SetUint(uint8Value)
	case reflect.Float64:
		floatValue, err := strconv.ParseFloat(str, 64)
		if err != nil {
			if err != nil {
				return err
			}
		}
		field.SetFloat(floatValue)
	case reflect.String:
		field.SetString(str)
	default:
		return fmt.Errorf("Unsupported field type: %v", typeOfStruct)
	}
	return nil
}

func getOrmColumnName(field reflect.StructField) string {
	gormString := field.Tag.Get("gorm")
	gormColumn := ""
	if gormString != "" {
		infoSlice := strings.Split(gormString, ";")
		for _, info := range infoSlice {
			if strings.HasPrefix(info, "column:") {
				gormColumn = strings.TrimPrefix(info, "column:")
			}
		}
	}
	return gormColumn
}
