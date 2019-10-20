package main

import (
	"log"
	"sync"
	"time"

	nats "github.com/nats-io/nats.go"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(4)
	start := make(chan struct{})
	//subscribers
	go func() {
		defer wg.Done()
		if err := runSubscriber("worker1", start, time.Millisecond*200, "hello"); err != nil {
			log.Printf("worker1 failed: %v", err)
		}
	}()
	go func() {
		defer wg.Done()
		if err := runSubscriber("worker2", start, time.Millisecond*400, "hello"); err != nil {
			log.Printf("worker2 failed: %v", err)
		}
	}()
	go func() {
		defer wg.Done()
		if err := runSubscriber("worker3", start, time.Millisecond*600, "hello"); err != nil {
			log.Printf("worker3 failed: %v", err)
		}
	}()
	//publisher
	go func() {
		defer wg.Done()
		if err := runPublisher(start, "hello", 300*time.Millisecond); err != nil {
			log.Printf("publisher failed: %v", err)
		}
	}()
	time.Sleep(time.Second)
	close(start)
	wg.Wait()
}

func runPublisher(start chan struct{}, subj string, timeout time.Duration) error {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	<-start
	resp, err := nc.Request(subj, []byte("hello"), timeout)
	if err != nil {
		return err
	}
	log.Printf("received resp '%s'", string(resp.Data))
	nc.Flush()
	return nil
}

func runSubscriber(worker string, start chan struct{}, sleep time.Duration, subj string) error {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	sub, err := nc.SubscribeSync(subj)
	if err != nil {
		return err
	}
	<-start
	time.Sleep(sleep)
	msg, err := sub.NextMsg(time.Second * 1)
	if err != nil {
		return err
	}
	log.Printf("%s received '%s'", worker, string(msg.Data))
	return msg.Respond([]byte("hello from " + worker))
}
