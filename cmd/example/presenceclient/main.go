package main

import (
	"context"
	"fmt"
	presenceClient "github.com/younesbeheshti/gocast_game/adapter/presence"
	"github.com/younesbeheshti/gocast_game/param"
)

func main() {

	//grpcClient := &grpc.ClientConn{}

	client, err := presenceClient.New(":8000")
	defer client.Close()
	if err != nil {
		panic(err)
	}

	resp, err := client.GetPresence(context.Background(), &param.GetPresenceRequest{UserIDs: []uint{1, 2, 3}})
	if err != nil {
		panic(err)
	}

	for _, item := range resp.Items {
		fmt.Println("item;", item.UserID, item.Timestamp)
	}

}
