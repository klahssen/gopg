package main

import (
	"context"
	"crypto/rand"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"encoding/base64"

	"github.com/klahssen/gopg/grpc/streaming/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct{}

var _ proto.CommandsAPIServer = (*server)(nil)

func (s *server) PushStream(stream proto.CommandsAPI_PushStreamServer) error {
	t0 := time.Now()
	cmds := make([]*proto.Command, 0, 100)
	dur := time.Millisecond * 1000
	timer := time.NewTimer(dur)
	type pkt struct {
		msg *proto.Command
		err error
	}
	c := make(chan *pkt, 1)
	defer close(c)
	defer log.Printf("processed in %s", time.Since(t0))
	for {
		go func() {
			cmd, err := stream.Recv()
			c <- &pkt{msg: cmd, err: err}
		}()
		select {
		case <-timer.C:
			return status.Errorf(codes.DeadlineExceeded, "timeout after %s", dur)
		case p := <-c:
			if p.err == io.EOF {
				//received all
				tx := make([]byte, 12)
				rand.Read(tx)
				txID := base64.StdEncoding.EncodeToString(tx)
				return stream.SendAndClose(&proto.Confirmation{TxID: txID})
			}
			//failure during reception
			if p.err != nil {
				return p.err
			}
			//log.Printf("new command: profileID='%s' segmentID=%#v", p.msg.ProfileID, p.msg.SegmentIDs)
			cmds = append(cmds, p.msg)
		}
	}
}
func (s *server) Push(ctx context.Context, payload *proto.Command) (*proto.Confirmation, error) {
	t0 := time.Now()
	defer log.Printf("processed in %s", time.Since(t0))
	tx := make([]byte, 12)
	rand.Read(tx)
	txID := base64.StdEncoding.EncodeToString(tx)
	return &proto.Confirmation{TxID: txID}, nil
}

func main() {
	port := flag.Int("port", 1234, "port on which to listen")
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterCommandsAPIServer(grpcServer, &server{})
	log.Printf("grpc server listening on port %d", *port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
