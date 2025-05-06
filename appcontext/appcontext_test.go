package appcontext_test

import (
	"context"
	"testing"

	"github.com/kontora13-go/coreutil/appcontext"
)

func TestContextWithValue(t *testing.T) {
	var ctx context.Context
	ctxKeyEnvAppName := "app_name"
	ctxValEnvAppName := "coreutil"
	ctxKeyEnvAppZone := "app_zone"
	ctxValEnvAppZone := "prod"

	// WithValue / GetValue
	ctx = context.Background()
	ctx = appcontext.WithValue(ctx, map[string]interface{}{
		ctxKeyEnvAppName: ctxValEnvAppName,
		ctxKeyEnvAppZone: ctxValEnvAppZone,
	})

	if c, ok := appcontext.GetValue(ctx, ctxKeyEnvAppName); !ok {
		t.Errorf("нет %s в context", ctxKeyEnvAppName)
	} else if c != ctxValEnvAppName {
		t.Errorf("ожидаемое значение: %v, фактическое значение: %v", ctxValEnvAppName, c)
	}
	if c, ok := appcontext.GetValue(ctx, ctxKeyEnvAppZone); !ok {
		t.Errorf("нет %s в context", ctxKeyEnvAppZone)
	} else if c != ctxValEnvAppZone {
		t.Errorf("ожидаемое значение: %v, фактическое значение: %v", ctxValEnvAppZone, c)
	}
}

func TestContextSetValue(t *testing.T) {
	var ctx context.Context
	ctxKeyEnvAppName := "app_name"
	ctxValEnvAppName := "coreutil"
	ctxKeyEnvAppZone := "app_zone"
	ctxValEnvAppZone := "app_zone"

	// SetValue / GetValue
	ctx = context.Background()
	ctx = appcontext.WithValue(ctx, map[string]interface{}{
		ctxKeyEnvAppName: ctxValEnvAppName + "_test",
	})
	ctx = appcontext.SetValue(ctx, ctxKeyEnvAppName, ctxValEnvAppName)
	ctx = appcontext.SetValue(ctx, ctxKeyEnvAppZone, ctxValEnvAppZone)

	if c, ok := appcontext.GetValue(ctx, ctxKeyEnvAppName); !ok {
		t.Errorf("нет %s в context", ctxKeyEnvAppName)
	} else if c != ctxValEnvAppName {
		t.Errorf("ожидаемое значение: %v, фактическое значение: %v", ctxValEnvAppName, c)
	}
	if c, ok := appcontext.GetValue(ctx, ctxValEnvAppZone); !ok {
		t.Errorf("нет %s в context", ctxKeyEnvAppZone)
	} else if c != ctxValEnvAppZone {
		t.Errorf("ожидаемое значение: %v, фактическое значение: %v", ctxValEnvAppZone, c)
	}
}

func TestContextSetValueIfEmpty(t *testing.T) {
	var ctx context.Context
	ctxKeyEnvAppName := "app_name"
	ctxValEnvAppName := "coreutil"
	ctxKeyEnvAppZone := "app_zone"
	ctxValEnvAppZone := "app_zone"

	// SetValueIfEmpty / GetValue
	ctx = context.Background()
	ctx = appcontext.WithValue(ctx, map[string]interface{}{
		ctxKeyEnvAppName: ctxValEnvAppName,
	})
	ctx = appcontext.SetValueIfEmpty(ctx, ctxKeyEnvAppName, ctxValEnvAppName+"_test")
	ctx = appcontext.SetValueIfEmpty(ctx, ctxKeyEnvAppZone, ctxValEnvAppZone)

	if c, ok := appcontext.GetValue(ctx, ctxKeyEnvAppName); !ok {
		t.Errorf("нет %s в context", ctxKeyEnvAppName)
	} else if c != ctxValEnvAppName {
		t.Errorf("ожидаемое значение: %v, фактическое значение: %v", ctxValEnvAppName, c)
	}
	if c, ok := appcontext.GetValue(ctx, ctxValEnvAppZone); !ok {
		t.Errorf("нет %s в context", ctxKeyEnvAppZone)
	} else if c != ctxValEnvAppZone {
		t.Errorf("ожидаемое значение: %v, фактическое значение: %v", ctxValEnvAppZone, c)
	}
}
