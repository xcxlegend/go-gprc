package service

import (
	"context"
	"grpcdemo/pkg/pb"
	"sync"
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
