package handler

import (
	"context"
	"go-scaffold/internal/grpc/pb"
)

type Ping struct {
}

func (t *Ping) Ping(ctx context.Context, arg *pb.PingRequest) (*pb.PingResponse, error) {
	return &pb.PingResponse{
		Message: "PONG",
	}, nil
}
