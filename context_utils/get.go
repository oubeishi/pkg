package context_utils

import (
	"context"
	"github.com/gin-gonic/gin"
)

func GetFillOptions(ctx context.Context, fillOptions ...string) map[string]bool {
	gCtx := GetGinContext(ctx)
	result := make(map[string]bool)
	//安全处理
	if gCtx.Request.Context().Value("fillOptions") != nil {
		result = gCtx.Request.Context().Value("fillOptions").(map[string]bool)
	}
	for _, item := range fillOptions {
		result[item] = true
	}
	return result
}

// 用了就删除，以免在其他地方被使用
func ConsumeFillOption(ctx context.Context, key string) (bool, bool) {
	gCtx := GetGinContext(ctx)
	result := make(map[string]bool)
	//安全处理
	if gCtx.Request.Context().Value("fillOptions") != nil {
		result = gCtx.Request.Context().Value("fillOptions").(map[string]bool)
	}
	if v, ok := result[key]; ok {
		delete(result, key)
		return v, true
	}
	return false, false
}

//	func ConsumesFillOptions(ctx context.Context, keys ...string) map[string]bool {
//		gCtx := GetGinContext(ctx)
//		oldFillOptions := make(map[string]bool)
//		result := make(map[string]bool)
//		//安全处理
//		if gCtx.Request.Context().Value("fillOptions") != nil {
//			oldFillOptions = gCtx.Request.Context().Value("fillOptions").(map[string]bool)
//		}
//		for _, key := range keys {
//			if v, ok := oldFillOptions[key]; ok {
//				result[key] = v
//				delete(oldFillOptions, key)
//			}
//		}
//		return result
//	}
func GetGinContext(ctx context.Context) *gin.Context {
	return ctx.Value("ginCtx").(*gin.Context)
}

func GetUserId(ctx context.Context) int {
	gCtx := GetGinContext(ctx)
	return gCtx.GetInt("userId")
}
