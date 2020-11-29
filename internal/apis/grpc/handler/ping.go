package handler

import (
	"context"
	pingPB "go-scaffold/internal/apis/grpc/proto/ping"
)

type Ping struct {
}

func (t *Ping) Ping(ctx context.Context, arg *pingPB.PingRequest) (*pingPB.PingResponse, error) {
	return &pingPB.PingResponse{
		Message: "PONG",
	}, nil
}
