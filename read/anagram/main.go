package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
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
	defer f.Close()
	scann := bufio.NewScanner(f)
	text := ""
	e := []entry{}
	ind := 1
	ok := true
	s := ""
	m := map[string][]entry{}
	for scann.Scan() {
		text = scann.Text()
		s = signature(text)
		if e, ok = m[s]; ok {
			e = append(e, entry{s: text, l: ind})
		} else {
			e = []entry{entry{s: text, l: ind}}
		}
		ind++
	}
	for k, vals := range m {
		fmt.Printf("found %d anagrams of '%s': %v\n", len(vals), k, vals)
	}
	fmt.Printf("treated in %s\n", time.Since(t0))
}

type entry struct {
	s string
	l int
}

func signature(w string) string {
	l := make([]string, len(w))
	for i, r := range w {
		l[i] = string(r)
	}
	sort.Strings(l)
	return strings.Join(l, "")
}
