// Copyright 2024 Kontora13. All rights reserved.
// Licensed under the Apache License, Version 2.0

// Работа с env-данными с использованием контекста приложения

package env

import (
	"context"

	"github.com/kontora13-go/coreutil/appcontext"
)

const (
	ctxKeyEnv            = "app_env"
	ctxKeyEnvAppZone     = "app_zone"
	ctxKeyEnvAppName     = "app_name"
	ctxKeyEnvAppVersion  = "app_version"
	ctxKeyEnvAppInstance = "app_instance"
)

type appCtxValue = appcontext.AppCtxValue

// newEnvCtxValue - создание пустого набора env-значений
func newEnvCtxValue() appCtxValue {
	return appcontext.NewAppCtxValue(map[string]any{
		ctxKeyEnvAppZone:     ZoneNotDefined,
		ctxKeyEnvAppName:     "",
		ctxKeyEnvAppVersion:  EmptyVersion,
		ctxKeyEnvAppInstance: "",
	})
}

// WithEnv - получение context с env приложения
func WithEnv(
	ctx context.Context,
	appZone Zone, appName string, appVersion Version, appInstance string,
) context.Context {

	var appCtx appCtxValue
	v := ctx.Value(ctxKeyEnv)
	if v != nil {
		var ok bool
		if appCtx, ok = v.(appCtxValue); !ok {
			appCtx = newEnvCtxValue()
			ctx = context.WithValue(ctx, ctxKeyEnv, appCtx)
		}
	} else {
		appCtx = newEnvCtxValue()
		ctx = context.WithValue(ctx, ctxKeyEnv, appCtx)
	}

	appCtx.Add(ctxKeyEnvAppZone, appZone)
	appCtx.Add(ctxKeyEnvAppName, appName)
	appCtx.Add(ctxKeyEnvAppVersion, appVersion)
	appCtx.Add(ctxKeyEnvAppInstance, appInstance)

	return ctx
}

// AppZone - извлечение Zone приложения из контекста
func AppZone(ctx context.Context) (val Zone, ok bool) {
	c := ctx.Value(ctxKeyEnv)
	if c == nil {
		return ZoneNotDefined, false
	}

	appCtx, ok := c.(appCtxValue)
	if !ok {
		return ZoneNotDefined, false
	}

	v := appCtx.Value(ctxKeyEnvAppZone)
	val, ok = v.(Zone)
	if !ok {
		return ZoneNotDefined, false
	}
	if err := val.IsValid(); err != nil {
		return ZoneNotDefined, false
	}

	return
}

// SetAppZone - установка Zone приложения в контекста
func SetAppZone(ctx context.Context, val Zone) context.Context {
	c := ctx.Value(ctxKeyEnv)
	if c == nil {
		return WithEnv(ctx, val, "", EmptyVersion, "")
	}

	appCtx, ok := c.(appCtxValue)
	if !ok {
		return WithEnv(ctx, val, "", EmptyVersion, "")
	}

	appCtx.Add(ctxKeyEnvAppZone, val)

	return ctx
}

// AppName - извлечение названия приложения из контекста
func AppName(ctx context.Context) (val string, ok bool) {
	c := ctx.Value(ctxKeyEnv)
	if c == nil {
		return
	}

	appCtx, ok := c.(appCtxValue)
	if !ok {
		return
	}

	v := appCtx.Value(ctxKeyEnvAppName)
	val, ok = v.(string)

	return
}

// SetAppName - запись названия приложения в контекст
func SetAppName(ctx context.Context, val string) context.Context {
	c := ctx.Value(ctxKeyEnv)
	if c == nil {
		return WithEnv(ctx, ZoneNotDefined, val, EmptyVersion, "")
	}

	appCtx, ok := c.(appCtxValue)
	if !ok {
		return WithEnv(ctx, ZoneNotDefined, val, EmptyVersion, "")
	}

	appCtx.Add(ctxKeyEnvAppName, val)

	return ctx
}

// AppVersion - извлечение версии приложения из контекста
func AppVersion(ctx context.Context) (val Version, ok bool) {
	c := ctx.Value(ctxKeyEnv)
	if c == nil {
		return
	}

	appCtx, ok := c.(appCtxValue)
	if !ok {
		return
	}

	v := appCtx.Value(ctxKeyEnvAppVersion)
	val, ok = v.(Version)

	return
}

// SetAppVersion - установка версии приложения в контекст
func SetAppVersion(ctx context.Context, val Version) context.Context {
	c := ctx.Value(ctxKeyEnv)
	if c == nil {
		return WithEnv(ctx, ZoneNotDefined, "", val, "")
	}

	appCtx, ok := c.(appCtxValue)
	if !ok {
		return WithEnv(ctx, ZoneNotDefined, "", val, "")
	}

	appCtx.Add(ctxKeyEnvAppVersion, val)

	return ctx
}

func AppInstance(ctx context.Context) (val string, ok bool) {
	c := ctx.Value(ctxKeyEnv)
	if c == nil {
		return
	}

	appCtx, ok := c.(appCtxValue)
	if !ok {
		return
	}

	v := appCtx.Value(ctxKeyEnvAppInstance)
	val, ok = v.(string)

	return
}

func SetAppInstance(ctx context.Context, val string) context.Context {
	c := ctx.Value(ctxKeyEnv)
	if c == nil {
		return WithEnv(ctx, ZoneNotDefined, "", EmptyVersion, val)
	}

	appCtx, ok := c.(appCtxValue)
	if !ok {
		return WithEnv(ctx, ZoneNotDefined, "", EmptyVersion, val)
	}

	appCtx.Add(ctxKeyEnvAppInstance, val)

	return ctx
}
