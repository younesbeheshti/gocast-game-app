package param

type GetPresenceRequest struct {
	UserIDs []uint
}

type GetPresenceResponse struct {
	Items []GetPresenceItem
}

type GetPresenceItem struct {
	UserID    uint
	Timestamp int64
}
