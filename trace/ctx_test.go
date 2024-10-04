package trace_test

import (
	"context"
	"log"
	"testing"

	"github.com/kontora13-go/coreutil/trace"
)

const ctxKeyTrace = "trace_id"

func TestTraceSetContextValues(t *testing.T) {
	ctx := context.Background()

	ctx = trace.WithTrace(ctx, "trace-id")
	log.Printf("ctx: %v", ctx.Value(ctxKeyTrace))
	c, ok := trace.TraceId(ctx)
	if !ok {
		t.Error("нет trace_id в context")
	}
	log.Printf("trace_id: %v", c)
	log.Println("---")

}
