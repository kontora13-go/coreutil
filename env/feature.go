// Copyright 2024 Kontora13. All rights reserved.
// Licensed under the Apache License, Version 2.0

// Feature-флаги для включения функций

package env

import (
	"context"

	"github.com/kontora13-go/coreutil/appcontext"
)

const (
	ctxKeyFeature = "app_feature"
)

// DefaultFeatureCtxValueCount - стартовая ёмкость хранилища значений
var DefaultFeatureCtxValueCount = 5

type appFeatureValue = appcontext.AppCtxValue

// newFeatureCtxValueContext - создание нового значение в контексте
func newFeatureCtxValueContext() appFeatureValue {
	return appcontext.NewAppCtxValue(
		make(map[string]any, DefaultFeatureCtxValueCount),
	)
}

// WithFeatures - получение context со списком features-флагов
func WithFeatures(ctx context.Context, val map[string]bool) context.Context {

	var featureCtx appFeatureValue
	v := ctx.Value(ctxKeyFeature)
	if v != nil {
		var ok bool
		if featureCtx, ok = v.(appFeatureValue); !ok {
			featureCtx = newFeatureCtxValueContext()
			ctx = context.WithValue(ctx, ctxKeyFeature, featureCtx)
		}
	} else {
		featureCtx = newFeatureCtxValueContext()
		ctx = context.WithValue(ctx, ctxKeyFeature, featureCtx)
	}

	for i, v := range val {
		featureCtx.Add(i, v)
	}

	return ctx
}

// CheckFeature - получение из контекста значения по ключу
func CheckFeature(ctx context.Context, key string) bool {
	c := ctx.Value(ctxKeyFeature)
	if c == nil {
		return false
	}

	appCtx, ok := c.(appFeatureValue)
	if !ok {
		return false
	}

	val := appCtx.Value(key)
	valFeature, ok := val.(bool)
	if !ok {
		return false
	}

	return valFeature
}

// SetFeature - установка в контекст значения по произвольному ключу
func SetFeature(ctx context.Context, key string, val bool) context.Context {
	c := ctx.Value(ctxKeyFeature)
	if c == nil {
		return WithFeatures(ctx, map[string]bool{
			key: val,
		})
	}

	appCtx, ok := c.(appFeatureValue)
	if !ok {
		return WithFeatures(ctx, map[string]bool{
			key: val,
		})
	}

	appCtx.Add(key, val)

	return ctx
}

// SetFeatureIfEmpty - установка в контекст значения по произвольному ключу, если оно ещё не задано
func SetFeatureIfEmpty(ctx context.Context, key string, val bool) context.Context {
	c := ctx.Value(ctxKeyFeature)
	if c == nil {
		return WithFeatures(ctx, map[string]bool{
			key: val,
		})
	}

	appCtx, ok := c.(appFeatureValue)
	if !ok {
		return WithFeatures(ctx, map[string]bool{
			key: val,
		})
	}

	if !appCtx.Exist(key) {
		appCtx.Add(key, val)
	}

	return ctx
}
