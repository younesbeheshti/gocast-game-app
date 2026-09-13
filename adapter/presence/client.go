package presence

import (
	"context"
	"github.com/younesbeheshti/gocast_game/contract/golang/presence"
	"github.com/younesbeheshti/gocast_game/param"
	"github.com/younesbeheshti/gocast_game/pkg/protobuf"
	"github.com/younesbeheshti/gocast_game/pkg/slice"
	"google.golang.org/grpc"
)

type Client struct {
	client presence.PresenceServiceClient
}

func New(conn *grpc.ClientConn) Client {
	return Client{
		client: presence.NewPresenceServiceClient(conn),
	}
}

func (c Client) GetPresence(ctx context.Context, req *param.GetPresenceRequest) (*param.GetPresenceResponse, error) {
	resp, err := c.client.GetPresence(ctx, &presence.GetPresenceRequest{UserIds: slice.MapFromUintToUint64(req.UserIDs)})
	if err != nil {
		return nil, err
	}

	return protobuf.MapGetPresenceResponseFromProtobuf(resp), nil
}
