package main

import (
	"context"
	"log"

	"github.com/opentracing/opentracing-go"
	otlog "github.com/opentracing/opentracing-go/log"
)

//open tracing log
func infoOTLog(ctx context.Context, msg string) {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span.LogFields(otlog.String("info", msg))
	}
	log.Println(msg)
}

//open tracing log
func errOTLog(ctx context.Context, err error) {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span.LogFields(otlog.Error(err))
	}
	log.Println(err)
}

//open tracing log
func tagOTLog(ctx context.Context, key, value string) {
	if span := opentracing.SpanFromContext(ctx); span != nil {
		span.SetTag(key, value)
	}
}
