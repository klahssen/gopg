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
	//pubsub()
	reqresp()
}

func pubsub() {
	protoFlag := flag.Bool("proto", false, "send protobuf msg")
	nPubsFlag := flag.Uint("p", 1, "number of publishers")
	nConsumersFlag := flag.Uint("c", 1, "number of consumers")
	nMsgsFlag := flag.Uint("m", 1000000, "number of messages")
	msgSizeFlag := flag.Uint("ms", 128, "number of publishers")
	flag.Parse()

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
	var msg []byte
	if *protoFlag {
		msg = Hello{}
	} else {
		msg = make([]byte, int(*msgSizeFlag))
	}
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
	wg1.Wait()
	wg2.Wait()
	//fmt.Printf("%d publisher(s) published %d messages each in %s", *nPubsFlag, *nMsgsFlag, time.Since(t0))
}

func reqresp() {
	nPubsFlag := flag.Uint("p", 1, "number of publishers")
	nConsumersFlag := flag.Uint("c", 1, "number of consumers")
	nMsgsFlag := flag.Uint("m", 1000000, "number of messages")
	msgSizeFlag := flag.Uint("ms", 128, "number of publishers")
	flag.Parse()

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
				m, _ := sub.NextMsg(time.Millisecond * 100)
				m.Respond([]byte("world"))
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
			j := 0
			for j < target {
				//nc.Publish(subj, msg)
				_, err := nc.Request(subj, msg, time.Second*2)
				/*if err != nil {
					log.Printf("[ERROR] %v", err)
				} else if resp == nil {
					log.Printf("[ERROR] response is nil")
				} else {
					log.Println(string(resp.Data))
				}*/
				if err == nil {
					j++
				}
			}
			nc.Flush()
			dur := time.Since(t0)
			log.Printf("%s published %d messages in %s (%.1f/sec)", name, *nMsgsFlag, dur, float64(time.Second)/float64(dur)*float64(target))
		}(fmt.Sprintf("publisher %d", i))
	}
	wg1.Wait()
	wg2.Wait()
	//fmt.Printf("%d publisher(s) published %d messages each in %s", *nPubsFlag, *nMsgsFlag, time.Since(t0))
}

func f2() {
	nPubsFlag := flag.Uint("p", 1, "number of publishers")
	nConsumersFlag := flag.Uint("c", 1, "number of consumers")
	nMsgsFlag := flag.Uint("m", 1000000, "number of messages")
	msgSizeFlag := flag.Uint("ms", 128, "number of publishers")
	flag.Parse()

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
			//sub, _ := nc.SubscribeSync(subj)
			target := int(*nMsgsFlag) * int(*nPubsFlag)

			done := make(chan int)
			nc.Subscribe(subj, func(m *nats.Msg) {
				n++
				//log.Printf("%s received %d msg", name, n)
				if n == target {
					close(done)
				}
			})
			<-done

			dur := time.Since(t0)
			log.Printf("%s consumed %d messages in %s (%.1f/sec)", name, target, dur, float64(time.Second)/float64(dur)*float64(target))
		}(fmt.Sprintf("consumer %d", i))
	}

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
	wg1.Wait()
	wg2.Wait()
	//fmt.Printf("%d publisher(s) published %d messages each in %s", *nPubsFlag, *nMsgsFlag, time.Since(t0))
}
