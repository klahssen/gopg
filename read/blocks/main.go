package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	buffSize = 48
)

func main() {
	f := flag.String("f", "", "filepath")
	flag.Parse()
	file, err := os.Open(*f)
	if err != nil {
		log.Fatalf("failed to open file: %s", err.Error())
	}
	fmt.Printf("opened file '%s'\n", *f)
	buffer := make([]byte, buffSize)
	for {
		n, err := file.Read(buffer[:cap(buffer)])
		if err != nil && err != io.EOF {
			log.Fatalf("failed to read: %s", err.Error()) // or something more appropriate...
		}
		buffer = buffer[:n]

		// do something with the buffer up to the number of bytes read.  Don't use anything past n because it might be scratch.
		fmt.Printf("%s\n\n\n", string(buffer[:n]))
		if err == io.EOF {
			break
		}
	}
}
