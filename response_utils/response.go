package response_utils

import (
	"ai_girlfriend_server/pkg/pagination"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"net/http"
)

const (
	prefix           = "ddd-gin-admin"
	userIdKey        = prefix + "/user-id"
	ReqBodyKey       = prefix + "/req-body"
	ResBodyKey       = prefix + "/res-body"
	LoggerReqBodyKey = prefix + "/logger-req-body"
)

type StatusText string

func (t StatusText) String() string {
	return string(t)
}

const (
	OKStatus    StatusText = "OK"
	ErrorStatus StatusText = "ERROR"
	FailStatus  StatusText = "FAIL"
)

type StatusResult struct {
	Status StatusText `json:"status"` // Result status
}

type ErrorResult struct {
	Error ErrorItem `json:"error"` // Error
}

type ErrorItem struct {
	Code    int    `json:"code"`    // Error Code
	Message string `json:"message"` // Error Message
}

type ListResult struct {
	List       any                    `json:"list"`                 // List
	Pagination *pagination.Pagination `json:"pagination,omitempty"` // Pagination
}

type FinalResponse map[string]any

func NewIDResult(id string) *IDResult {
	return &IDResult{
		ID: id,
	}
}

type IDResult struct {
	ID string `json:"id"`
}

func ResPage(c *gin.Context, v any, pr *pagination.Pagination) {
	list := ListResult{
		List:       v,
		Pagination: pr,
	}
	ResSuccess(c, list, http.StatusOK)
}
func ResError(c *gin.Context, err error, status ...int) {
	//ctx := c.Request.Context()
	var res *ResponseError
	switch err.(type) {
	case *ResponseError:
		res = err.(*ResponseError)
		break
	case *mysql.MySQLError:
		res = ConvertToResponseError(*err.(*mysql.MySQLError))
		break
	default:
		res = &ResponseError{
			Code:       500,
			Message:    err.Error(),
			StatusCode: 500,
			ERR:        err,
		}
	}

	if len(status) > 0 {
		res.StatusCode = status[0]
	}

	if err := res.ERR; err != nil {
		if res.Message == "" {
			res.Message = err.Error()
		}

		if status := res.StatusCode; status >= 400 && status < 500 {
			//logger.WithContext(ctx).Warnf(err.Error())

		} else if status >= 500 {
			//logger.WithContext(logger.NewStackContext(ctx, err)).Errorf(err.Error())
		}
	}

	mapString := make(map[string]interface{})
	mapString["code"] = res.Code
	mapString["message"] = res.Message
	ResJSON(c, res.StatusCode, mapString)
}

// ResSuccess response_utils success
//
// Parameters:
//
//	param1 interface{}: data
//	param2 int: code
func ResSuccess(c *gin.Context, params ...interface{}) {
	mapString := FinalResponse{
		"code": 200,
	}
	if len(params) > 0 {
		mapString["data"] = params[0]
		if len(params) > 1 {
			mapString["code"] = params[1]
		}
		if len(params) > 2 {
			mapString["message"] = params[2]
		}
	} else {
		ResJSON(c, http.StatusOK, params)
	}
	ResJSON(c, http.StatusOK, mapString)
}

func ResJSON(c *gin.Context, status int, v interface{}) {
	buf, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	c.Set(ResBodyKey, buf)
	c.Data(status, "application/json; charset=utf-8", buf)
	c.Abort()
}

func IAbort(c context.Context, message string, nums ...int) {
	//ginContext := context_utils.GetGinContext(c)
	switch len(nums) {
	case 2:
		//ginContext.JSON(nums[1], gin.H{
		//	"code":    nums[0],
		//	"message": message,
		//})
		panic(&ResponseError{
			Code:       nums[0],
			Message:    message,
			StatusCode: nums[1],
		})
		break
	case 1:
		panic(&ResponseError{
			Code:       nums[0],
			Message:    message,
			StatusCode: http.StatusOK,
		})
		break
	default:
		panic(&ResponseError{
			Code:       400,
			Message:    message,
			StatusCode: http.StatusOK,
		})
		break
	}
	return
}

func NewErrorRes(code, statusCode int, msg string, args ...interface{}) error {
	res := &ResponseError{
		Code:       code,
		Message:    fmt.Sprintf(msg, args...),
		StatusCode: statusCode,
	}
	return res
}
