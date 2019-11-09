package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/klahssen/gopg/grpc/streaming/proto"
	"google.golang.org/grpc"
)

func getClient(addr string) (*grpc.ClientConn, proto.CommandsAPIClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}
	client := proto.NewCommandsAPIClient(conn)
	return conn, client, nil
}

func pushStream(client proto.CommandsAPIClient, nProfiles, nSegsPerProfile int) (string, error) {
	//t0 := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	stream, err := client.PushStream(ctx)
	if err != nil {
		return "", fmt.Errorf("")
	}
	var cmd *proto.Command
	segs := make([]string, 0, nSegsPerProfile)
	for j := 1; j <= nSegsPerProfile; j++ {
		segs = append(segs, fmt.Sprintf("sgmt_%03d", j))
	}
	for i := 1; i <= nProfiles; i++ {
		cmd = &proto.Command{ProfileID: fmt.Sprintf("usr_%03d", 1), SegmentIDs: segs}
		if err = stream.Send(cmd); err != nil {
			if err == io.EOF {
				break
			}
		}
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		return "", fmt.Errorf("CloseAndRecv: %v", err)
	}
	//log.Printf("pushed %d items in %s", n, time.Since(t0))
	return resp.TxID, nil
}

func push(client proto.CommandsAPIClient, nProfiles, nSegsPerProfile int) ([]string, error) {
	//t0 := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	var cmd *proto.Command
	txIDs := make([]string, 0, nProfiles)
	segs := make([]string, 0, nSegsPerProfile)
	for j := 1; j <= nSegsPerProfile; j++ {
		segs = append(segs, fmt.Sprintf("sgmt_%03d", j))
	}
	for i := 1; i <= nProfiles; i++ {
		cmd = &proto.Command{ProfileID: fmt.Sprintf("usr_%03d", 1), SegmentIDs: segs}
		if resp, err := client.Push(ctx, cmd); err != nil {
			if err != nil {
				return nil, fmt.Errorf("push item [%d]: err: %v", i, err)
			}
			txIDs = append(txIDs, resp.TxID)
		}
	}
	//log.Printf("pushed %d items in %s", n, time.Since(t0))
	return txIDs, nil
}
