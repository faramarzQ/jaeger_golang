package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"jaeger_golang/internal"
	"jaeger_golang/jaeger"

	"github.com/opentracing/opentracing-go"
)

func main() {
	tracer, closer := jaeger.InitializeJaeger()
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc("/process", processHandler)

	log.Printf("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// An API handler executing a multi second task
func processHandler(w http.ResponseWriter, r *http.Request) {
	span := jaeger.Tracer.StartSpan("process-function")
	defer span.Finish()

	ctx := opentracing.ContextWithSpan(context.Background(), span)

	// a time consuming operation
	time.Sleep(1 * time.Second)

	internal.Foo(ctx)

	// a time consuming operation
	time.Sleep(1 * time.Second)

	w.Write([]byte("Request processed successfully."))
}
