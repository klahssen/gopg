package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	path := flag.String("f", "", "filepath")
	flag.Parse()
	if *path == "" {
		fmt.Fprint(os.Stderr, "empty filepath\n")
		os.Exit(1)
	}
	b, err := ioutil.ReadFile(*path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read file '%s': %s\n", *path, err.Error())
		os.Exit(1)
	}
	dest := data{}
	if err := json.Unmarshal(b, &dest); err != nil {
		fmt.Fprintf(os.Stderr, "failed to unmarshal '%s': %s\n", *path, err.Error())
		os.Exit(1)
	}
	dest.Read = true
	fmt.Printf("Read: %+v\n", dest)
}

type data struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
	Read    bool     `json:"read"`
}
