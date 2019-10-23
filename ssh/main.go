package main

import (
	"io/ioutil"
	"log"
	"net"
	"reflect"
	"sort"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

func main() {
	b, err := ioutil.ReadFile("./keys/id_rsa_nopass")
	if err != nil {
		log.Fatalf("failed to read private key file: %v", err)
	}
	signer, err := ssh.ParsePrivateKey(b)
	if err != nil {
		log.Fatalf("failed to parse private key: %v", err)
	}
	noopCallback := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}
	config := &ssh.ClientConfig{
		User: "test",
		Auth: []ssh.AuthMethod{
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: noopCallback,
	}
	// Connect to the remote server and perform the SSH handshake.
	client, err := ssh.Dial("tcp", "127.0.0.1:22", config)
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	sc, err := sftp.NewClient(client)
	if err != nil {
		log.Fatalf("failed to get new sftp client from ssh connexion: %v", err)
	}
	defer client.Close()
	path := "/incoming/file3.txt"
	f, err := sc.Create(path)
	if err != nil {
		log.Fatalf("failed to create file: %v", err)
	}
	content := []byte("file2")
	f.Write(content)
	f.Close()
	f2, err := sc.Open(path)
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	dest := make([]byte, len(content))
	f2.Read(dest)
	f2.Close()
	if !reflect.DeepEqual(dest, content) {
		log.Fatalf("mismatch between file writtent and read")
	}
	if err := sc.Remove(path); err != nil {
		log.Fatalf("failed to remove file after creation: %v", err)
	}
	infos, err := sc.ReadDir("/") // /incoming == /incoming/ == incoming == incoming/
	if err != nil {
		log.Fatalf("failed to read dir '/': %v", err)
	}
	for _, info := range infos {
		log.Printf("folder: %v, '%s'", info.IsDir(), info.Name())
	}

	infos, err = sc.ReadDir("incoming") // /incoming == /incoming/ == incoming == incoming/
	if err != nil {
		log.Fatalf("failed to read dir 'incoming': %v", err)
	}
	files := make([]string, 0, len(infos))
	for _, info := range infos {
		if !info.IsDir() {
			files = append(files, info.Name())
		}
	}
	sort.Strings(files)
	expected := []string{
		"file1.txt", "file2.txt",
	}
	if !reflect.DeepEqual(files, expected) {
		log.Fatalf("expected files %#v found %#v", expected, files)
	}
}
