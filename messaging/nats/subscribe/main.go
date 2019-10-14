package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	nats "github.com/nats-io/nats.go"
)

const (
	nMessages  = 1000000
	nConsumers = 3
	//nsgSize    = 256
)

func main() {
	//log.Printf("benchmark test [%d msgs, %d consumer(s), %d bytes/msg]", nMessages, nConsumers, msgSize)
	log.Printf("benchmark: consume test [%d msgs, %d consumer(s)]", nMessages, nConsumers)
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)
	//nc, err := nats.Connect("0.0.0.0:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	wg := sync.WaitGroup{}
	wg.Add(nConsumers)
	t0 := time.Now()
	for ic := 1; ic <= nConsumers; ic++ {
		go func(name string) {
			defer wg.Done()
			done := make(chan struct{})
			c := make(chan struct{}, nMessages)
			go func() {
				n := 0
				for {
					<-c
					n++
					log.Printf("%s received %d msgs", name, n)
					if n >= nMessages {
						done <- struct{}{}
						return
					}
				}
			}()
			_, _ = nc.Subscribe("hello", func(m *nats.Msg) {
				c <- struct{}{}
				//log.Printf("%s received %d msgs", name, n)
			})
			<-done
		}(fmt.Sprintf("worker_%d", ic))
	}
	wg.Wait()
	dur := time.Since(t0)
	log.Printf("%d consumer(s) each consumed %d messages in %s (~%.1f/sec)", nConsumers, nMessages, dur, float64(time.Second)/float64(dur)*float64(nMessages))
}
