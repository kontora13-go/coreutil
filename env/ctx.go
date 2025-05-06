// Copyright 2024-2025 Kontora13. All rights reserved.
// Licensed under the Apache License, Version 2.0

// Работа с env-данными с использованием контекста приложения

package env

import (
	"context"
)

const (
	// Ключи для хранения данных env-значений в context приложения
	ctxKeyEnv = "app_env"
)

type appEnv struct {
	appZone     Zone
	appName     string
	appVersion  Version
	appInstance string
}

func newAppEnv() *appEnv {
	return &appEnv{
		appZone:    ZoneNotDefined,
		appVersion: EmptyVersion,
	}
}

// WithEnv - получение context с env приложения
func WithEnv(
	ctx context.Context,
	appZone Zone, appName string, appVersion Version, appInstance string,
) context.Context {

	var appCtx *appEnv
	v := ctx.Value(ctxKeyEnv)
	if v != nil {
		var ok bool
		if appCtx, ok = v.(*appEnv); !ok {
			appCtx = newAppEnv()
			ctx = context.WithValue(ctx, ctxKeyEnv, appCtx)
		}
	} else {
		appCtx = newAppEnv()
		ctx = context.WithValue(ctx, ctxKeyEnv, appCtx)
	}

	appCtx.appZone = appZone
	appCtx.appName = appName
	appCtx.appVersion = appVersion
	appCtx.appInstance = appInstance

	return ctx
}

// AppZone - извлечение Zone приложения из контекста
func AppZone(ctx context.Context) (val Zone, ok bool) {
	c := ctx.Value(ctxKeyEnv)
	if c == nil {
		return ZoneNotDefined, false
	}

	appCtx, ok := c.(*appEnv)
	if !ok {
		return ZoneNotDefined, false
	}

	val = appCtx.appZone
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

	appCtx, ok := c.(*appEnv)
	if !ok {
		return WithEnv(ctx, val, "", EmptyVersion, "")
	}

	appCtx.appZone = val

	return ctx
}

// AppName - извлечение названия приложения из контекста
func AppName(ctx context.Context) (val string, ok bool) {
	c := ctx.Value(ctxKeyEnv)
	if c == nil {
		return
	}

	appCtx, ok := c.(*appEnv)
	if !ok {
		return
	}

	val = appCtx.appName

	return
}

// SetAppName - запись названия приложения в контекст
func SetAppName(ctx context.Context, val string) context.Context {
	c := ctx.Value(ctxKeyEnv)
	if c == nil {
		return WithEnv(ctx, ZoneNotDefined, val, EmptyVersion, "")
	}

	appCtx, ok := c.(*appEnv)
	if !ok {
		return WithEnv(ctx, ZoneNotDefined, val, EmptyVersion, "")
	}

	appCtx.appName = val

	return ctx
}

// AppVersion - извлечение версии приложения из контекста
func AppVersion(ctx context.Context) (val Version, ok bool) {
	c := ctx.Value(ctxKeyEnv)
	if c == nil {
		return
	}

	appCtx, ok := c.(*appEnv)
	if !ok {
		return
	}

	val = appCtx.appVersion

	return
}

// SetAppVersion - установка версии приложения в контекст
func SetAppVersion(ctx context.Context, val Version) context.Context {
	c := ctx.Value(ctxKeyEnv)
	if c == nil {
		return WithEnv(ctx, ZoneNotDefined, "", val, "")
	}

	appCtx, ok := c.(*appEnv)
	if !ok {
		return WithEnv(ctx, ZoneNotDefined, "", val, "")
	}

	appCtx.appVersion = val

	return ctx
}

func AppInstance(ctx context.Context) (val string, ok bool) {
	c := ctx.Value(ctxKeyEnv)
	if c == nil {
		return
	}

	appCtx, ok := c.(*appEnv)
	if !ok {
		return
	}

	val = appCtx.appInstance

	return
}

func SetAppInstance(ctx context.Context, val string) context.Context {
	c := ctx.Value(ctxKeyEnv)
	if c == nil {
		return WithEnv(ctx, ZoneNotDefined, "", EmptyVersion, val)
	}

	appCtx, ok := c.(*appEnv)
	if !ok {
		return WithEnv(ctx, ZoneNotDefined, "", EmptyVersion, val)
	}

	appCtx.appInstance = val

	return ctx
}
