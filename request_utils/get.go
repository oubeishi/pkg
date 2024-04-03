package request_utils

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oubeishi/pkg/response_utils"
	"io"
	"reflect"
	"strconv"
)

func IgetInt(c *gin.Context, key string) int {
	queryValue := c.Query(key)
	if queryValue != "" {
		intValue, _ := strconv.Atoi(queryValue)
		return intValue
	}
	FormValue := c.PostForm(key)
	if FormValue != "" {
		intValue, _ := strconv.Atoi(FormValue)
		return intValue
	}
	// 从body的json中获取

	return 0
}

func GetMap(c *gin.Context, structure any) map[string]any {

	body, _ := io.ReadAll(c.Request.Body)
	copyBody := body
	var requestBody map[string]any
	structureMap := make(map[string]any)
	err := json.Unmarshal(body, &requestBody)
	if err != nil {
		response_utils.IAbort(nil, fmt.Sprintf("请求参数错误，请检查参数类型，错误信息：%s", err.Error()), 400, 400)
	}
	// 根据传入的结构体类型，创建一个新的结构体
	newStructure := reflect.New(reflect.TypeOf(structure)).Interface()
	err = json.Unmarshal(copyBody, newStructure)
	if err != nil {
		response_utils.IAbort(nil, fmt.Sprintf("请求参数错误，请检查参数类型，错误信息：%s", err.Error()), 400, 400)
	}
	typeOfStructure := reflect.TypeOf(newStructure)

	switch typeOfStructure.Kind() {
	case reflect.Ptr:
		typeOfStructure = typeOfStructure.Elem()
		if typeOfStructure.Kind() == reflect.Struct {
			// 使用 .NumField() 获取字段数量
			numFields := typeOfStructure.NumField()
			// 遍历结构体的字段
			for i := 0; i < numFields; i++ {
				field := typeOfStructure.Field(i) // 获取字段的反射类型
				tag := field.Tag.Get("json")
				if requestBody[tag] == nil {
					continue
				}
				structureMap[tag] = reflect.ValueOf(newStructure).Elem().Field(i).Interface()
			}
		}
	case reflect.Struct:

		numFields := typeOfStructure.NumField()
		// 遍历结构体的字段
		for i := 0; i < numFields; i++ {
			field := typeOfStructure.Field(i) // 获取字段的反射类型
			tag := field.Tag.Get("json")
			if requestBody[tag] == nil {
				continue
			}
			structureMap[tag] = reflect.ValueOf(newStructure).Elem().Field(i).Interface()
		}
	}
	instanceValue := reflect.ValueOf(newStructure)

	// 确保结构体指针的值是可取地址的
	//if instanceValue.Kind() == reflect.Ptr {
	//	// 获取指针指向的值
	//	instanceValue = instanceValue.Elem()
	//}
	//CustomCheck是指针方法
	methodName := "CustomCheck"
	method := instanceValue.MethodByName(methodName)
	if method.IsValid() {

		// 执行方法，传递指针
		method.Call(nil)
	} else {
		fmt.Printf("Struct does not have method %s\n", methodName)
	}
	return structureMap
}
