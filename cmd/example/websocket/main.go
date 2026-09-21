package main

import (
	"encoding/json"
	"fmt"
	"github.com/younesbeheshti/gocast_game/entity"
	"net"
	"net/http"
	"time"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

func producer(remoteAddr string, channel chan string) {
	for {
		channel <- remoteAddr
		time.Sleep(time.Second * 5)
	}
}

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, _, _, err := ws.UpgradeHTTP(r, w)
		if err != nil {
			// handle error
		}

		channel := make(chan string)
		go producer(r.RemoteAddr, channel)
		go writeMessage(conn, channel)

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
	for {
		msg, op, err := wsutil.ReadClientData(conn)
		if err != nil {
			panic(err)
		}
		fmt.Println(msg, op)

		var notif entity.Notification
		err = json.Unmarshal(msg, &notif)
		if err != nil {
			panic(err)
		}
		fmt.Println(notif)

	}
}

func writeMessage(conn net.Conn, channel chan string) {
	for data := range channel {
		err := wsutil.WriteServerMessage(conn, ws.OpText, []byte(data))
		if err != nil {
			panic(err)
		}
	}
}
