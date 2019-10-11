package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	nats "github.com/nats-io/nats.go"
)

func main() {
	log.Println("connect")
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)
	//nc, err := nats.Connect("0.0.0.0:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	simplePublish(nc)
	simpleRequestResponse(nc)
	multipleConsumers(10, 1000000, nc)
}

func multipleConsumers(nConsumers, nMessages int, nc *nats.Conn) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg := sync.WaitGroup{}
	wg.Add(nConsumers)

	for i := 1; i <= nConsumers; i++ {
		go func(i int) {
			counter := 0
			defer wg.Done()
			sub, err := nc.Subscribe("foo", func(m *nats.Msg) {
				//log.Printf("[Consumer %02d] Received : %s", i, string(m.Data))
				counter++
			})
			if err != nil {
				log.Fatal(err)
			}
			<-ctx.Done()
			log.Printf("[Consumer %02d] received %d/%d messages (%.2f)", i, counter, nMessages, float64(counter)/float64(nMessages)*100.0)
			//_ = sub.Drain()
			_ = sub.Unsubscribe()
		}(i)
	}
	time.Sleep(time.Millisecond * 100)
	msg := ""
	t0 := time.Now()
	for j := 1; j <= nMessages; j++ {
		msg = fmt.Sprintf("message_%03d", j)
		if err := nc.Publish("foo", []byte(msg)); err != nil {
			log.Printf("failed to publish message '%s'", msg)
		}
		//nc.Flush()
		time.Sleep(time.Microsecond * 1)
	}
	//time.Sleep(time.Millisecond * 500)
	cancel()
	log.Printf("all messages published in %s", time.Since(t0))
	wg.Wait()
	//fmt.Println("counter", counter)
}

func simplePublish(nc *nats.Conn) {
	fmt.Println("")
	// Simple Async Subscriber
	sub, err := nc.Subscribe("foo", func(m *nats.Msg) {
		log.Printf("Received : %s", string(m.Data))
		m.Respond([]byte("response to message"))
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("publish to foo")
	// Simple Publisher
	if err := nc.Publish("foo", []byte("Hello World")); err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Millisecond * 1)
	nc.Flush()
	log.Println("unsubscribe")
	// Unsubscribe
	sub.Unsubscribe()
	log.Println("drain ...")
	// Drain
	sub.Drain()
}

func simpleRequestResponse(nc *nats.Conn) {
	fmt.Println("")
	// Simple Async Subscriber
	sub, err := nc.Subscribe("foo", func(m *nats.Msg) {
		log.Printf("request: %s", string(m.Data))
		m.Respond([]byte("fine thank you"))
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("send a request wait for response")
	// Simple Publisher
	timeout := time.Millisecond * 3
	m, err := nc.Request("foo", []byte("How are you?"), timeout)
	if err != nil {
		log.Fatal(err)
	}
	if m != nil {
		log.Printf("response: %s", string(m.Data))
	}
	log.Println("unsubscribe")
	// Unsubscribe
	sub.Unsubscribe()
	log.Println("drain ...")
	// Drain
	sub.Drain()
}

func f2(nc *nats.Conn) {
	/*
		// Simple Sync Subscriber
		sub, err := nc.SubscribeSync("foo")
		timeout := time.Second * 1
		m, err := sub.NextMsg(timeout)

		// Channel Subscriber
		ch := make(chan *nats.Msg, 64)
		sub, err := nc.ChanSubscribe("foo", ch)
		msg := <-ch
	*/

}
