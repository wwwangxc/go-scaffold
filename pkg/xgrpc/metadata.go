package xgrpc

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
)

var (
	ErrContextEmpty     = errors.New("context can't be empty.")
	ErrMetadataNotFound = errors.New("metadata not found.")
)

// GetMetadataFromIncomingContext return mdatadata from incoming context.
func GetMetadataFromIncomingContext(ctx context.Context) (metadata.MD, error) {
	if ctx == nil {
		return nil, ErrContextEmpty
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, ErrMetadataNotFound
	}
	return md, nil
}

// GetMetadataFromOutgoingContext return metadata from outgoing context.
func GetMetadataFromOutgoingContext(ctx context.Context) (metadata.MD, error) {
	if ctx == nil {
		return nil, ErrContextEmpty
	}
	md, ok := metadata.FromOutgoingContext(ctx)
	if !ok {
		return nil, ErrMetadataNotFound
	}
	return md, nil
}
