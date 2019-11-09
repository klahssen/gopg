package main

import (
	"log"
	"testing"
)

func BenchmarkPushStream100x100(b *testing.B) {
	conn, client, err := getClient("127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for i := 1; i <= b.N; i++ {
		pushStream(client, 100, 100)
	}
}

func BenchmarkUnaryPush100x100(b *testing.B) {
	conn, client, err := getClient("127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for i := 1; i <= b.N; i++ {
		push(client, 100, 100)
	}
}

func BenchmarkPushStream1000x1000(b *testing.B) {
	conn, client, err := getClient("127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for i := 1; i <= b.N; i++ {
		pushStream(client, 1000, 1000)
	}
}

func BenchmarkUnaryPush1000x1000(b *testing.B) {
	conn, client, err := getClient("127.0.0.1:1234")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for i := 1; i <= b.N; i++ {
		push(client, 1000, 1000)
	}
}
