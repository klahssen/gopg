package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

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
			return cancellable(ctx)
		}, func(error) {
			cancel()
		})
	}
	fmt.Printf("The group was terminated with: %v\n", g.Run())
}

func cancellable(ctx context.Context) error {
	fmt.Println("started cancellable job")
	t := time.NewTimer(time.Millisecond * 1500)
	select {
	case <-t.C:
		fmt.Println("job done!")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
