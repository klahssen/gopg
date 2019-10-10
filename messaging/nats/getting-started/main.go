package main

import (
	"fmt"
	"log"
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
	nc.Publish("foo", []byte("Hello World"))
	time.Sleep(time.Millisecond * 1)
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
