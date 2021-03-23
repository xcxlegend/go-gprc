package test

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpcdemo/pkg/pb"
	"log"
	"testing"
	"time"
)

func TestGrpcClient(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:9900", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := pb.NewDemoClient(conn)
	//res, err := client.Put(context.Background(), &pb.PutValRequest{
	//	Key: "test",
	//	Val: "hello",
	//})
	//fmt.Println(err)
	//fmt.Println(res)
	//
	//res2, err := client.Get(context.Background(), &pb.GetValRequest{Key: "test"})
	//fmt.Println(err)
	//fmt.Println(res2)

	stream, err := client.Stream(context.Background())
	go func() {
		for {
			data, _ := stream.Recv()
			fmt.Println(data)
		}
	}()
	for {
		stream.Send(&pb.Msg{Msg: "client message"})
		time.Sleep(time.Second)
	}

}
