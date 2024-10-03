package env_test

import (
	"context"
	"log"
	"testing"

	"github.com/kontora13-go/coreutil/env"
)

const (
	ctxKeyEnv = "app_env"
)

func TestEnvMultiContext(t *testing.T) {
	ctx := context.Background()

	v, err := env.NewVersion("0.0.1")
	if err != nil {
		t.Error(err)
	}

	ctx1 := env.WithEnv(ctx,
		env.ZoneTest, "coreutil", v, "local")

	log.Printf("ctx1: %v", ctx1.Value(ctxKeyEnv))
	log.Println("---")

	ctx2 := context.WithValue(ctx1, "test2", 2)

	log.Printf("ctx1: %v + %v", ctx1.Value(ctxKeyEnv), ctx1.Value("test2"))
	log.Printf("ctx2: %v + %v", ctx2.Value(ctxKeyEnv), ctx2.Value("test2"))

}

func TestEnvDoubleContext(t *testing.T) {
	ctx := context.Background()

	v, err := env.NewVersion("0.0.1")
	if err != nil {
		t.Error(err)
	}

	ctx1 := env.WithEnv(ctx,
		env.ZoneTest, "coreutil", v, "local")

	log.Printf("ctx1: %v", ctx1.Value(ctxKeyEnv))
	log.Println("---")

	v, err = env.NewVersion("0.0.2")
	if err != nil {
		t.Error(err)
	}

	ctx2 := context.WithValue(ctx1, "test2", 2)
	ctx2 = env.WithEnv(ctx2,
		env.ZoneTest, "coreutil1", v, "local")

	log.Printf("ctx1: %v + %v", ctx1.Value(ctxKeyEnv), ctx1.Value("test2"))
	log.Printf("ctx2: %v + %v", ctx2.Value(ctxKeyEnv), ctx2.Value("test2"))
}

func TestEnvSetContextValues(t *testing.T) {
	ctx := context.Background()

	ctx = env.SetAppName(ctx, "coreutil")
	log.Printf("ctx: %v", ctx.Value(ctxKeyEnv))
	c, ok := env.AppName(ctx)
	if !ok {
		t.Error("нет app_name в context")
	}
	log.Printf("app_name: %v", c)
	log.Println("---")

	ctx = env.SetAppZone(ctx, env.ZoneTest)
	log.Printf("ctx: %v", ctx.Value(ctxKeyEnv))
	z, ok := env.AppZone(ctx)
	if !ok {
		t.Error("нет app_zone в context")
	}
	log.Printf("app_zone: %v", z.String())
	log.Println("---")

	ctx = env.SetAppInstance(ctx, "local")
	log.Printf("ctx: %v", ctx.Value(ctxKeyEnv))
	c, ok = env.AppInstance(ctx)
	if !ok {
		t.Error("нет app_instance в context")
	}
	log.Printf("app_instance: %v", c)
	log.Println("---")

	v, err := env.NewVersion("0.0.1")
	if err != nil {
		t.Error(err)
	}
	ctx = env.SetAppVersion(ctx, v)
	log.Printf("ctx: %v", ctx.Value(ctxKeyEnv))
	v, ok = env.AppVersion(ctx)
	if !ok {
		t.Error("нет app_zone в context")
	}
	log.Printf("app_version: %v", v.String())
	log.Println("---")
}
