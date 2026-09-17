package presenceserver

import (
	"context"
	"fmt"
	"github.com/younesbeheshti/gocast_game/contract/protogolang/presence"
	"github.com/younesbeheshti/gocast_game/param"
	"github.com/younesbeheshti/gocast_game/pkg/protobuf"
	"github.com/younesbeheshti/gocast_game/pkg/slice"
	"github.com/younesbeheshti/gocast_game/service/presenceservice"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	presence.UnimplementedPresenceServiceServer
	svc presenceservice.Service
}

func New(svc presenceservice.Service) Server {
	return Server{
		svc:                                svc,
		UnimplementedPresenceServiceServer: presence.UnimplementedPresenceServiceServer{},
	}
}

func (s Server) Start() {
	addr := fmt.Sprintf(":%d", 8000)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	presence.RegisterPresenceServiceServer(grpcServer, s)

	fmt.Println("Starting gRPC Server", addr)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal("Failed to start server")
	}
}

func (s Server) GetPresence(ctx context.Context, req *presence.GetPresenceRequest) (*presence.GetPresenceResponse, error) {
	resp, err := s.svc.GetPresence(ctx, &param.GetPresenceRequest{UserIDs: slice.MapFromUint64ToUint(req.GetUserIds())})
	if err != nil {
		return nil, err
	}

	return protobuf.MapGetPresenceResponseToProtobuf(resp), nil
}
