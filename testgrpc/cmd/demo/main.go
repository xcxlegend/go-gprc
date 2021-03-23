package main

import (
	"google.golang.org/grpc"
	"grpcdemo/internal/service"
	"grpcdemo/pkg/pb"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9900")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDemoServer(grpcServer, &service.DemoServer{})
	log.Fatal(grpcServer.Serve(lis))
}
