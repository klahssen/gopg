package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/oklog/run"
)

func main() {
	var g run.Group
	{
		ctx, cancel := context.WithCancel(context.Background())
		g.Add(func() error {
			fmt.Println("program waiting for ctrl+C to exit")
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("os signal: %s", sig)
			case <-ctx.Done():
				return ctx.Err()
			}
		}, func(error) {
			cancel()
		})
	}
	{
		ctx, cancel := context.WithCancel(context.Background())
		g.Add(func() error {
			fmt.Println("started listener")
			<-ctx.Done()
			fmt.Println("listener canceled")
			return fmt.Errorf("listener canceled")
		}, func(error) {
			cancel()
		})
	}
	fmt.Printf("The group was terminated with: %v\n", g.Run())
}
