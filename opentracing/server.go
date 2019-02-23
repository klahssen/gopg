package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

//startListening starts a local http.Server on specified port that can be closed through the context. It will take max 500ms to close it
func startListening(ctx context.Context, port int16) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	server := http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      mux,
		ReadTimeout:  time.Millisecond * 500,
		WriteTimeout: time.Millisecond * 1000,
		IdleTimeout:  time.Millisecond * 1500,
	}
	errChan := make(chan error, 1)

	go func() {
		if err := server.ListenAndServe(); err != nil {
			errChan <- err
		}
	}()
	var e error
	select {
	case e = <-errChan:
		e = fmt.Errorf("server: ListenAndServer: %v", e)
	case <-ctx.Done():
		infoOTLog(ctx, "server: received close signal, wait for all requests to be processed")
		infoOTLog(ctx, fmt.Sprintf("server: context.Err: %v", ctx.Err()))
		ctx0, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
		defer cancel()
		if e = server.Shutdown(ctx0); e != nil {
			e = fmt.Errorf("server: Shutdown: %v", e)
		}
	}
	return e
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var serverSpan opentracing.Span
	appSpecificOperationName := "handleRoot"
	wireContext, err := opentracing.GlobalTracer().Extract(
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(r.Header))
	if err != nil {
		// Optionally record something about err here
		errOTLog(ctx, fmt.Errorf("handleRoot: extract context from headers: err: %v", err))
	}
	infoOTLog(ctx, "extracted wire context")
	// Create the span referring to the RPC client if available.
	// If wireContext == nil, a root span will be created.
	serverSpan = opentracing.StartSpan(
		appSpecificOperationName,
		ext.RPCServerOption(wireContext))

	defer serverSpan.Finish()
	ctx = opentracing.ContextWithSpan(context.Background(), serverSpan)
	infoOTLog(ctx, "fetch...")
	fakeFetch(ctx)
	fmt.Fprintf(w, "hello")
}

func fakeFetch(ctx context.Context) {
	parentSpan := opentracing.SpanFromContext(ctx)
	sp := opentracing.StartSpan(
		"fakeFetch",
		opentracing.ChildOf(parentSpan.Context()))
	defer sp.Finish()
	ctxa := opentracing.ContextWithSpan(ctx, sp)
	infoOTLog(ctxa, "in fakeFetch")
	time.Sleep(time.Millisecond * 100)
	infoOTLog(ctx, "fetched")
}
