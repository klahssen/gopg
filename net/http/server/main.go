package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	port := "8080"
	pport := flag.String("p", "", "port on which the server will listen")
	flag.Parse()
	if *pport == "" {
		if p := os.Getenv("SERVER_PORT"); p != "" {
			//fmt.Printf("using SERVER_PORT environment variable\n")
			port = p
		} else {
			//fmt.Printf("using default PORT\n")
		}
	} else {
		//fmt.Printf("using PORT from -p flag\n")
		port = *pport
	}
	p := 0
	var err error
	if p, err = strconv.Atoi(port); err != nil {
		panic(fmt.Sprintf("invalid port '%s': %s", port, err.Error()))
	}

	log.Printf("Listening on port %d", p)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", p), getMux()); err != nil {
		panic(err)
	}
}

func getMux() http.Handler {
	mux := http.DefaultServeMux
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello\n"))
}
