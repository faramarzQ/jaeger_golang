package internal

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
)

// Foo processes a 6 seconds task
func Foo(ctx context.Context) {
	span, _ := opentracing.StartSpanFromContext(ctx, "foo-function")
	defer span.Finish()

	// a time consuming operation
	time.Sleep(2 * time.Second)

	Bar(ctx)
	go Bar(ctx)

	// a time consuming operation
	time.Sleep(1 * time.Second)
}
