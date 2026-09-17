package presence

import (
	"context"
	"github.com/younesbeheshti/gocast_game/contract/protogolang/presence"
	"github.com/younesbeheshti/gocast_game/param"
	"github.com/younesbeheshti/gocast_game/pkg/protobuf"
	"github.com/younesbeheshti/gocast_game/pkg/slice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn   *grpc.ClientConn
	client presence.PresenceServiceClient
}

func New(address string) (*Client, error) {
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:   conn,
		client: presence.NewPresenceServiceClient(conn),
	}, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c Client) GetPresence(ctx context.Context, req *param.GetPresenceRequest) (*param.GetPresenceResponse, error) {

	resp, err := c.client.GetPresence(ctx, &presence.GetPresenceRequest{UserIds: slice.MapFromUintToUint64(req.UserIDs)})
	if err != nil {
		return nil, err
	}

	return protobuf.MapGetPresenceResponseFromProtobuf(resp), nil
}
