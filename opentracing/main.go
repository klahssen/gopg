package main

import (
	"context"
	"log"
	"time"
)

const (
	port int16 = 8080
)

func main() {
	err := registerServerTracer()
	if err != nil {
		log.Printf("failed to register server tracer: %v", err)
		return
	}
	//opentracing.InitGlobalTracer(dapperish.NewTracer("opentracing_tester"))
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*2000)
	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Millisecond*300)
	defer cancel()
	defer cancel2()
	errChan := make(chan error, 1)
	go func() {
		errChan <- startListening(ctx, port)
	}()
	go func() {
		if err := makeSomeRequest(ctx2); err != nil {
			log.Printf("makeSomeRequest: err: %v", err)
		} else {
			log.Printf("makeSomeRequest: done")
		}
	}()
	err = <-errChan
	log.Printf("main: closed server: err: %v", err)
	//runtime.Goexit()
}
