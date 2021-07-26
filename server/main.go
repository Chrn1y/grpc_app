package main

import (
	"context"
	"io"
	"net"

	"example.com/grpcapp/proto"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCServer struct {
	proto.UnimplementedAppServer
}

func (s *GRPCServer) Add(ctx context.Context, req *proto.AddValues) (*proto.Value, error) {
	return &proto.Value{X: req.GetX() + req.GetY()}, nil
}

func (s *GRPCServer) Sum(stream proto.App_SumServer) error {
	var res int64
	for {
		x, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.Value{
				X: res,
			})
		}
		if err != nil {
			return err
		}
		res += x.GetX()
	}
}

func (s *GRPCServer) Ones(x *proto.Value, stream proto.App_OnesServer) error {
	for it := x.GetX(); it > 0; it -= 1 {
		if err := stream.Send(&proto.Value{
			X: 1}); err != nil {
			return err
		}
	}
	return nil
}

func (s *GRPCServer) Repeat(stream proto.App_RepeatServer) error {
	return status.Errorf(codes.Unimplemented, "method Repeat not implemented")
}

func main() {
	s := grpc.NewServer()
	srv := &GRPCServer{}
	proto.RegisterAppServer(s, srv)
	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		return
	}
	if err := s.Serve(l); err != nil {
		return
	}
}
