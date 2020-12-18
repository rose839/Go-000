package server

import (
	"context"
	"net"

	"google.golang.org/grpc"
)

// Server represent a grpc server.
type Server struct {
	*grpc.Server

	addr     string
	grpcOpts []grpc.ServerOption
}

// NewServer is used to create a new grpc server.
func NewServer(addr string, opts ...grpc.ServerOption) *Server {
	srv := &Server{
		addr:     addr,
		grpcOpts: opts,
	}

	srv.Server = grpc.NewServer(opts...)
	return srv
}

// Start is used to start listen on specific addr.
func (s *Server) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	return s.Serve(lis)
}
