package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" //this will attach some endpoint that the profiler will query
)

func main() {
	http.HandleFunc("/", h)
	if err := http.ListenAndServe("localhost:8000", nil); err != nil {
		log.Fatalf("server crashed: %s", err.Error())
	}
}

func h(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hi")
}

/*
trigger profiles with go tool pprof

30 sec CPU profile:
go tool pprof http://localhost:8000/debug/pprof/profile

heap:
go tool pprof http://localhost:8000/debug/pprof/heap

goroutine blocking profile:
go tool pprof http://localhost:8000/debug/pprof/block

or the list of measures:
go tool pprof http://localhost:8000/debug/pprof/profile

*/
