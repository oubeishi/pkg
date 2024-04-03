package reflectx

import (
	"fmt"
	"reflect"
)

// StringMapString2Struct 字符串map转结构体
func StringMapString2Struct(inputMap map[string]string, resultStruct any) (any, error) {
	resultValue := reflect.ValueOf(resultStruct)
	if resultValue.Kind() != reflect.Ptr || resultValue.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("目标不是结构体指针")
	}

	resultElem := resultValue.Elem()
	resultType := resultElem.Type()

	for i := 0; i < resultElem.NumField(); i++ {
		field := resultElem.Field(i)
		fieldType := resultType.Field(i)

		// 获取结构体标签
		tagValue := fieldType.Tag.Get("redis")

		// 使用Snake Case格式
		tagValue = snakeCase(tagValue)

		// 从 map 中获取对应的值
		mapValue, ok := inputMap[tagValue]
		if !ok {
			continue
		}
		setStringDataToField(mapValue, field)
		//typeOfStruct := field.Type()
		//// 转换 map 中的值到结构体字段
		//switch typeOfStruct.Kind() {
		//case reflect.Int:
		//	intValue, err := strconv.Atoi(mapValue)
		//	if err != nil {
		//		return nil, err
		//	}
		//	field.SetInt(int64(intValue))
		//case reflect.Uint8:
		//	uint8Value, err := strconv.ParseUint(mapValue, 10, 8)
		//	if err != nil {
		//		return nil, err
		//	}
		//	field.SetUint(uint8Value)
		//case reflect.Float64:
		//	floatValue, err := strconv.ParseFloat(mapValue, 64)
		//	if err != nil {
		//		if err != nil {
		//			return nil, err
		//		}
		//	}
		//	field.SetFloat(floatValue)
		//case reflect.String:
		//	field.SetString(mapValue)
		//default:
		//	return nil, fmt.Errorf("Unsupported field type: %v", fieldType)
		//}
	}

	return resultStruct, nil
}
