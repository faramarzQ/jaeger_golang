package internal

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
)

// Foo processes a 3 seconds task
func Bar(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "bar-function")
	defer span.Finish()

	// a time consuming operation
	time.Sleep(3 * time.Second)
}
