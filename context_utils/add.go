package context_utils

import "context"

func AddFillOptions(ctx context.Context, fillOptions ...string) {
	gCtx := GetGinContext(ctx)
	result := make(map[string]bool)
	//安全处理
	if gCtx.Request.Context().Value("fillOptions") != nil {
		result = gCtx.Request.Context().Value("fillOptions").(map[string]bool)
	}
	for _, item := range fillOptions {
		result[item] = true
	}
	gCtx.Request = gCtx.Request.WithContext(context.WithValue(gCtx.Request.Context(), "fillOptions", result))
}
