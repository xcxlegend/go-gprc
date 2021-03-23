package service

import (
	"context"
	"fmt"
	"grpcdemo/pkg/pb"
	"io"
	"sync"
	"time"
)

var vals = sync.Map{}

type DemoServer struct {
}

func (ds *DemoServer) Get(ctx context.Context, req *pb.GetValRequest) (*pb.GetValReply, error) {
	val, _ := vals.Load(req.GetKey())
	return &pb.GetValReply{
		Key: req.Key,
		Val: val.(string),
	}, nil
}

func (ds *DemoServer) Put(ctx context.Context, req *pb.PutValRequest) (*pb.PutValReply, error) {
	vals.Store(req.Key, req.Val)
	return &pb.PutValReply{
		Ok: true,
	}, nil
}

func (ds *DemoServer) Stream(stream pb.Demo_StreamServer) error {
	ctx, done := context.WithCancel(context.Background())
	go func() {
		defer func() {
			done()
		}()
		for {
			data, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				return
			}
			fmt.Println(data.Msg)
		}
	}()
	t := time.Tick(time.Second)
LOOP:
	for {
		select {
		case <-t:
			stream.Send(&pb.Msg{Msg: "server message"})
		case <-ctx.Done():
			break LOOP
		}
	}
	fmt.Println("client close")
	return nil
}
