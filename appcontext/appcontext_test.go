package appcontext_test

import (
	"context"
	"log"
	"testing"

	"github.com/kontora13-go/coreutil/appcontext"
	"github.com/kontora13-go/coreutil/env"
)

const (
	ctxKeyValue = "app_value"
)

func TestValueSetContextValues(t *testing.T) {
	ctx := context.Background()

	ctxKeyEnvAppName := "app_name"
	ctxKeyEnvAppZone := "app_zone"

	ctx = appcontext.SetValue(ctx, ctxKeyEnvAppName, "coreutil")
	log.Printf("ctx: %v", ctx.Value(ctxKeyValue))
	c, ok := appcontext.GetValue(ctx, ctxKeyEnvAppName)
	if !ok {
		t.Error("нет app_name в context")
	}
	log.Printf("app_name: %v", c)
	log.Println("---")

	ctx = appcontext.SetValueIfEmpty(ctx, ctxKeyEnvAppZone, env.ZoneTest)
	log.Printf("ctx: %v", ctx.Value(ctxKeyValue))
	z, ok := appcontext.GetValue(ctx, ctxKeyEnvAppZone)
	if !ok {
		t.Error("нет app_zone в context")
	}
	log.Printf("app_zone: %v", z)
	log.Println("---")

	ctx = appcontext.SetValueIfEmpty(ctx, ctxKeyEnvAppName, "coreutil1")
	log.Printf("ctx: %v", ctx.Value(ctxKeyValue))
	c, ok = appcontext.GetValue(ctx, ctxKeyEnvAppName)
	if !ok {
		t.Error("нет app_name в context")
	}
	log.Printf("app_name: %v", c)
	log.Println("---")

}
