package broker

import "github.com/younesbeheshti/gocast_game/entity"

type Publisher interface {
	Publish(event entity.Event, payload string)
}
