package main

import (
	"fmt"
	"net/http"
)

/*
WARNING:
The default http server has NO TIMEOUTS
*/

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Simple Helloworld server is running on port 8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
