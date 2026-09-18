package protobufencoder

import (
	"encoding/base64"
	"github.com/younesbeheshti/gocast_game/contract/protogolang/matching"
	"github.com/younesbeheshti/gocast_game/contract/protogolang/notification"
	"github.com/younesbeheshti/gocast_game/entity"
	"github.com/younesbeheshti/gocast_game/pkg/slice"
	"google.golang.org/protobuf/proto"
)

func EncodeEvent(event entity.Event, data any) string {

	var payload []byte

	switch event {
	case entity.MatchingUsersMatchedEvent:
		mu, ok := data.(entity.MatchedPlayers)
		if !ok {
			// TODO: log error
			// TODO: update metrics
			return ""
		}

		pbMu := matching.MatchedUsers{
			Category: string(mu.Category),
			UserIds:  slice.MapFromUintToUint64(mu.UserIDs),
		}

		var err error
		payload, err = proto.Marshal(&pbMu)
		if err != nil {
			// TODO: log error
			// TODO: update metrics
			return ""
		}

	case entity.NotificationEvent:
		mu, ok := data.(entity.Notification)
		if !ok {
			return ""
		}

		pbMu := notification.Notification{
			Type:    mu.Type,
			Payload: mu.Payload,
		}

		var err error
		payload, err = proto.Marshal(&pbMu)
		if err != nil {
			// TODO: log error
			// TODO: update metrics
			return ""
		}
	}

	return base64.StdEncoding.EncodeToString(payload)

}

func DecodeEvent(event entity.Event, data string) any {
	payload, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		// TODO: log error
		// TODO: update metrics
		return nil
	}

	switch event {
	case entity.MatchingUsersMatchedEvent:
		pbMu := &matching.MatchedUsers{}
		if err := proto.Unmarshal(payload, pbMu); err != nil {
			// TODO: log error
			// TODO: update metrics
			return nil
		}

		return entity.MatchedPlayers{
			Category: entity.Category(pbMu.Category),
			UserIDs:  slice.MapFromUint64ToUint(pbMu.UserIds),
		}
	case entity.NotificationEvent:
		pbMu := &notification.Notification{}
		if err := proto.Unmarshal(payload, pbMu); err != nil {
			return nil
		}
		return entity.Notification{
			Type:    pbMu.Type,
			Payload: pbMu.Payload,
		}
	}

	return nil
}
