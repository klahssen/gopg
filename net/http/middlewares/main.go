package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(handler))
	http.Handle("/protected", authHeaderMiddleware(http.HandlerFunc(handler)))
	fmt.Printf("Listening on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hi!")
}
