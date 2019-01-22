package main

import (
	"time"
)

func get() {
	//fmt.Printf("called get\n")
	iterate(4)
}

func iterate(n int) {
	//fmt.Printf("do %d iterations\n", n)
	for i := 0; i < n; i++ {
		wait(time.Nanosecond * 12)
	}
}

func wait(dur time.Duration) {
	//fmt.Printf("wait %s\n", dur)
	time.Sleep(dur)
}
