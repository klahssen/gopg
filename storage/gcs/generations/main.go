package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"reflect"
	"sync"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
)

const (
	bktName = "gcstratyo-playground"
)

func main() {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("------- list files ...")
	if err = list(ctx, client, "", "/"); err != nil {
		log.Println(err)
		return
	}
	log.Println("------- tryReadWhileWriting ...")
	tryReadWhileWriting(ctx, client, "file1")
	log.Println("------- tryCreateDeleteRecreate ...")
	tryCreateDeleteRecreate(ctx, client, "file1")
	log.Println("------- createAndTryDeleteWithOpenedReader ...")
	createAndTryDeleteWithOpenedReader(ctx, client, "file1")
}
func list(ctx context.Context, client *storage.Client, prefix string, delimiter string) error {
	q := &storage.Query{Prefix: prefix, Delimiter: delimiter, Versions: true}
	it := client.Bucket(bktName).Objects(ctx, q)
	for {
		at, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		log.Println(at.Name, at.Generation)
	}
	return nil
}
func tryReadWhileWriting(ctx context.Context, client *storage.Client, filename string) {
	var err error
	if err = writeNew(ctx, client, filename, []byte("content1")); err != nil {
		log.Println(err)
		return
	}
	time.Sleep(time.Millisecond * 100)
	wg := sync.WaitGroup{}
	wg.Add(2)
	cErr := make(chan error, 10)
	go func() {
		//get reader
		defer wg.Done()
		if err := read(ctx, client, filename); err != nil {
			cErr <- err
		}
	}()
	go func() {
		time.Sleep(time.Millisecond * 100)
		//get writer and write 2nd gen
		defer wg.Done()
		if err := write(ctx, client, filename, []byte("content2")); err != nil {
			cErr <- err
		}
	}()

	wg.Wait()
	time.Sleep(time.Millisecond * 100)
	if err = read(ctx, client, filename); err != nil {
		cErr <- err
	}
	close(cErr)
	for err = range cErr {
		log.Printf("err: %v", err)
	}
}

func tryCreateDeleteRecreate(ctx context.Context, client *storage.Client, filename string) {
	var err error
	if err = writeNew(ctx, client, filename, []byte("content1")); err != nil {
		log.Println(err)
		return
	}
	if err = delete(ctx, client, filename); err != nil {
		log.Println(err)
		return
	}
	if err = writeNew(ctx, client, filename, []byte("content2")); err != nil {
		log.Println(err)
		return
	}
	if err = read(ctx, client, filename); err != nil {
		log.Println(err)
		return
	}

}

func createAndTryDeleteWithOpenedReader(ctx context.Context, client *storage.Client, filename string) {
	var err error
	content := []byte("content1")
	if err = writeNew(ctx, client, filename, content); err != nil {
		log.Println(err)
		return
	}
	h1 := client.Bucket(bktName).Object(filename)
	r, err := h1.NewReader(ctx)
	if err != nil {
		log.Printf("failed to get new reader: %v", err)
		return
	}
	defer r.Close()
	if err = delete(ctx, client, filename); err != nil {
		log.Printf("failed to delete: %v", err)
		return
	}
	b := make([]byte, len(content))
	if _, err := r.Read(b); err != nil {
		log.Printf("failed to read: %v", err)
		return
	}
	log.Printf("%q", b)
}

func delete(ctx context.Context, client *storage.Client, filename string) error {
	h := client.Bucket(bktName).Object(filename)
	return h.Delete(ctx)
}

func read(ctx context.Context, client *storage.Client, filename string) error {
	h := client.Bucket(bktName).Object(filename)
	a, err := h.Attrs(ctx)
	if err != nil {
		return fmt.Errorf("reader: get attributes: %v", err)
	}
	log.Printf("reader: object: '%s', generation: %d", a.Name, a.Generation)
	r, err := h.NewReader(ctx)
	if err != nil {
		return fmt.Errorf("new reader: %v", err)
	}
	defer r.Close()
	p := make([]byte, 100)
	var n int
	for {
		n, err = r.Read(p)
		if err != nil && err != io.EOF {
			return fmt.Errorf("reader: r.Read: %v", err)
		}
		if n == 0 || err == io.EOF {
			return nil
		}
		log.Printf("reader: read %q", p[:n])
	}
}

func writeNew(ctx context.Context, client *storage.Client, filename string, content []byte) error {
	h := client.Bucket(bktName).Object(filename)
	w := h.NewWriter(ctx)
	log.Printf("write new file '%s' in bucket '%s'", filename, bktName)
	_, err := w.Write(content)
	if err != nil {
		return fmt.Errorf("writer: w.Write: %v", err)
	}
	w.Close()
	r, err := h.NewReader(ctx)
	if err != nil {
		return fmt.Errorf("writer: failed to create read after write: %v", err)
	}
	p2 := make([]byte, len(content))
	_, err = r.Read(p2)
	if err != nil {
		return fmt.Errorf("writer: failed to read after write: %v", err)
	}
	defer r.Close()
	if !reflect.DeepEqual(p2, content) {
		return fmt.Errorf("writer: read %q expected %q", p2, content)
	}
	return nil
}

func write(ctx context.Context, client *storage.Client, filename string, content []byte) error {
	h := client.Bucket(bktName).Object(filename)
	a, err := h.Attrs(ctx)
	if err != nil {
		return fmt.Errorf("writer: get attributes: %v", err)
	}
	log.Printf("writer: object: '%s', generation: %d", a.Name, a.Generation)
	w := h.NewWriter(ctx)
	_, err = w.Write(content)
	if err != nil {
		return fmt.Errorf("writer: w.Write: %v", err)
	}
	w.Close()
	r, err := h.NewReader(ctx)
	if err != nil {
		return fmt.Errorf("writer: failed to create read after write: %v", err)
	}
	p2 := make([]byte, len(content))
	_, err = r.Read(p2)
	if err != nil {
		return fmt.Errorf("writer: failed to read after write: %v", err)
	}
	defer r.Close()
	if !reflect.DeepEqual(p2, content) {
		return fmt.Errorf("writer: read %q expected %q", p2, content)
	}
	return nil
}
