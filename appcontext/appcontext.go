// Copyright 2024 Kontora13. All rights reserved.
// Licensed under the Apache License, Version 2.0

// Пакет работы с контекстом приложения.
// Позволяет добавлять в контекст произвольные значения
// и извлекать их из контекста

package appcontext

import (
	"context"
)

const (
	ctxKeyValue = "app_value"
)

// DefaultAppCtxValueCount - стартовая ёмкость хранилища значений
var DefaultAppCtxValueCount = 5

// newAppCtxValueContext - создание нового значение в контексте
func newAppCtxValueContext() AppCtxValue {
	return NewAppCtxValue(
		make(map[string]any, DefaultAppCtxValueCount),
	)
}

// WithValue - получение context со списком values
func WithValue(ctx context.Context, val map[string]any) context.Context {

	var appCtx AppCtxValue
	v := ctx.Value(ctxKeyValue)
	if v != nil {
		var ok bool
		if appCtx, ok = v.(AppCtxValue); !ok {
			appCtx = newAppCtxValueContext()
			ctx = context.WithValue(ctx, ctxKeyValue, appCtx)
		}
	} else {
		appCtx = newAppCtxValueContext()
		ctx = context.WithValue(ctx, ctxKeyValue, appCtx)
	}

	for i, v := range val {
		appCtx.Add(i, v)
	}

	return ctx
}

// GetValue - получение из контекста значения по ключу
func GetValue(ctx context.Context, key string) (val any, ok bool) {
	c := ctx.Value(ctxKeyValue)
	if c == nil {
		return
	}

	appCtx, ok := c.(AppCtxValue)
	if !ok {
		return
	}

	val = appCtx.Value(key)

	return
}

// SetValue - установка в контекст значения по произвольному ключу
func SetValue(ctx context.Context, key string, val any) context.Context {
	c := ctx.Value(ctxKeyValue)
	if c == nil {
		return WithValue(ctx, map[string]any{
			key: val,
		})
	}

	appCtx, ok := c.(AppCtxValue)
	if !ok {
		return WithValue(ctx, map[string]any{
			key: val,
		})
	}

	appCtx.Add(key, val)

	return ctx
}

// SetValueIfEmpty - установка в контекст значения по произвольному ключу, если оно ещё не задано
func SetValueIfEmpty(ctx context.Context, key string, val any) context.Context {
	c := ctx.Value(ctxKeyValue)
	if c == nil {
		return WithValue(ctx, map[string]any{
			key: val,
		})
	}

	appCtx, ok := c.(AppCtxValue)
	if !ok {
		return WithValue(ctx, map[string]any{
			key: val,
		})
	}

	if !appCtx.Exist(key) {
		appCtx.Add(key, val)
	}

	return ctx
}
