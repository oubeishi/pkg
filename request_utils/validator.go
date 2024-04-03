package request_utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func HandleValidateError(err error, c *gin.Context) {
	var validationErrors []interface{}
	// 根据类型返回不同的错误
	switch err.(type) {
	case *validator.InvalidValidationError:
		validationErrors = handleInvalidValidationError(err.(*validator.InvalidValidationError))
		break
	case validator.ValidationErrors:
		validationErrors = HandleValidationErrors(err.(validator.ValidationErrors))
		break
	case *json.UnmarshalTypeError:
		validationErrors = HandleJsonUnmarshalTypeError(err.(*json.UnmarshalTypeError))
		break
	default:
		validationErrors = append(validationErrors, map[string]string{
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusUnprocessableEntity, validationErrors)
}

func handleInvalidValidationError(err *validator.InvalidValidationError) []interface{} {
	var validationErrors []interface{}
	validationErrors = append(validationErrors, map[string]string{
		"message": err.Error(),
	})
	return validationErrors
}

func HandleValidationErrors(errs validator.ValidationErrors) []interface{} {
	var validationErrors []interface{}
	for _, e := range errs {
		validationErrors = append(validationErrors, map[string]string{
			"field":   e.Field(),
			"message": e.Error(),
		})
	}
	return validationErrors
}

func HandleJsonUnmarshalTypeError(err *json.UnmarshalTypeError) []interface{} {
	var validationErrors []interface{}
	validationErrors = append(validationErrors, map[string]string{
		"field":   err.Field,
		"struct":  err.Struct,
		"message": err.Error(),
	})
	return validationErrors
}
