package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	nats "github.com/nats-io/nats.go"
)

const (
	nMessages   = 1000000
	nPublishers = 1
	nConsumers  = 1
	msgSize     = 256
)

func main() {
	log.Printf("benchmark: publish test [%d msgs, %d publisher(s), %d bytes/msg]", nMessages, nPublishers, msgSize)
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)
	//nc, err := nats.Connect("0.0.0.0:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	//workload chan
	c := make(chan struct{}, nMessages)
	//input
	go func() {
		for i := 1; i <= nMessages; i++ {
			c <- struct{}{}
		}
		close(c)
	}()
	wg := sync.WaitGroup{}
	wg.Add(nPublishers)
	//t0 := time.Now()
	for ip := 1; ip <= nPublishers; ip++ {
		go func(name string) {
			defer wg.Done()
			msg := make([]byte, msgSize)
			t0 := time.Now()
			nFailures := 0
			n := 0
			for range c {
				<-c
				_, err := nc.Request("hello", msg, time.Millisecond*5)
				if err != nil {
					nFailures++
				} else {
					n++
				}
			}
			dur := time.Since(t0)
			log.Printf("%s published %d req/resp in %s, %d failures", name, n, dur, nFailures)
		}(fmt.Sprintf("publisher_%d", ip))
	}
	for ic := 1; ic <= nConsumers; ic++ {
		go func() {
			//defer wg.Done()
			for range c {
				<-c
				_, _ = nc.Subscribe("hello", func(m *nats.Msg) {
					m.Respond([]byte("ok"))
				})
			}
		}()
	}
	wg.Wait()
	//dur := time.Since(t0)
	//log.Printf("published %d messages in %s (~%.1f/sec)", nMessages, dur, float64(time.Second)/float64(dur)*float64(nMessages))
}
