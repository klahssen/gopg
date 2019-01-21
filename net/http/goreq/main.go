package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	var b io.Reader
	url := flag.String("u", "", "url")
	body := flag.String("b", "", "body")
	meth := flag.String("m", "GET", "method")
	flag.Parse()
	if *url == "" {
		fmt.Fprint(os.Stderr, "empty url\n")
		os.Exit(1)
	}
	if *meth == "" {
		fmt.Fprint(os.Stderr, "empty method\n")
		os.Exit(1)
	}
	if *body != "" {
		b.Read([]byte(*body))
	}
	req, err := http.NewRequest(*meth, *url, b)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid request: %s\n", err.Error())
		os.Exit(1)
	}
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "request failed: %s\n", err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	rb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read response: %s\n", err.Error())
		os.Exit(1)
	}
	//reset the response body to the original unread state
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(rb))
	fmt.Printf("%s", string(rb))
}
