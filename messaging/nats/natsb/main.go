// Copyright 2015-2019 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	//"github.com/nats-io/nats.go/bench"
)

// Some sane defaults
const (
	DefaultNumMsgs     = 100000
	DefaultNumPubs     = 1
	DefaultNumSubs     = 0
	DefaultMessageSize = 128
)

func usage() {
	log.Printf("Usage: nats-bench [-s server (%s)] [-np NUM_PUBLISHERS] [-ns NUM_SUBSCRIBERS] [-n NUM_MSGS] [-ms MESSAGE_SIZE] <subject>\n", nats.DefaultURL)
	flag.PrintDefaults()
}

func showUsageAndExit(exitcode int) {
	usage()
	os.Exit(exitcode)
}

func main() {
	var urls = flag.String("s", nats.DefaultURL, "The nats server URLs (separated by comma)")
	var numPubs = flag.Int("np", DefaultNumPubs, "Number of Concurrent Publishers")
	var numSubs = flag.Int("ns", DefaultNumSubs, "Number of Concurrent Subscribers")
	var numMsgs = flag.Int("n", DefaultNumMsgs, "Number of Messages to Publish")
	var msgSize = flag.Int("ms", DefaultMessageSize, "Size of the message.")
	var showHelp = flag.Bool("h", false, "Show help message")

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	if *showHelp {
		showUsageAndExit(0)
	}

	args := flag.Args()
	if len(args) != 1 {
		showUsageAndExit(1)
	}

	if *numMsgs <= 0 {
		log.Fatal("Number of messages should be greater than zero.")
	}

	//benchmark = bench.NewBenchmark("NATS", *numSubs, *numPubs)

	var startwg sync.WaitGroup
	var donewg sync.WaitGroup

	donewg.Add(*numPubs + *numSubs)

	// Run Subscribers first
	startwg.Add(*numSubs)
	for i := 0; i < *numSubs; i++ {
		nc, err := nats.Connect(*urls)
		if err != nil {
			log.Fatalf("Can't connect: %v\n", err)
		}
		defer nc.Close()

		go runSubscriber(strconv.Itoa(i+1), nc, &startwg, &donewg, *numMsgs, *msgSize)
	}
	startwg.Wait()

	// Now Publishers
	startwg.Add(*numPubs)
	//pubCounts := bench.MsgsPerClient(*numMsgs, *numPubs)
	for i := 0; i < *numPubs; i++ {
		nc, err := nats.Connect(*urls)
		if err != nil {
			log.Fatalf("Can't connect: %v\n", err)
		}
		defer nc.Close()

		//go runPublisher(nc, &startwg, &donewg, pubCounts[i], *msgSize)
		go runPublisher(strconv.Itoa(i+1), nc, &startwg, &donewg, (*numMsgs)/(*numPubs), *msgSize)
	}

	log.Printf("Starting benchmark [msgs=%d, msgsize=%d, pubs=%d, subs=%d]\n", *numMsgs, *msgSize, *numPubs, *numSubs)

	startwg.Wait()
	donewg.Wait()

	//benchmark.Close()

	//fmt.Print(benchmark.Report())
}

func runPublisher(name string, nc *nats.Conn, startwg, donewg *sync.WaitGroup, numMsgs int, msgSize int) {
	startwg.Done()

	args := flag.Args()
	subj := args[0]
	var msg []byte
	if msgSize > 0 {
		msg = make([]byte, msgSize)
	}

	start := time.Now()

	for i := 0; i < numMsgs; i++ {
		nc.Publish(subj, msg)
	}
	nc.Flush()
	dur := time.Since(start)
	//benchmark.AddPubSample(bench.NewSample(numMsgs, msgSize, start, time.Now(), nc))
	log.Printf("[PUBLISHER] %s: pub %d messages in %s (%.1fmsgs/sec)", name, numMsgs, dur, float64(numMsgs)*float64(time.Second)/float64(dur))
	donewg.Done()
}

func runSubscriber(name string, nc *nats.Conn, startwg, donewg *sync.WaitGroup, numMsgs int, msgSize int) {
	defer nc.Close()
	args := flag.Args()
	subj := args[0]

	received := 0
	ch := make(chan time.Time, 2)
	_, _ = nc.Subscribe(subj, func(msg *nats.Msg) {
		received++
		if received == 1 {
			ch <- time.Now()
		}
		if received >= numMsgs {
			ch <- time.Now()
		}
	})
	//sub.SetPendingLimits(-1, -1)
	nc.Flush()
	startwg.Done()

	<-ch //start
	start := time.Now()
	<-ch //end
	dur := time.Since(start)
	//benchmark.AddSubSample(bench.NewSample(numMsgs, msgSize, start, end, nc))
	log.Printf("[SUBSCRIBER] %s: consumed %d messages in %s (%.1fmsgs/sec)", name, numMsgs, dur, float64(numMsgs)*float64(time.Second)/float64(dur))
	donewg.Done()
}
