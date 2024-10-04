// Copyright 2024 Kontora13. All rights reserved.
// Licensed under the Apache License, Version 2.0

// Работа с контекстом трассировки запроса

package trace

import (
	"context"

	"github.com/kontora13-go/coreutil/appcontext"
)

const (
	ctxKeyTrace   = "app_trace"
	ctxKeyTraceId = "trace_id"
)

type appCtxTrace = appcontext.AppCtxValue

// newTraceContext - создание нового объекта для записи данных трассировки в контекст
func newTraceContext() appCtxTrace {
	return appcontext.NewAppCtxValue(map[string]any{
		ctxKeyTraceId: "",
	})
}

// WithTrace - получение context с данными трассировки приложения
func WithTrace(ctx context.Context, traceId string) context.Context {
	var appCtx appCtxTrace
	v := ctx.Value(ctxKeyTrace)
	if v != nil {
		var ok bool
		if appCtx, ok = v.(appCtxTrace); !ok {
			appCtx = newTraceContext()
			ctx = context.WithValue(ctx, ctxKeyTrace, appCtx)
		}
	} else {
		appCtx = newTraceContext()
		ctx = context.WithValue(ctx, ctxKeyTrace, appCtx)
	}

	appCtx.Add(ctxKeyTraceId, traceId)

	return ctx
}

// TraceId - получение идентификатора трассировки из контекста
func TraceId(ctx context.Context) (val string, ok bool) {
	c := ctx.Value(ctxKeyTrace)
	if c == nil {
		return
	}

	appCtx, ok := c.(appCtxTrace)
	if !ok {
		return
	}

	v := appCtx.Value(ctxKeyTraceId)
	val, ok = v.(string)

	return
}

// SetTraceId - запись идентификатора трассировки в контекст
func SetTraceId(ctx context.Context, val string) context.Context {
	c := ctx.Value(ctxKeyTrace)
	if c == nil {
		return WithTrace(ctx, val)
	}

	appCtx, ok := c.(appCtxTrace)
	if !ok {
		return WithTrace(ctx, val)
	}

	appCtx.Add(ctxKeyTraceId, val)

	return ctx
}
