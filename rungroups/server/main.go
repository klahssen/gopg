package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/oklog/run"
)

func main() {
	var g run.Group
	{
		ln, _ := net.Listen("tcp", ":0")
		g.Add(func() error {
			defer fmt.Printf("http.Serve returned\n")
			return http.Serve(ln, http.NewServeMux())
		}, func(error) {
			ln.Close()
		})
	}
	{
		dur := time.Millisecond * 500
		ctx, cancel := context.WithTimeout(context.Background(), dur)
		g.Add(func() error {
			fmt.Printf("started timer of %s", dur)
			<-ctx.Done()
			return ctx.Err()
		}, func(error) {
			cancel()
		})
	}
	fmt.Printf("The group was terminated with: %v\n", g.Run())
}
