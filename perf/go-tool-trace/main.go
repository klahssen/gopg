package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	tracefile := "trace.out"
	f, err := os.Create(tracefile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create file '%s': %s\n", tracefile, err.Error())
		os.Exit(1)
	}
	defer f.Close()
	if err = trace.Start(f); err != nil {
		fmt.Fprintf(os.Stderr, "failed to start tracing: %s\n", err.Error())
		os.Exit(1)
	}
	defer trace.Stop()
	fmt.Println("hello!")
}
