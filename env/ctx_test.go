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

var (
	appZone       = env.ZoneTest
	appName       = "coreutil"
	appVersion, _ = env.NewVersion("0.0.1")
	appInstance   = "local"
)

// checkEnv проверяет значения в context на соответствие установленным в переменных
func checkEnv(ctx context.Context, t *testing.T) {
	if v, ok := env.AppZone(ctx); !ok {
		t.Errorf("нет значения env.Zone в context")
	} else if v != appZone {
		t.Errorf("ожидаемое значение ctx.Value(appZone) = %v, получено: %v", appZone, v)
	}
	if v, ok := env.AppName(ctx); !ok {
		t.Errorf("нет значения env.Name в context")
	} else if v != appName {
		t.Errorf("ожидаемое значение ctx.Value(appName) = %v, получено: %v", appName, v)
	}
	if v, ok := env.AppVersion(ctx); !ok {
		t.Errorf("нет значения env.Version в context")
	} else if v.String() != appVersion.String() {
		t.Errorf("ожидаемое значение ctx.Value(appVersion) = %v, получено: %v", appVersion, v)
	}
	if v, ok := env.AppInstance(ctx); !ok {
		t.Errorf("нет значения env.Instance в context")
	} else if v != appInstance {
		t.Errorf("ожидаемое значение ctx.Value(appInstance) = %v, получено: %v", appInstance, v)
	}
}

func TestEnvMultiContext(t *testing.T) {
	ctx := context.Background()

	// Устанавливаем значения в context
	ctx = env.WithEnv(ctx, appZone, appName, appVersion, appInstance)
	log.Printf("ctx: %v", ctx.Value(ctxKeyEnv))
	// Проверяем значения в context
	checkEnv(ctx, t)

	// Создаем новый контекст с новыми значениями
	ctx = context.WithValue(ctx, "test", 2)
	log.Printf("ctx: %v + %v", ctx.Value(ctxKeyEnv), ctx.Value("test"))
	// Проверяем значения в context
	checkEnv(ctx, t)

}

func TestEnvDoubleContext(t *testing.T) {
	ctx := context.Background()

	// Устанавливаем значения в context
	ctx1 := env.WithEnv(ctx, appZone, appName, appVersion, appInstance)
	log.Printf("ctx1: %v", ctx1.Value(ctxKeyEnv))
	log.Println("---")
	// Проверяем значения в context
	checkEnv(ctx1, t)

	appVersion, _ = env.NewVersion("0.0.2")

	ctx2 := env.WithEnv(ctx, appZone, appName, appVersion, appInstance)
	log.Printf("ctx1: %v", ctx1.Value(ctxKeyEnv))
	log.Printf("ctx2: %v", ctx2.Value(ctxKeyEnv))
	log.Println("---")
	// Проверяем значения в context
	checkEnv(ctx2, t)
}

func TestEnvSetContextValues(t *testing.T) {
	ctx := context.Background()

	ctx = env.SetAppName(ctx, appName)
	log.Printf("ctx: %v", ctx.Value(ctxKeyEnv))
	if v, ok := env.AppName(ctx); !ok {
		t.Error("нет AppName в context")
	} else if v != appName {
		t.Errorf("ожидаемое значение: %v, фактическое значение: %v", appName, v)
	} else {
		log.Printf("app_name: %v", v)

	}
	log.Println("---")

	ctx = env.SetAppZone(ctx, appZone)
	log.Printf("ctx: %v", ctx.Value(ctxKeyEnv))
	if v, ok := env.AppZone(ctx); !ok {
		t.Error("нет AppName в context")
	} else if v != appZone {
		t.Errorf("ожидаемое значение: %v, фактическое значение: %v", appZone, v)
	} else {
		log.Printf("app_zone: %v", v)

	}
	log.Println("---")

	ctx = env.SetAppInstance(ctx, appInstance)
	log.Printf("ctx: %v", ctx.Value(ctxKeyEnv))
	if v, ok := env.AppInstance(ctx); !ok {
		t.Error("нет AppName в context")
	} else if v != appInstance {
		t.Errorf("ожидаемое значение: %v, фактическое значение: %v", appInstance, v)
	} else {
		log.Printf("app_instance: %v", v)

	}
	log.Println("---")

	ctx = env.SetAppVersion(ctx, appVersion)
	log.Printf("ctx: %v", ctx.Value(ctxKeyEnv))
	if v, ok := env.AppVersion(ctx); !ok {
		t.Error("нет AppName в context")
	} else if v.String() != appVersion.String() {
		t.Errorf("ожидаемое значение: %v, фактическое значение: %v", appVersion, v)
	} else {
		log.Printf("app_version: %v", v)

	}
	log.Println("---")
}

func BenchmarkEnvSetContextValues(b *testing.B) {
	ctx := env.WithEnv(context.Background(), appZone, appName, appVersion, appInstance)
	for i := 0; i < b.N; i++ {
		ctx = env.SetAppName(ctx, appName)
		ctx = env.SetAppZone(ctx, appZone)
		ctx = env.SetAppInstance(ctx, appInstance)
		ctx = env.SetAppVersion(ctx, appVersion)
	}
}

func BenchmarkEnvGetContextValues(b *testing.B) {
	ctx := env.WithEnv(context.Background(), appZone, appName, appVersion, appInstance)
	for i := 0; i < b.N; i++ {
		_, _ = env.AppName(ctx)
		_, _ = env.AppZone(ctx)
		_, _ = env.AppInstance(ctx)
		_, _ = env.AppVersion(ctx)
	}
}
