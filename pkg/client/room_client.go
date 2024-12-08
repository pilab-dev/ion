package client

import (
	"fmt"

	roomv1 "github.com/pilab-dev/ion/pkg/gen/pb/room/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// RoomClient represents a RoomClient which manage peers
type RoomClient struct {
	// grpc room service
	roomv1.RoomServiceClient
}

// NewRoomClient create a room client instance.
func NewRoomClient(addr string) (*RoomClient, error) {
	dialer, err := grpc.NewClient(fmt.Sprintf("dns:%s", addr), grpc.WithTransportCredentials(
		insecure.NewCredentials(),
	))
	if err != nil {
		return nil, fmt.Errorf("failed to dial: %v", err)
	}

	return &RoomClient{
		roomv1.NewRoomServiceClient(dialer),
	}, nil
}
