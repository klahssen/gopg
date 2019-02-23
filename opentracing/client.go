package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/opentracing/opentracing-go"
)

func makeSomeRequest(ctx context.Context) error {
	span := opentracing.StartSpan("makeSomeRequest")
	defer span.Finish()
	//if span := opentracing.SpanFromContext(ctx); span != nil {
	httpClient := &http.Client{}
	httpReq, _ := http.NewRequest("GET", fmt.Sprintf("http://localhost:%d/", port), nil)
	// Transmit the span's TraceContext as HTTP headers on our
	// outbound request.
	opentracing.GlobalTracer().Inject(
		span.Context(),
		opentracing.HTTPHeaders,
		opentracing.HTTPHeadersCarrier(httpReq.Header))
	resp, err := httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	infoOTLog(ctx, string(b))
	return err
}
