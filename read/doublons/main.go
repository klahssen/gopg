package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	t0 := time.Now()
	filepath := flag.String("f", "", "filepath")
	flag.Parse()
	f, err := os.Open(*filepath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file: %s\n", err.Error())
		os.Exit(1)
	}
	scann := bufio.NewScanner(f)
	text := ""
	line := 0
	ind := 1
	ok := true
	m := map[string]int{}
	for scann.Scan() {
		text = scann.Text()
		if line, ok = m[text]; ok {
			fmt.Printf("line %d: '%s' doublon from line %d\n", ind, text, line)
			continue
		}
		m[text] = ind
		ind++
	}
	fmt.Printf("treated in %s\n", time.Since(t0))
}
