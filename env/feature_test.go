package env_test

import (
	"context"
	"fmt"
	"github.com/kontora13-go/coreutil/env"
	"testing"
)

func TestFeatures(t *testing.T) {
	ctx := context.Background()
	ctx = env.WithFeatures(ctx, map[string]bool{"feature1": true})

	if !env.CheckFeature(ctx, "feature1") {
		t.Error("неверно значение feature1")
	}

	ctx = env.SetFeature(ctx, "feature2", true)
	if !env.CheckFeature(ctx, "feature2") {
		t.Error("неверно значение feature2")
	}
	ctx = env.SetFeature(ctx, "feature2", false)
	if env.CheckFeature(ctx, "feature2") {
		t.Error("неверно значение feature2")
	}

	ctx = env.SetFeatureIfEmpty(ctx, "feature1", false)
	if !env.CheckFeature(ctx, "feature1") {
		t.Error("неверно значение feature1")
	}
}

func BenchmarkFeaturesSetValues(b *testing.B) {

	ctx := context.Background()
	ctx = env.WithFeatures(ctx, map[string]bool{"feature": true})
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ctx = env.SetFeature(ctx, fmt.Sprintf("feature%d", i), true)
	}
}

func BenchmarkFeaturesCheckValues(b *testing.B) {

	ctx := context.Background()
	ctx = env.WithFeatures(ctx, map[string]bool{"feature": true})

	for i := 0; i < b.N; i++ {
		ctx = env.SetFeature(ctx, fmt.Sprintf("feature%d", i), true)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = env.CheckFeature(ctx, fmt.Sprintf("feature%d", i))
	}
}
