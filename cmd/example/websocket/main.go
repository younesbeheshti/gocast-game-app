package main

import (
	"fmt"
	"github.com/younesbeheshti/gocast_game/entity"
	"github.com/younesbeheshti/gocast_game/pkg/protobufencoder"
	"net"
	"net/http"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func print() {
	str := protobufencoder.EncodeEvent(entity.NotificationEvent, entity.Notification{
		Type:    "ping",
		Payload: "hello world"})
	fmt.Println("msg:", str)
}

func main() {
	print()
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			// handle error
		}
		go readMessage(conn)
		//go func() {
		//	defer conn.Close()
		//
		//	for {
		//		msg, op, err := wsutil.ReadClientData(conn)
		//		if err != nil {
		//			// handle error
		//			panic(err)
		//		}
		//		err = wsutil.WriteServerMessage(conn, op, msg)
		//		if err != nil {
		//			// handle error
		//			panic(err)
		//		}
		//	}
		//}()
	}))
}

func readMessage(conn net.Conn) {
	defer conn.Close()
	topic := entity.NotificationEvent
	for {
		msg, op, err := wsutil.ReadClientData(conn)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(msg), string(op))
		payload := protobufencoder.DecodeEvent(topic, string(msg))
		data, ok := payload.(entity.Notification)
		if !ok {
			panic("invalid payload")
		}
		fmt.Println(data)

	}
}
