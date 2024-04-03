package context_utils

import "context"

func SetFillOptions(ctx context.Context, fillOptions map[string]bool) {
	gCtx := GetGinContext(ctx)
	gCtx.Request = gCtx.Request.WithContext(context.WithValue(gCtx.Request.Context(), "fillOptions", fillOptions))
}
