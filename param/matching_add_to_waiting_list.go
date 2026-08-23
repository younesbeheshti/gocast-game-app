package param

import (
	"github.com/younesbeheshti/gocast_game/entity"
	"time"
)

type AddToWaitingListRequest struct {
	UserID   uint            `json:"user_id"`
	Category entity.Category `json:"category"`
}

type AddToWaitingListResponse struct {
	Timeout time.Duration `json:"timeout_in_nanoseconds"`
}
