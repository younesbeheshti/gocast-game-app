package protobuf

import (
	"github.com/younesbeheshti/gocast_game/contract/golang/presence"
	"github.com/younesbeheshti/gocast_game/param"
)

func MapGetPresenceResponseToProtobuf(g *param.GetPresenceResponse) *presence.GetPresenceResponse {
	r := &presence.GetPresenceResponse{}

	for _, v := range g.Items {
		r.Items = append(r.Items, &presence.GetPresenceItem{
			UserId:    uint64(v.UserID),
			Timestamp: v.Timestamp,
		})
	}

	return r

}

func MapGetPresenceResponseFromProtobuf(g *presence.GetPresenceResponse) *param.GetPresenceResponse {
	r := &param.GetPresenceResponse{}

	for _, v := range g.Items {
		r.Items = append(r.Items, param.GetPresenceItem{
			UserID:    uint(v.UserId),
			Timestamp: v.Timestamp,
		})
	}

	return r

}
