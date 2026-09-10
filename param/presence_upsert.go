package param

type UpsertPresenceRequest struct {
	UserID    uint  `json:"user_id"`
	Timestamp int64 `json:"timestamp"`
}

type UpsertPresenceResponse struct{}
