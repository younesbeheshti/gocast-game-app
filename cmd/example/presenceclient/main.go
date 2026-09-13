package main

import (
	"context"
	"fmt"
	presenceClient "github.com/younesbeheshti/gocast_game/adapter/presence"
	"github.com/younesbeheshti/gocast_game/param"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	//grpcClient := &grpc.ClientConn{}

	conn, err := grpc.NewClient(
		":8000",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := presenceClient.New(conn)
	resp, err := client.GetPresence(context.Background(), &param.GetPresenceRequest{UserIDs: []uint{1, 2, 3}})
	if err != nil {
		panic(err)
	}

	for _, item := range resp.Items {
		fmt.Println("item;", item.UserID, item.Timestamp)
	}
}
