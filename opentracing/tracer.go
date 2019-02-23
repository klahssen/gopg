package main

import (
	"io"
	"log"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	config "github.com/uber/jaeger-client-go/config"
)

/*
func registerServerTracer() error {
	//use default port (6831) and default max packet length (65000)
	sender, err := jaeger.NewUDPTransport("", 0)
	if err != nil {
		return err
	}
	cfg := config.Configuration{}
	metrics := jaeger.NewNullMetrics()
	logger := jaeger.StdLogger
	reportLogger := jaeger.ReporterOptions.Logger(logger)
	reportMetrics := jaeger.ReporterOptions.Metrics(metrics)
	reporter := jaeger.NewRemoteReporter(sender, reportMetrics, reportLogger)
	sampler := jaeger.NewConstSampler(true)
	tracer, _ := jaeger.NewTracer(
		"example_server",
		sampler,
		reporter,
	)
	opentracing.SetGlobalTracer(tracer)
	return nil
}
*/

func registerServerTracer() error {
	//use default port (6831) and default max packet length (65000)
	sender, err := jaeger.NewUDPTransport("", 0)
	if err != nil {
		return err
	}
	metrics := jaeger.NewNullMetrics()
	logger := jaeger.StdLogger
	reportLogger := jaeger.ReporterOptions.Logger(logger)
	reportMetrics := jaeger.ReporterOptions.Metrics(metrics)
	reporter := jaeger.NewRemoteReporter(sender, reportMetrics, reportLogger)
	sampler := jaeger.NewConstSampler(true)
	tracer, _ := jaeger.NewTracer(
		"example_server",
		sampler,
		reporter,
	)
	opentracing.SetGlobalTracer(tracer)
	return nil
}

// newTracer returns a new tracer
func newTracer(serviceName, hostPort string) (opentracing.Tracer, io.Closer) {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  hostPort, // localhost:5775
		},
	}
	tracer, closer, err := cfg.New(
		serviceName,
		config.Logger(jaeger.StdLogger),
	)
	if err != nil {
		log.Fatal(err)
	}

	return tracer, closer
}
