package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	subj = "hello"
)

func main() {
	f()
}

func f() {
	nPubsFlag := flag.Uint("p", 1, "number of publishers")
	nConsumersFlag := flag.Uint("c", 1, "number of consumers")
	nMsgsFlag := flag.Uint("m", 1000000, "number of messages")
	msgSizeFlag := flag.Uint("ms", 128, "number of publishers")
	flag.Parse()

	msg := make([]byte, int(*msgSizeFlag))
	wg1 := sync.WaitGroup{}
	wg1.Add(int(*nPubsFlag))
	for i := 1; i <= int(*nPubsFlag); i++ {
		go func(name string) {
			defer wg1.Done()
			target := int(*nMsgsFlag)
			nc, err := nats.Connect(nats.DefaultURL)
			if err != nil {
				log.Fatal(err)
			}
			t0 := time.Now()
			for j := 1; j < target; j++ {
				nc.Publish(subj, msg)
			}
			nc.Flush()
			dur := time.Since(t0)
			log.Printf("%s published %d messages in %s (%.1f/sec)", name, *nMsgsFlag, dur, float64(time.Second)/float64(dur)*float64(target))
		}(fmt.Sprintf("publisher %d", i))
	}
	wg2 := sync.WaitGroup{}
	wg2.Add(int(*nConsumersFlag))
	for i := 1; i <= int(*nConsumersFlag); i++ {
		go func(name string) {
			defer wg2.Done()
			nc, err := nats.Connect(nats.DefaultURL)
			if err != nil {
				log.Fatal(err)
			}
			n := 0

			t0 := time.Now()
			sub, _ := nc.SubscribeSync(subj)
			target := int(*nMsgsFlag) * int(*nPubsFlag)

			for {
				sub.NextMsg(time.Millisecond * 1)
				n++
				//log.Printf("%s received %d msg", name, n)
				if n == target {
					break
				}
			}

			/*
				done := make(chan int)
				nc.Subscribe(subj, func(m *nats.Msg) {
					n++
					//log.Printf("%s received %d msg", name, n)
					if n == int(*nMsgsFlag) {
						close(done)
					}
				})
				<-done
			*/
			dur := time.Since(t0)
			log.Printf("%s consumed %d messages in %s (%.1f/sec)", name, target, dur, float64(time.Second)/float64(dur)*float64(target))
		}(fmt.Sprintf("consumer %d", i))
	}
	wg1.Wait()
	wg2.Wait()
	//fmt.Printf("%d publisher(s) published %d messages each in %s", *nPubsFlag, *nMsgsFlag, time.Since(t0))
}
