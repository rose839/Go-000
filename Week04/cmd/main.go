package main

import (
	"context"
	pb "week04/api"
	"week04/internal/server"
)

func main() {
	service := InitializeOrderService(nil, nil)

	s := server.NewServer("127.0.0.1:8080")
	pb.RegisterSaveOrderServer(s, service)
	s.Start(context.Background())
}
